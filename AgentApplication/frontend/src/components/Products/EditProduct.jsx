import React from 'react';
import { Button, Grid, TextField } from "@material-ui/core";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";
import Slider from "react-slick";
import { makeStyles } from "@material-ui/core/styles";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { useEffect } from 'react';
import {useParams,Redirect} from "react-router-dom"
import { AssignmentReturnedTwoTone } from '@material-ui/icons';

const EditProduct=()=>{
  const { productId } = useParams();
  const [redirectToProductList,setRedirectToProductList]=useState(false)
  const [imageChanged,setImageChanged]=useState(false)

    const useStyles = makeStyles((theme) => ({
        settings: {
          dots: true,
          infinite: true,
          speed: 500,
          slidesToShow: 1,
          slidesToScroll: 1,
        },
        rtl: {
          rtl: true,
        },
      }));
    const classes = useStyles();
    const [selectedFile, setSelectedFile] = useState(null);
    const [image, setImage] = useState([]);
    const loggedUserId = localStorage.getItem("id");
    const [fetchedProduct,setFetchedProduct]=useState({ id:"",name:"", description: "", image: "" ,
    availableQuantity:"",price:""})
    const settings = {
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1
      };

    useEffect(() => {
    fetchProduct()
      }, []);

    const fetchProduct =  () => {
        axios
        .post("/products/findById/?id="+productId)
        .then((res) => {
            console.log(res.data)
            setFetchedProduct(res.data)

        })
        .catch((error) => {
          alert(error)
        });
    };

    const ChangeImage = (event) => {
      setImageChanged(true)
      setSelectedFile(null);
      var formData = new FormData();
      var file = event.target.files[0];
      formData.append("myFile", file);
      const reader = new FileReader();
      var url = reader.readAsDataURL(file);
      reader.onloadend = function (e) {
        setSelectedFile(reader.result);
        setImage(formData);
      }.bind(this);
    };

  

    const editProduct= () => {
      console.log(fetchedProduct)
      saveProduct();
    };

   

    const saveProduct = () => {
      var imageId = uuidv4().toString() + "A" + ".jpg";
      if (imageChanged){
        axios
        .post(
          "/media/uploadImage?id=" +
           imageId.substring(0, imageId.length - 4),
          image,
          {
            headers: { "Content-Type": "multipart/form-data" },
          }
        )
        .then((res) => {
          var product = {
            id: fetchedProduct.id,
            name: fetchedProduct.name,
            description: fetchedProduct.description,
            image: imageId ,
            availableQuantity:fetchedProduct.availableQuantity,
            price: fetchedProduct.price
          };
          axios
          .post("/products/update", product)
          .then((res1) => {
            alert("Successfully edit product")
            setRedirectToProductList(true)
            window.location.reload(false);
          })
          .catch((error) => {
          })
        })
        .catch((error) => {
      });

       
    
    }
    else{
      var product = {
        id: fetchedProduct.id,
        name: fetchedProduct.name,
        description: fetchedProduct.description,
        image: fetchedProduct.image ,
        availableQuantity:fetchedProduct.availableQuantity,
        price: fetchedProduct.price
      };
      axios
      .post("/products/update", product)
      .then((res1) => {
        alert("Successfully edit product")
        setRedirectToProductList(true)
        window.location.reload(false);
      })
      .catch((error) => {
      });
    }
     
       
    };


  
  
    return (
        <div style={{width:"800px",margin:"0 auto"}}>
          {redirectToProductList === true && <Redirect to="/" />}
          <Grid container style={{ marginTop: "10%",marginLeft:"5%"}}>
            <Grid item xs={3}></Grid>
            <Grid item xs={6}>
                <div style={{marginBottom:"10%",marginLeft:"15%"}}>
                    <label style={{fontSize:"25px"}}>Edit product</label>
                </div>
              <form
                noValidate
                autoComplete="off"
                style={{ width: "80%", margin: "auto" }}
                
                 
              >
                <TextField
                  color="primary"
                  variant="outlined"
                  size="small"
                  placeholder="Name"
                  width="200px"
                  height="100px"
                  value={fetchedProduct.name}
                  onChange={(e) => setFetchedProduct({ ...fetchedProduct, name: e.target.value })}
                />
                <br></br>
                <br></br>
                <TextField
                  color="primary"
                  variant="outlined"
                  size="small"
                  placeholder="Description"
                  width="200px"
                  height="100px"
                  value={fetchedProduct.description}
                  onChange={(e) => setFetchedProduct({ ...fetchedProduct, description: e.target.value })}
                />
                <br></br>
                <br></br>
                <TextField
                  color="primary"
                  variant="outlined"
                  size="small"
                  placeholder="Available quantity"
                  width="200px"
                  height="100px"
                  value={fetchedProduct.availableQuantity}
                  onChange={(e) => setFetchedProduct({ ...fetchedProduct, availableQuantity:e.target.value })}
                />
                <br></br>
                <br></br>
                <TextField
                  color="primary" 
                  variant="outlined"
                  size="small"
                  placeholder="Price"
                  value={fetchedProduct.price}
                  onChange={(e) => setFetchedProduct({ ...fetchedProduct, price:e.target.value })}
                /> 
                <br></br>
                <br></br>
                <Button
                        variant="contained"
                        component="label"
                        style={{ margin: "auto" }} >
                        Change image
                        <input
                        hidden
                        accept="image/*"
                        multiple
                        type="file"
                        name="myFile"
                        onChange={(event) => ChangeImage(event)}
                        />
                </Button>
                <Grid container style={{ marginTop: "3%" }}>
            <Grid item xs={2} />
            <Grid item xs={8}>
              <div>
                <Slider {...settings}>
                    
                    <div>
                        <img
                          width="200px"
                          height="200px"
                          src={selectedFile===null ?  `http://localhost:8070/media/getImage/?id=${fetchedProduct.image}` : selectedFile}
                        />
                    </div>
                </Slider>
              </div>
            </Grid>
            <Grid item xs={2} />
          </Grid>
              </form>
            </Grid>
            <Grid item xs={3}></Grid>
          </Grid>
          <div style={{ marginTop: "2%" }}>
            <Button
              variant="contained"
              color="primary"
              style={{ marginRight: "2%",marginLeft:"52%" }}
              onClick={editProduct}
            >
              Submit
            </Button>
          </div>
        </div>
      );
}
export default EditProduct 