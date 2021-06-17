import {
  AppBar,
  Toolbar,
  Grid,
  Button,
  TextField,
  Typography,
  Avatar,
  Grow,
  Paper,
  Popper,
  MenuItem,
  MenuList,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import { Autocomplete } from "@material-ui/lab";
import { Redirect } from "react-router-dom";
import { useState, useRef, useEffect } from "react";
import React from "react";
import axios from "axios";
import avatar from "../images/nistagramAvatar.jpg";
import {
  HomeOutlined,
  EmailOutlined,
  FavoriteBorderOutlined,
  ExploreOutlined,
  Settings,
  BookmarkBorderOutlined,
  AccountCircleOutlined,
  ThumbsUpDownOutlined,
} from "@material-ui/icons";

import ClickAwayListener from "@material-ui/core/ClickAwayListener";

const NavBar = () => {
  const username = localStorage.getItem("username");

  const [searchedUser, setSearchedUser] = useState([]);
  const [searchedUsername, setSearchedUsername] = useState();
  const [open, setOpen] = useState(false);
  const anchorRef = useRef(null);

  const [redirectToSearchedUser, setRedirection] = useState(false);

  const logout = () => {
    clearLocalStorage();
  };

  const clearLocalStorage = () => {
    localStorage.clear();
  };

  const handleChangeInput = (text) => {
    if (text.length !== 0) {
      axios
        .get("/api/user/search/" + text)
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

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event) => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
  };

  function handleListKeyDown(event) {
    if (event.key === "Tab") {
      event.preventDefault();
      setOpen(false);
    }
  }

  const prevOpen = useRef(open);

  useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;
  }, [open]);

  const dropDowMenuForProfile = (
    <Popper
      open={open}
      anchorEl={anchorRef.current}
      role={undefined}
      transition
      disablePortal
      style={{ width: "15%", zIndex: "1" }}
    >
      {({ TransitionProps, placement }) => (
        <Grow
          {...TransitionProps}
          style={{
            transformOrigin:
              placement === "bottom" ? "center top" : "center bottom",
          }}
        >
          <Paper>
            <ClickAwayListener onClickAway={handleClose}>
              <MenuList
                autoFocusItem={open}
                id="menu-list-grow"
                onKeyDown={handleListKeyDown}
              >
                <MenuItem onClick={handleClose}>
                  <Grid container>
                    <Grid item xs={3}>
                      <Link
                        to={"/homePage/" + `${username}`}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <AccountCircleOutlined />
                      </Link>
                    </Grid>
                    <Grid item xs={9}>
                      <Link
                        to={"/homePage/" + `${username}`}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <div style={{ width: "100%" }}>Profile</div>
                      </Link>
                    </Grid>
                  </Grid>
                </MenuItem>
                <MenuItem onClick={handleClose}>
                  <Grid container>
                    <Grid item xs={3}>
                      <Link
                        to={"/" + `${username}` + "/saved"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <BookmarkBorderOutlined />
                      </Link>
                    </Grid>
                    <Grid item xs={9}>
                      <Link
                        to={"/" + `${username}` + "/saved"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <div style={{ width: "100%" }}>Saved</div>
                      </Link>
                    </Grid>
                  </Grid>
                </MenuItem>
                <MenuItem onClick={handleClose}>
                  <Grid container>
                    <Grid item xs={3}>
                      <Link
                        to={"/accounts/edit/"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <Settings />
                      </Link>
                    </Grid>
                    <Grid item xs={9}>
                      <Link
                        to={"/accounts/edit/"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <div style={{ width: "100%" }}>Settings</div>
                      </Link>
                    </Grid>
                  </Grid>
                </MenuItem>
                <MenuItem onClick={handleClose}>
                  <Grid container>
                    <Grid item xs={3}>
                      <Link
                        to={"/" + `${username}` + "/liked-disliked/"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <ThumbsUpDownOutlined />
                      </Link>
                    </Grid>
                    <Grid item xs={9}>
                      <Link
                        to={"/" + `${username}` + "/liked-disliked/"}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <div style={{ width: "100%" }}>
                          Liked/Disliked Posts
                        </div>
                      </Link>
                    </Grid>
                  </Grid>
                </MenuItem>
                <MenuItem onClick={logout} style={{ borderTop: "inset 0.5px" }}>
                  <Grid container>
                    <a
                      href={"/"}
                      style={{
                        textDecoration: "none",
                        color: "black",
                        width: "100%",
                      }}
                    >
                      Logout
                    </a>
                  </Grid>
                </MenuItem>
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

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
          <Grid item xs={2} style={{ margin: "auto" }}>
            <HomeOutlined
              style={{
                color: "gray",
                width: "33px",
                height: "33px",
                marginRight: "5%",
              }}
            />
            <EmailOutlined
              style={{
                color: "gray",
                width: "29px",
                height: "29px",
              }}
            />
          </Grid>
          <Grid
            container
            item
            xs={2}
            style={{ textAlign: "left", margin: "auto" }}
          >
            <ExploreOutlined
              style={{
                color: "gray",
                width: "29px",
                height: "33px",
                marginLeft: "5%",
              }}
            />
            <FavoriteBorderOutlined
              style={{
                color: "gray",
                width: "29px",
                height: "33px",
                marginLeft: "5%",
              }}
            />
            <div>
              <Avatar
                alt="N"
                src={avatar}
                ref={anchorRef}
                aria-controls={open ? "menu-list-grow" : undefined}
                aria-haspopup="true"
                onClick={handleToggle}
                style={{
                  width: "29px",
                  height: "33px",
                  marginTop: "1.5%",
                  marginLeft: "5%",
                  cursor: "pointer",
                }}
              />

              {dropDowMenuForProfile}
            </div>
          </Grid>
          <Grid item xs={2}></Grid>
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
