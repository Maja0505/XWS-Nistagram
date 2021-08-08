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
  RoomRounded,
} from "@material-ui/icons";
import Badge from "@material-ui/core/Badge";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";

const NavBar = () => {
  const username = localStorage.getItem("username");
  const loggedUserId = localStorage.getItem("id");

  const [searchedContent, setSearchedContent] = useState([]);
  const [redirectionString, setRedirectionString] = useState();
  const [open, setOpen] = useState(false);
  const [openNotifications, setOpenNotifications] = useState(false);
  const [allNotifications, setAllNotifications] = useState([]);

  const anchorRef = useRef(null);
  const [isHastag, setIsHastag] = useState(false);
  const [isUser, setIsUser] = useState(false);
  const [invisible, setInvisible] = useState(
    localStorage.getItem("invisibleNotification")
  );

  const [redirection, setRedirection] = useState(false);

  const logout = () => {
    clearLocalStorage();
  };

  const clearLocalStorage = () => {
    localStorage.clear();
  };

  const handleChangeInput = (text) => {
    if (text.length !== 0) {
      if (text.substring(0, 1) === "#") {
        setIsHastag(true);
        setIsUser(false);
        if (text.length > 1) {
          axios
            .get("/api/post/get-tag-suggestions/" + text.substring(1))
            .then((res) => {
              console.log(res.data);
              setSearchedContent(res.data);
            });
        } else {
          axios.get("/api/post/get-all-tags").then((res) => {
            console.log(res.data);
            setSearchedContent(res.data);
          }).catch((error) => {
            //console.log(error);
          });
        }
      } else {
        setIsHastag(false);
        setIsUser(true);
        axios
          .get("/api/user/search/" + username + "/" + text)
          .then((res) => {
            console.log(res.data);
            setSearchedContent(res.data);
            axios
              .get("/api/post/get-location-suggestions/" + text)
              .then((res) => {
                console.log(res.data);
                if (res.data !== null) {
                  setSearchedContent((prevState) => [
                    ...prevState,
                    ...res.data,
                  ]);
                }
              })
              .catch((error) => {
                setSearchedContent([]);
              });
          })
          .catch((error) => {
            axios
              .get("/api/post/get-location-suggestions/" + text)
              .then((res) => {
                console.log(res.data);
                if (res.data !== null) {
                  setSearchedContent(res.data);
                }
              })
              .catch((error) => {
                setSearchedContent([]);
              });
          });
      }
    } else {
      setSearchedContent([]);
    }
  };

  const goToSearchContent = (content) => {
    if (isUser && content !== null) {
      if (content.Username !== undefined) {
        setRedirectionString("/homePage/" + content.Username);
      } else {
        setRedirectionString("/explore/locations/" + content + "/");
      }
    }
    if (isHastag && content !== null) {
      setRedirectionString("/explore/tags/" + content.substring(1) + "/");
    }

    setRedirection(true);
  };

  const handleToggle = () => {
    setOpenNotifications(false);
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event) => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
    setOpenNotifications(false);
  };

  function handleListKeyDown(event) {
    if (event.key === "Tab") {
      event.preventDefault();
      setOpen(false);
    }
  }

  const prevOpen = useRef(open);
  const prevOpenNotifications = useRef(openNotifications);

  useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    if (prevOpenNotifications.current === true && openNotifications === false) {
      anchorRef.current.focus();
    }

    prevOpenNotifications.current = open;
    prevOpen.current = open;
  }, [open, openNotifications]);

  const handleNotificationButton = () => {
    if (!openNotifications) {
      axios.get("/api/notification/channels/" + loggedUserId).then((res) => {
        console.log(res.data);
        if (res.data) {
          setAllNotifications(res.data);
        }
        setOpenNotifications((prevOpenNotifications) => !prevOpenNotifications);
        setOpen(false);
        localStorage.setItem("invisibleNotification", true);
        setInvisible(localStorage.getItem("invisibleNotification"));
      }).catch((error) => {
        //console.log(error);
      });
    }
  };

  const dropDowMenuForNotifications = (
    <Popper
      open={openNotifications}
      anchorEl={anchorRef.current}
      role={undefined}
      transition
      disablePortal
      style={{ width: "30%", zIndex: "1" }}
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
                autoFocusItem={openNotifications}
                id="menu-list-grow"
                onKeyDown={handleListKeyDown}
              >
                {allNotifications.length === 0 && <p>aaaaaa</p>}
                {allNotifications.map((notification) => (
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
                      <Grid item xs={7}>
                        <Link
                          to={
                            notification.post_id !== undefined &&
                            notification.post_id !== null
                              ? "/dialog/" + `${notification.post_id}`
                              : "/homePage/" + `${notification.user_who_follow}`
                          }
                          style={{ textDecoration: "none", color: "black" }}
                        >
                          {(notification.content === "started following you." ||
                            notification.content ===
                              "requested to following you." ||
                            notification.content === "liked your photo." ||
                            notification.content === "disliked your photo." ||
                            notification.content ===
                              "tagged you in a post.") && (
                            <div style={{ width: "100%" }}>
                              {notification.user_who_follow +
                                " " +
                                notification.content}
                            </div>
                          )}
                          {notification.content === "commented your post:" && (
                            <div style={{ width: "100%" }}>
                              {notification.user_who_follow +
                                " " +
                                notification.content +
                                " " +
                                notification.comment}
                            </div>
                          )}
                        </Link>
                      </Grid>
                      <Grid item xs={2}>
                        {notification.post_id !== undefined &&
                          notification.post_id !== null && (
                            <img
                              width="100%"
                              height="100%"
                              src={`http://localhost:8080/api/media/get-media-image/${notification.media}`}
                            />
                          )}
                      </Grid>
                    </Grid>
                  </MenuItem>
                ))}
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

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
        renderOption={(option) => (
          <Grid container>
            <Grid item xs={2}>
              {isHastag && (
                <Avatar
                  alt="#"
                  style={{
                    backgroundColor: "#ECECEC",
                    border: "1px solid black",
                    color: "black",
                  }}
                >
                  #
                </Avatar>
              )}
              {isUser && option.Username !== undefined && (
                <Avatar
                  alt="N"
                  src={avatar}
                  style={{ border: "1px solid" }}
                ></Avatar>
              )}
              {isUser && option.Username === undefined && (
                <Avatar
                  alt="N"
                  style={{
                    backgroundColor: "#ECECEC",
                    border: "1px solid black",
                    color: "black",
                  }}
                >
                  <RoomRounded />
                </Avatar>
              )}
            </Grid>
            <Grid item xs={10} style={{ marginTop: "3%" }}>
              {option.Username !== undefined ? option.Username : option}
            </Grid>
          </Grid>
        )}
        options={
          searchedContent !== null && searchedContent.length !== 0
            ? searchedContent.map((o) => o)
            : []
        }
        getOptionLabel={(option) =>
          option.Username !== undefined ? option.Username : option
        }
        onChange={(event, value) => goToSearchContent(value)}
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
        {username !== "admin" && (
          <Grid item xs={8} container style={{ textAlign: "right" }}>
            {searchBar}
            <Grid item xs={2} style={{ margin: "auto" }}>
              <Link to="/">
                <HomeOutlined
                  style={{
                    color: "gray",
                    width: "33px",
                    height: "33px",
                    marginRight: "5%",
                    cursor: "pointer",
                  }}
                />
              </Link>
              <Link to="/messages/">
              <EmailOutlined
                style={{
                  color: "gray",
                  width: "29px",
                  height: "29px",
                }}
              />
              </Link>
            </Grid>
            <Grid
              container
              item
              xs={2}
              style={{ textAlign: "left", margin: "auto" }}
            >
              <Link to="/follow-suggestions/">
                <ExploreOutlined
                  style={{
                    color: "gray",
                    width: "29px",
                    height: "33px",
                    marginLeft: "5%",
                    cursor: "pointer",
                  }}
                />
              </Link>
              <Badge color="secondary" variant="dot" invisible={invisible}>
                <FavoriteBorderOutlined
                  style={{
                    color: "gray",
                    width: "29px",
                    height: "33px",
                    marginLeft: "5%",
                    cursor: "pointer",
                  }}
                  ref={anchorRef}
                  aria-controls={
                    openNotifications ? "menu-list-grow" : undefined
                  }
                  aria-haspopup="true"
                  onClick={handleNotificationButton}
                />
              </Badge>
              {dropDowMenuForNotifications}

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
        )}

        {username === "admin" && (
          <Grid item xs={8} container style={{ textAlign: "right" }}>
            <Grid item xs={8} />
            <Grid item xs={2}>
              <Button variant="text" onClick={logout}>
                <a
                  href={"/"}
                  style={{
                    textDecoration: "none",
                    color: "red",
                    width: "100%",
                  }}
                >
                  Logout
                </a>
              </Button>
            </Grid>
          </Grid>
        )}
      </Grid>
    </Toolbar>
  );
  return (
    <>
      {redirection === true && <Redirect to={redirectionString} />}
      <AppBar position="static">
        {(username === null || username === undefined) &&
          NavBarForUnregisteredUser}
        {username !== null && username !== undefined && NavBarForRegistredUser}
      </AppBar>
    </>
  );
};
export default NavBar;
