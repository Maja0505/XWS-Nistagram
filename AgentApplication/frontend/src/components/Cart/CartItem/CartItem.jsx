import React,{useEffect} from 'react';
import { Typography, Button, Card, CardActions, CardContent, CardMedia } from '@material-ui/core';

import useStyles from './styles';

const CartItem = ({ item, onUpdateCartQty, onRemoveFromCart }) => {
  const classes = useStyles();

  const handleUpdateCartQty = (lineItemId, newQuantity) => onUpdateCartQty(lineItemId, newQuantity);

  const handleRemoveFromCart = (lineItemId) => onRemoveFromCart(lineItemId);

  return (
    <Card className="cart-item">
         <div>
            <img  style={{width:"100%",height:"100%"}}  src={`http://localhost:8070/media/getImage/?id=${item.Product.image}`} title={item.Product.name} />
         </div>
       <CardContent className={classes.cardContent}>
        <Typography variant="h4">{item.Product.name}</Typography>
      </CardContent>
      <CardActions className={classes.cardActions}>
        <div className={classes.buttons}>
          <Button type="button" size="small" onClick={() => handleUpdateCartQty(item.id, item.amount - 1)}>-</Button>
          <Typography>&nbsp;{item.amount}&nbsp;</Typography>
          <Button type="button" size="small" onClick={() => handleUpdateCartQty(item.id, item.amount + 1)}>+</Button>
        </div>
        <Button variant="contained" type="button" color="secondary" style={{backgroundColor:"gray"}} onClick={() => handleRemoveFromCart(item.id)}>Remove</Button>
      </CardActions>
    </Card>
  );
};

export default CartItem;
