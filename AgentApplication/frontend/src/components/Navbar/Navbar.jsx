import React, { useState,useEffect  } from 'react';
import { AppBar, Toolbar, IconButton, Badge, MenuItem, Menu, Typography } from '@material-ui/core';
import { ShoppingCart,AccountCircle,Add,ArrowForward} from '@material-ui/icons';
import {Link,useLocation } from 'react-router-dom';

import logo from '../../assets/commerce.png';
import useStyles from './styles';


const PrimarySearchAppBar = ({ totalItems,handleClickOnCart}) => {
  const classes = useStyles();
  const location = useLocation();

  const logout= () => {
        localStorage.setItem("username","");
        localStorage.setItem("id","");
        localStorage.setItem("role","")
        localStorage.setItem("loggedIn", "false")
        window.location.reload(false);
  };

  

  return (
    <>
      <AppBar position="fixed" className={classes.appBar} color="inherit">
        <Toolbar>
          <Typography component={Link} to="/" variant="h6" className={classes.title} color="inherit">
            <img src={logo} alt="webShop" height="25px" className={classes.image} /> Web shop
          </Typography>
          <div className={classes.grow} />
          {location.pathname === '/' && (
          <div className={classes.button}>
          
           {(localStorage.getItem("role")==="User") && 
           <IconButton component={Link} to="/cart" aria-label="Show cart items" color="inherit" onClick={handleClickOnCart}>
              <Badge badgeContent={totalItems} color="secondary">
                <ShoppingCart />
              </Badge>
            </IconButton>
          }
          
            {(localStorage.getItem("role")==="Agent")  &&<IconButton component={Link} to="/addProduct" aria-label="Add product" color="inherit">
               <Add />
           </IconButton>}
           {localStorage.getItem("loggedIn")==="false" ?
            <IconButton component={Link} to="/login" aria-label="Login" color="inherit">
               <AccountCircle />
           </IconButton>: 
           <IconButton component={Link}  aria-label="Logout" color="inherit" onClick={logout}>
               <ArrowForward />
           </IconButton> 
           }
         </div>
          )} 
        </Toolbar>
      </AppBar>
    </>
  );
};

export default PrimarySearchAppBar;
