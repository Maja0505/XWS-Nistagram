import React from 'react';
import { Card, CardMedia, CardContent, CardActions, Typography, IconButton } from '@material-ui/core';
import { AddShoppingCart,DeleteOutline,Edit } from '@material-ui/icons';
import {Link} from "react-router-dom"
import useStyles from './styles';

const Product = ({ product, onAddToCart,onEditProduct,onDeleteProduct }) => {
  const classes = useStyles();

  const handleAddToCart = () => onAddToCart(product.id, 1);
  const handleEditProduct = () => onEditProduct(product.id);
  const handleDeleteProduct = () => onDeleteProduct(product.id);

  return (
    <Card className={classes.root}>
      <div>
      <img  style={{width:"100%",height:"100%"}}  src={`http://localhost:8070/media/getImage/?id=${product.image}`} title={product.name} />
      </div>
      <CardContent>
        <div className={classes.cardContent}>
          <Typography gutterBottom variant="h5" component="h2">
            {product.name}
          </Typography>
          <Typography gutterBottom variant="h5" component="h2">
            ${product.price}
          </Typography>
        </div>
        <Typography dangerouslySetInnerHTML={{ __html: product.description }} variant="body2" color="textSecondary" component="p" />
      </CardContent>
       
        <CardActions disableSpacing className={classes.cardActions}>
        {localStorage.getItem("role")!="Agent" && localStorage.getItem("loggedIn")==="true" &&
        <IconButton aria-label="Add to Cart" onClick={handleAddToCart}>
          <AddShoppingCart />
        </IconButton> }
      
        {localStorage.getItem("role")==="Agent" && 
        <Link to={"/editProduct/" + product.id}>
        <IconButton aria-label="Add to Cart">
          <Edit />
        </IconButton>
        </Link>}
        {localStorage.getItem("role")==="Agent" && <IconButton aria-label="Add to Cart" onClick={handleDeleteProduct}>
          <DeleteOutline />
        </IconButton>}
        
      </CardActions>
    </Card>
  );
};

export default Product;

