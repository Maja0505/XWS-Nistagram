import { Grid, TextField, Button, useRadioGroup } from "@material-ui/core";
import React,{ useState } from "react";
import { Redirect } from "react-router-dom";


import axios from "axios";

const RegistrationForm= () => {
  const [user, setUser] = useState({ FirstName:"",LastName:"", Email: "", Password: "" ,ConfirmedPassword:""});
  const [redirectToLogin, setRedirectToLogin] = useState(false);

  const register = () => {
    axios
      .post("/users/registerUser", user)
      .then((res) => {
          alert("Successfull registration!")
          setRedirectToLogin(true)
      })
      .catch((error, res) => {
        alert(error);
        console.log(error.message);
      });
  };

  return (
    <div style={{width:"800px",margin:"0 auto"}}>
        {redirectToLogin === true && <Redirect to="/login" />}
      <Grid container style={{ marginTop: "10%",marginLeft:"5%"}}>
        <Grid item xs={3}></Grid>
        <Grid item xs={6}>
            <div style={{marginBottom:"10%",marginLeft:"15%"}}>
                <label style={{fontSize:"25px"}}>Registration form</label>
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
              placeholder="First name"
              width="200px"
              height="100px"
              onChange={(e) => setUser({ ...user, FirstName: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Last name"
              width="200px"
              height="100px"
              onChange={(e) => setUser({ ...user, LastName: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Email"
              width="200px"
              height="100px"
              onChange={(e) => setUser({ ...user, Email: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              type="password"
              placeholder="Password"
              onChange={(e) => setUser({ ...user, Password: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              type="password"
              placeholder="Confirm password"
              onChange={(e) => setUser({ ...user, ConfirmedPassword: e.target.value })}
            />
          </form>
        </Grid>
        <Grid item xs={3}></Grid>
      </Grid>
      <div style={{ marginTop: "2%" }}>
        <Button
          variant="contained"
          color="primary"
          style={{ marginRight: "2%",marginLeft:"52%" }}
          onClick={register}
        >
          Register
        </Button>
      </div>
    </div>
  );
};

export default RegistrationForm;