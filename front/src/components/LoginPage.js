import { Grid, TextField, Button } from "@material-ui/core";
import { useState } from "react";

import axios from "axios";

const LoginPage = () => {
  const [user, setUser] = useState({ username: "", password: "" });

  const singUp = () => {
    window.location.href = "http://localhost:3000/registration";
  };

  const loginIn = () => {
    axios.get("http://localhost:8000/user/" + user.username).then((res) => {
      console.log(res.data);
    });
  };

  return (
    <div>
      <Grid container style={{ marginTop: "6%" }}>
        <Grid item xs={3}></Grid>
        <Grid item xs={6}>
          <form
            noValidate
            autoComplete="off"
            style={{ width: "80%", margin: "auto" }}
          >
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Username"
              onChange={(e) => setUser({ ...user, username: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              type="password"
              placeholder="Password"
              onChange={(e) => setUser({ ...user, password: e.target.value })}
            />
          </form>
        </Grid>
        <Grid item xs={3}></Grid>
      </Grid>
      <div style={{ marginTop: "2%" }}>
        <Button
          variant="contained"
          color="primary"
          style={{ marginRight: "2%" }}
          onClick={loginIn}
        >
          LOG IN
        </Button>
        <Button variant="contained" color="inherit" onClick={singUp}>
          SING UP
        </Button>
      </div>
    </div>
  );
};

export default LoginPage;
