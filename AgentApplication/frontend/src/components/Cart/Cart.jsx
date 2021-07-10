import React,{useEffect,useState} from 'react';
import { Container, Typography, Button, Grid } from '@material-ui/core';
import { Link } from 'react-router-dom';

import CartItem from './CartItem/CartItem';
import useStyles from './styles';

const Cart = ({ cart, onUpdateCartQty, onRemoveFromCart }) => {
  const classes = useStyles();
  const [totalSum,setTotalSum]=useState(0)

  const renderEmptyCart = () => (
    <Typography variant="subtitle1">No items in your shopping cart.
    </Typography>
  );
  

 if (!cart.orders) return 'Loading';

  const renderCart = () => (
    <>
      <Grid container spacing={3}>
        {cart.orders.map((lineItem) => (
          <Grid item xs={12} sm={4} key={lineItem.id}>
            <CartItem item={lineItem} onUpdateCartQty={onUpdateCartQty} onRemoveFromCart={onRemoveFromCart} />
          </Grid>
        ))}
      </Grid>
      <div className={classes.cardDetails}>
        <div>
           <Button className={classes.checkoutButton} style={{marginTop:"20px",marginRight:"20px"}} component={Link} to="/checkout" size="large" type="button" variant="contained" color="primary">Checkout</Button>
        </div>
      </div>
    </>
  );

  return (
    <Container>
      <div className={classes.toolbar} />
      <Typography className={classes.title} variant="h3" gutterBottom>Your Shopping Cart</Typography>
      { !cart.orders.length ? renderEmptyCart() : renderCart() }
    </Container>
  );
};

export default Cart;
