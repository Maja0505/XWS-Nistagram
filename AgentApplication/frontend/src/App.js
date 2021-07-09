import React, { useState, useEffect } from 'react';
import { CssBaseline } from '@material-ui/core';
import { BrowserRouter as Router, Switch, Route,Redirect } from 'react-router-dom';

import { Navbar, Products, Cart, Checkout,Login,Register,AddProduct,EditProduct } from './components';
import axios from "axios";

const App = () => {
  const [products, setProducts] = useState([]);
  const [cart, setCart] = useState({});
  const [order, setOrder] = useState({});
  const [errorMessage, setErrorMessage] = useState('');
  const [redirectToProductList,setRedirectToProductList]=useState(false)

  const clickedOnCart =  () => {
   setCart({})
   fetchCart()
  }
  const fetchProducts = async () => {
    axios
    .get("/products/findAll")
    .then((res) => {
      if (res.data) {
        console.log(res.data)
        setProducts(res.data);
      } else {
        setProducts([]);
      }
    })
    .catch((error) => {
      setProducts([]);
    });
  };

  const fetchCart = async () => {
    axios
    .post("/shoppingCart/findByUser?userId="+localStorage.getItem("id"))
    .then((res) => {
      setCart(res.data)
      console.log(res.data)

    })
    .catch((error) => {
      alert(error)
    });
  };

  const handleAddToCart = async (productId, quantity) => {
    axios
    .post("/shoppingCart/addOrderToCart",{	id:"",
                                            ProductID:productId,
                                            Amount:quantity,
                                            ShoppingCartID:cart.id})
    .then((res) => {
        console.log("nakon dodavanja: ",res.data)
    })
    .catch((error) => {
      alert(error)
    });
  };


  const handleDeleteProduct =  (productId) => {
    axios
    .post("/products/delete/?id="+productId)
    .then((res) => {
      fetchProducts();

    })
    .catch((error) => {
      alert(error)
    });
  };


  const handleUpdateCartQty = async (lineItemId, quantity) => {
    axios
    .post("/shoppingCart/updateOrderQuantity?orderId="+lineItemId+"&amount="+quantity)
    .then((res) => {
      console.log(res.data)
      fetchCart()
    })
    .catch((error) => {
      alert(error)
    });
  };

  const handleRemoveFromCart = async (lineItemId) => {
    axios
    .post("/shoppingCart/deleteOrderFromCart?orderId="+lineItemId+"&shoppingCartId="+cart.id)
    .then((res) => {
      fetchCart()
    })
    .catch((error) => {
      alert(error)
    });
  };

  
 
  const handleCheckout = async (addressInfo) => {
    let totalSum=0;
    for(let i=0;i<cart.orders.length;i++){
      totalSum += parseInt(cart.orders[i].Product.price) * parseInt(cart.orders[i].amount)
    }
    axios
    .post("/address/createAddress",addressInfo)
    .then((res) => {
      console.log(res.data)
      var purchase={
        id:"",
        orders:cart.orders,
        totalPrice:totalSum,
        user_id:cart.user_id,
        address_id:res.data.id
      }
      axios
    .post("/purchase/createPurchase",purchase)
    .then((res) => {
      alert("Items successfully purchased")
      setRedirectToProductList(true)

    })
    .catch((error) => {
      alert(error)
    });

    })
    .catch((error) => {
      alert(error)
    });
    
  };

  useEffect(() => {
    fetchProducts();
    if(localStorage.getItem("loggedIn")===true)
    {
      fetchCart();
    }
  }, []);

  

  

  

  return (
    
    <Router>
      
      <div style={{ display: 'flex' }}>
      {redirectToProductList === true && <Redirect to="/" />}
        <CssBaseline />
        <Navbar totalItems={cart.total_items} handleClickOnCart={clickedOnCart}  />
        <Switch>
          <Route exact path="/">
            <Products products={products} onAddToCart={handleAddToCart}  onDeleteProduct={handleDeleteProduct} handleUpdateCartQty />
          </Route>
          <Route exact path="/cart">
            <Cart cart={cart} onUpdateCartQty={handleUpdateCartQty} onRemoveFromCart={handleRemoveFromCart} />
          </Route>
          <Route exact path="/login">
            <Login />
          </Route>

          <Route path="/checkout" exact>
            <Checkout  order={order} handleCheckout={handleCheckout} error={errorMessage} />
          </Route>
          <Route path="/register" exact>
            <Register />
          </Route>
          <Route path="/addProduct" exact>
            <AddProduct />
          </Route>
          <Route
              exact
              path="/editProduct/:productId"
              render={(props) => <EditProduct {...props} />}
            ></Route>
        </Switch>
      </div>
    </Router>
  );
};

export default App;
