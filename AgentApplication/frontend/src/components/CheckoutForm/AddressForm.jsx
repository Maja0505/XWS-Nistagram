import React, { useState, useEffect } from 'react';
import { InputLabel, Select, MenuItem, Button, Grid, Typography,TextField } from '@material-ui/core';
import { useForm, FormProvider } from 'react-hook-form';
import { Link } from 'react-router-dom';

import { commerce } from '../../lib/commerce';
import FormInput from './CustomTextField';

const AddressForm = ({ handleCheckout }) => {
  const handleSubmit= (addressInfo) => handleCheckout(addressInfo);
  const [addressInfo,SetAddressInfo]=useState({ id:null,address:"",city:"", zip: "", country: ""})


  return (
    <>
      <Typography variant="h6" gutterBottom>Shipping address</Typography>
          <Grid container spacing={3}>
            <TextField required style={{margin:"20px"}}  label="Address" onChange={(e) => SetAddressInfo({ ...addressInfo, address:e.target.value })} />
            <TextField required style={{margin:"20px"}} label="City" onChange={(e) => SetAddressInfo({ ...addressInfo, city:e.target.value })} />
            <TextField required style={{margin:"20px"}} label="Zip / Postal code" onChange={(e) => SetAddressInfo({ ...addressInfo, zip:e.target.value })} />
            <TextField required style={{margin:"20px"}} label="ShippingCountry" onChange={(e) => SetAddressInfo({ ...addressInfo, country:e.target.value })} />
          </Grid>
          <br />
          <div style={{ display: 'flex', justifyContent: 'space-between' }}>
            <Button component={Link} variant="outlined" to="/cart">Back to Cart</Button>
            <Button type="submit" variant="contained" color="primary" onClick={(e)=>handleSubmit(addressInfo)}>Confirm</Button>
          </div>
    </>
  );
};

export default AddressForm;
