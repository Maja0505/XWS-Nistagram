import React from 'react';
import { Button, Grid, TextField } from "@material-ui/core";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";
import Slider from "react-slick";
import { makeStyles } from "@material-ui/core/styles";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import {useParams,Redirect} from "react-router-dom"

const AddProduct=({})=>{
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
    const [product, setProduct] = useState({ ID:"",Name:"", Description: "", Image: "" ,
    AvailableQuantity:"",Price:""});
    const classes = useStyles();
    const [selectedFile, setSelectedFile] = useState([]);
    const [image, setImage] = useState([]);
    const [description, setDescription] = useState("");
    const loggedUserId = localStorage.getItem("id");
    const [isVideo, setIsVideo] = useState([]);
    const [imagesIdsForSave, setImagesIdsForSave] = useState([]);
    const [puklaSlika, setPuklaSlika] = useState(false);
    const [redirectToProductList,setRedirectToProductList]=useState(false)

    const settings = {
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1,
      };

    const createProduct= () => {
      setProduct({...product,Price:parseInt(product.price,10)+1});
      setProduct({...product,AvailableQuantity:parseFloat(product.AvailableQuantity)});
      
      for (let index = 0; index < image.length; index++) {
          uploadImage(image[index], index);
      }
      saveProduct();
    };

   
   

    const saveProduct = () => {
      var productForSave = {
        id: product.ID,
        name: product.Name,
        description: product.Description,
        image:imagesIdsForSave[0] ,
        availableQuantity:product.AvailableQuantity,
        price: product.Price
      };
      console.log(productForSave)
      if (!puklaSlika) {
        axios
          .post("/products/create", productForSave)
          .then((res1) => {
            alert("Successfully added product")
            setPuklaSlika(false);
            setRedirectToProductList(true)
            window.location.reload(false);

          })
          .catch((error) => {
            alert(error)
            setPuklaSlika(false);
          });
      }
    };


    const uploadImage = (imageForUpload, index) => {
        var imageId = uuidv4().toString() + "A" + ".jpg";
        var array = imagesIdsForSave;
        array.push(imageId);
        setImagesIdsForSave(array);
        axios
          .post(
            "/media/uploadImage?id=" +
             imageId.substring(0, imageId.length - 4) +
              "&formKey=" +
              "image" +
              index,
            imageForUpload,
            {
              headers: { "Content-Type": "multipart/form-data" },
            }
          )
          .then((res) => {})
          .catch((error) => {
            setPuklaSlika(true);
          });
      };

    const HandleUploadMedia = (event) => {
        setSelectedFile([]);
        setIsVideo([]);
        setImage([]);
    
        var formData = new FormData();
        for (let index = 0; index < event.target.files.length; index++) {
          if (event.target.files[index].type === "video/mp4") {
            var array = isVideo;
            array.push(true);
            setIsVideo(array);
          } else {
            var array = isVideo;
            array.push(false);
            setIsVideo(array);
          }
    
          var file = event.target.files[index];
          formData.append("myFile",file);
          const reader = new FileReader();
          var url = reader.readAsDataURL(file);
          reader.onloadend = function (e) {
            setSelectedFile((prevState) => [...prevState, reader.result]);
          }.bind(this);
          setImage((prevState) => [...prevState,formData]);
          setProduct({ ...product, Image:imagesIdsForSave[0] })
          console.log(formData);
        }
        console.log(isVideo);
      };

      
  
    return (
        <div style={{width:"800px",margin:"0 auto"}}>
          {redirectToProductList === true && <Redirect to="/" />}
          <Grid container style={{ marginTop: "10%",marginLeft:"5%"}}>
            <Grid item xs={3}></Grid>
            <Grid item xs={6}>
                <div style={{marginBottom:"10%",marginLeft:"15%"}}>
                    <label style={{fontSize:"25px"}}>Add product</label>
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
                  onChange={(e) => setProduct({ ...product, Name: e.target.value })}
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
                  onChange={(e) => setProduct({ ...product, Description: e.target.value })}
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
                  onChange={(e) => setProduct({ ...product, AvailableQuantity:e.target.value })}
                />
                <br></br>
                <br></br>
                <TextField
                  color="primary" 
                  variant="outlined"
                  size="small"
                  placeholder="Price"
                  onChange={(e) => setProduct({ ...product, Price:e.target.value })}
                /> 
                <br></br>
                <br></br>
                <Button
                        variant="contained"
                        component="label"
                        style={{ margin: "auto" }} >
                        {selectedFile.length === 0 ? `Upload image` : `Change image`}
                        <input
                        hidden
                        accept="image/*,video/mp4,video/x-m4v,video/*"
                        multiple
                        type="file"
                        name="myFile"
                        onChange={(event) => HandleUploadMedia(event)}
                        />
                </Button>
                <Grid container style={{ marginTop: "3%" }}>
            <Grid item xs={2} />
            <Grid item xs={8}>
              <div>
                <Slider {...settings}>
                  {selectedFile.map((media, index) => (
                    <div>
                      {!isVideo[index] && (
                        <img
                          width="200px"
                          height="200px"
                          src={selectedFile[index]}
                        />
                      )}
                      {isVideo[index] && (
                        <video width="100%" controls>
                          <source src={selectedFile[index]} type="video/mp4" />
                        </video>
                      )}
                    </div>
                  ))}
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
              onClick={createProduct}
            >
              Submit
            </Button>
          </div>
        </div>
      );
}
export default AddProduct 