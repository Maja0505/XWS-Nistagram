import { Grid, TextField, Button } from "@material-ui/core";
import React,{ useState } from "react";
import { Link,Redirect } from 'react-router-dom';


import axios from "axios";

const Login = () => {
  const [user, setUser] = useState({ Email: "", Password: "" });
  const[redirect,setRedirect] = useState(false);
  

  const login = () => {
   
    axios
      .post("/users/login" , user)
      .then((res) => {
        console.log(res.data)
      
        localStorage.setItem("username", res.data.email);
        localStorage.setItem("id", res.data.id);
        console.log(res.data.id)
        localStorage.setItem("role", res.data.role)
        localStorage.setItem("loggedIn", "true")
        setRedirect(true)
        console.log(localStorage.getItem("role"))
       
      })
      .catch((error) => {
        alert("Wrong username or password");
      });
  };

  return (
    <div style={{width:"800px",margin:"0 auto"}}>
      {redirect === true && <Redirect to="/" />}
      <Grid container style={{ marginTop: "20%",marginLeft:"5%"}}>
        <Grid item xs={3}></Grid>
        <Grid item xs={6}>
            <label style={{marginLeft:"28%",fontSize:"25px"}}>Log in</label>
          <form
            noValidate
            autoComplete="off"
            style={{ width: "80%", margin: "auto" }}
            
            
          >
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
          </form>
        </Grid>
        <Grid item xs={3}></Grid>
      </Grid>
      <div style={{marginTop:"2%",marginLeft:"31%"}}>
          <label>Don't have an account? Click here to </label><Link to="/register">register</Link>
      </div>
      <div style={{ marginTop: "2%" }}>
        <Button
          variant="contained"
          color="primary"
          style={{ marginRight: "2%",marginLeft:"52%" }}
          onClick={login}
        >
          LOG IN
        </Button>
      </div>
    </div>
  );
};

export default Login;