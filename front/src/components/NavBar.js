import {
  AppBar,
  Toolbar,
  Grid,
  Button,
  TextField,
  Typography,
  Avatar,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import { Autocomplete } from "@material-ui/lab";
import { Redirect } from "react-router-dom";
import { useState } from "react";
import React from "react";
import axios from "axios";
import avatar from "../images/nistagramAvatar.jpg";

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
          setSearchedUser([]);
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

  const searchBar = (
    <Grid item xs={6} style={{ textAlign: "center" }}>
      <Autocomplete
        freeSolo
        renderOption={(option, { selected }) => (
          <React.Fragment>
            <Grid container>
              <Grid item xs={2}>
                <Avatar
                  alt="N"
                  src={avatar}
                  style={{ border: "1px solid" }}
                ></Avatar>
              </Grid>
              <Grid item xs={10} style={{ marginTop: "3%" }}>
                {option}
              </Grid>
            </Grid>
          </React.Fragment>
        )}
        options={
          searchedUser.length !== 0
            ? searchedUser.map((option) => option.Username)
            : []
        }
        onChange={(event, value) => goToUserProfile(value)}
        renderInput={(params) => (
          <>
            <TextField
              {...params}
              variant="outlined"
              size="small"
              style={{ width: "70%" }}
              onChange={(e) => handleChangeInput(e.target.value)}
            ></TextField>
          </>
        )}
      />
    </Grid>
  );

  const NavBarForUnregisteredUser = (
    <Toolbar style={{ backgroundColor: "white" }}>
      <Grid container>
        <Grid item xs={4}>
          <Typography
            variant="h5"
            style={{ color: "gray", fontFamily: "cursive", margin: "auto" }}
          >
            Nistagram
          </Typography>
        </Grid>
        <Grid item xs={8} container style={{ textAlign: "right" }}>
          {searchBar}
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
        <Grid item xs={4}>
          <Typography
            variant="h5"
            style={{ color: "gray", fontFamily: "cursive", margin: "auto" }}
          >
            Nistagram
          </Typography>
        </Grid>
        <Grid item xs={8} container style={{ textAlign: "right" }}>
          {searchBar}
          <Grid item xs={2} />
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
        {(username === null || username === undefined) &&
          NavBarForUnregisteredUser}
        {username !== null && username !== undefined && NavBarForRegistredUser}
      </AppBar>
    </>
  );
};

export default NavBar;
