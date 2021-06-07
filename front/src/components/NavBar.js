import { AppBar, Toolbar, Grid, Button, TextField } from "@material-ui/core";
import { Link } from "react-router-dom";
import { Autocomplete } from "@material-ui/lab";
import { Redirect } from "react-router-dom";
import { useState } from "react";

import axios from "axios";

const NavBar = () => {
  const username = localStorage.getItem("username");

  const [searchedUser, setSearchedUser] = useState([]);
  const [searchedUsername, setSearchedUsername] = useState();

  const [redirectToSearchedUser, setRedirection] = useState(false);

  const clearLocalStorage = () => {
    localStorage.clear();
  };

  const handleChangeInput = (text) => {
    if (text.length !== 0) {
      axios
        .get("/user/search/" + text)
        .then((res) => {
          setSearchedUser(res.data);
        })
        .catch((error) => {
          alert(error.response.data);
        });
    } else {
      setSearchedUser([]);
    }
  };

  const goToUserProfile = (username) => {
    if (username !== undefined && username !== null) {
      setSearchedUsername("/homePage/" + username);
      setRedirection(true);
    }
  };

  const NavBarForUnregisteredUser = (
    <Toolbar style={{ backgroundColor: "white" }}>
      <Grid container>
        <Grid item xs={4}></Grid>
        <Grid item xs={8} container style={{ textAlign: "right" }}>
          <Grid item xs={6} style={{ textAlign: "center" }}>
            <Autocomplete
              freeSolo
              options={
                searchedUser.length !== 0
                  ? searchedUser.map((option) => option.Username)
                  : []
              }
              onChange={(event, value) => goToUserProfile(value)}
              renderInput={(params) => (
                <TextField
                  {...params}
                  variant="outlined"
                  size="small"
                  style={{ width: "70%" }}
                  onChange={(e) => handleChangeInput(e.target.value)}
                />
              )}
            />
          </Grid>
          <Grid item xs={2} />
          <Grid item xs={2}>
            <Button variant="contained" color="primary">
              <Link
                to="/login"
                style={{ textDecoration: "none", color: "white" }}
              >
                Log in
              </Link>
            </Button>
          </Grid>
          <Grid item xs={2}>
            <Button variant="text" onClick={clearLocalStorage}>
              <Link to="/registration" style={{ textDecoration: "none" }}>
                Sing up
              </Link>
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Toolbar>
  );

  const NavBarForRegistredUser = (
    <Toolbar style={{ backgroundColor: "white" }}>
      <Grid container>
        <Grid item xs={6}></Grid>
        <Grid item xs={6} container style={{ textAlign: "right" }}>
          <Grid item xs={2} />
          <Grid item xs={2} />
          <Grid item xs={2}></Grid>
          <Grid item xs={2}></Grid>
          <Grid item xs={2}>
            <Button variant="text" onClick={clearLocalStorage}>
              <a href="/" style={{ textDecoration: "none" }}>
                Sing out
              </a>
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Toolbar>
  );

  return (
    <>
      {redirectToSearchedUser === true && <Redirect to={searchedUsername} />}
      <AppBar position="static">
        {(username === null || username === undefined) && NavBarForUnregisteredUser}
        {username !== null && username !== undefined && NavBarForRegistredUser}
      </AppBar>
    </>
  );
};

export default NavBar;
