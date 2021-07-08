import {
  Grid,
  Typography,
  Button,
  FormLabel,
  Paper,
  Tabs,
  Tab,
  Divider,
} from "@material-ui/core";

import { useState, useEffect } from "react";
import axios from "axios";
import avatar from "../images/nistagramAvatar.jpg";
import highlights from "../images/highlights.jpg";

import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import Posts from "./Posts";
import verification from "../images/verification.png";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import { Grow, Popper, MenuItem, MenuList } from "@material-ui/core";
import { useRef } from "react";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";
import DialogForBlockUser from "./DialogForBlockUser";
import DialogForMuteUser from "./DialogForMuteUser";
import FollowRequest from "./FollowRequests";
import Collections from "./Collections.js";
import {
  GridOn,
  BookmarkBorder,
  AssignmentIndOutlined,
  AddPhotoAlternateOutlined,
} from "@material-ui/icons";
import UsersList from "./UsersList";
import AddPost from "./AddPost";
import Story from "./Story";
import PostWhereUserTagged from "./PostWhereUserTagged.js";

const UserHomePage = () => {
  const [user, setUser] = useState();
  const [tabValue, setTabValue] = useState(0);
  const { username } = useParams();
  const [following, setFollowing] = useState(false);
  const [muted, setMuted] = useState(false);
  const [close, setClose] = useState(false);
  const [redirection, setRedirection] = useState(false);
  const [requested, setRequested] = useState(false);
  const [privateProfile, setPrivateProfile] = useState(false);

  const loggedUsername = localStorage.getItem("username");
  const loggedInId = localStorage.getItem("id");
  const [load1, setLoad1] = useState(false);
  const [load2, setLoad2] = useState(false);
  const [load3, setLoad3] = useState(false);
  const [open, setOpen] = useState(false);
  const anchorRef = useRef(null);
  const [openDialogForBlock, setOpenDialogForBlock] = useState(false);
  const [openDialogForMute, setOpenDialogForMute] = useState(false);
  const [allFollows, setAllFollows] = useState([]);
  const [allFollowers, setAllFollowers] = useState([]);
  const [openDialogForFollows, setOpenDialogForFollows] = useState(false);
  const [openDialogForFollowers, setOpenDialogForFollowers] = useState(false);
  const [openDialogForFollowRequests, setOpenDialogForFollowRequests] =
    useState(false);
  const [openHighlightsDialog, setOpenHighlightsDialog] = useState(false);
  const [highlightStories, setHighlightStories] = useState([]);

  const loggedUserId = localStorage.getItem("id");

  const clickShowFollowRequests = () => {
    setOpenDialogForFollowRequests(true);
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

  useEffect(() => {
    console.log(username);
    console.log(loggedUsername);

    axios
      .get("/api/user/" + username)
      .then((res) => {
        console.log(res.data);
        setUser(res.data);
        setPrivateProfile(res.data.ProfileSettings.Public);
        if (res.data.IdString !== loggedInId) {
          axios
            .get(
              "/api/user-follow/checkBlock/" +
                loggedInId +
                "/" +
                res.data.IdString
            )
            .then((res) => {
              console.log(res.data);

              setRedirection(res.data);
              setLoad1(true);
            })
            .catch((error) => {
              alert(error);
            });

          axios
            .get(
              "/api/user-follow/checkRequested/" +
                loggedInId +
                "/" +
                res.data.IdString
            )
            .then((res) => {
              console.log(res.data);
              setRequested(res.data);
              setLoad2(true);
            })
            .catch((error) => {
              alert(error);
            });

          axios
            .get(
              "/api/user-follow/checkFollowing/" +
                loggedInId +
                "/" +
                res.data.IdString
            )
            .then((res) => {
              console.log(res.data);
              setFollowing(res.data);
              setLoad3(true);
            })
            .catch((error) => {
              alert(error);
            });

          axios
            .get(
              "/api/user-follow/checkMuted/" +
                loggedInId +
                "/" +
                res.data.IdString
            )
            .then((res) => {
              console.log(res.data);
              setMuted(res.data);
              setLoad3(true);
            })
            .catch((error) => {
              alert(error);
            });

          axios
            .get(
              "/api/user-follow/checkClosed/" +
                loggedInId +
                "/" +
                res.data.IdString
            )
            .then((res) => {
              console.log(res.data);
              setClose(res.data);
              setLoad3(true);
            })
            .catch((error) => {
              alert(error);
            });

          axios
            .get("/api/user-follow/allFollows/" + res.data.IdString)
            .then((res) => {
              if (res.data) {
                setAllFollows(res.data);
              } else {
                setAllFollows([]);
              }
            })
            .catch((error) => {
              setAllFollows([]);
              alert(error);
            });

          axios
            .get("/api/user-follow/allFollowers/" + res.data.IdString)
            .then((res) => {
              if (res.data) {
                setAllFollowers(res.data);
              } else {
                setAllFollowers([]);
              }
            })
            .catch((error) => {
              setAllFollowers([]);
              alert(error);
            });
        } else {
          axios
            .get("/api/user-follow/allFollows/" + res.data.IdString)
            .then((res) => {
              if (res.data) {
                setAllFollows(res.data);
              } else {
                setAllFollows([]);
              }
            })
            .catch((error) => {
              setAllFollows([]);
              alert(error);
            });

          axios
            .get("/api/user-follow/allFollowers/" + res.data.IdString)
            .then((res) => {
              if (res.data) {
                setAllFollowers(res.data);
              } else {
                setAllFollowers([]);
              }
            })
            .catch((error) => {
              setAllFollowers([]);
              alert(error);
            });
          setLoad1(true);
          setLoad2(true);
          setLoad3(true);
        }
        axios
          .get("/api/post/story/all-highlights/" + res.data.IdString)
          .then((res) => {
            if (res.data) {
              setHighlightStories(res.data);
            }
          });
      })
      .catch((error) => {
        console.log(error);
      });
  }, [username, loggedUsername]);

  const handleClickUnmute = () => {
    var muteDto = {
      User: loggedUserId,
      Friend: user.ID,
      Mute: false,
    };
    axios.put("/api/user-follow/setMuteFriend", muteDto).then((res) => {
      console.log("uspelo");
      setOpen((prevOpen) => !prevOpen);
      setMuted(false);
    });
  };

  const handleOpenDialogForBlock = () => {
    setOpenDialogForBlock(true);
    setOpen((prevOpen) => !prevOpen);
  };

  const handleOpenDialogForMute = () => {
    setOpenDialogForMute(true);
    setOpen((prevOpen) => !prevOpen);
  };

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  const requestedClicked = () => {
    setRequested(!requested);
  };

  const handleSetToClose = () => {
    var closeDto = {
      User: loggedUserId,
      Friend: user.IdString,
      Close: true,
    };
    axios.put("/api/user-follow/setCloseFriend", closeDto).then((res) => {
      console.log("uspesno");
      setClose(true);
    });
  };

  const handleSetToRemoveFromClose = () => {
    var closeDto = {
      User: loggedUserId,
      Friend: user.IdString,
      Close: false,
    };
    axios.put("/api/user-follow/setCloseFriend", closeDto).then((res) => {
      console.log("uspesno");
      setClose(false);
    });
  };

  const followClicked = () => {
    if (privateProfile) {
      var follow = {
        User: loggedInId,
        FollowedUser: user.ID,
        Private: true,
      };
      axios.post("/api/user-follow/followUser", follow).then((res) => {
        console.log("uspesno");
        let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + user.IdString + '"' + ', "content": "requested to following you."}')
        };
        setRequested(true);
        

      });
    } else {
      var follow = {
        User: loggedInId,
        FollowedUser: user.ID,
        Private: false,
      };
      axios.post("/api/user-follow/followUser", follow).then((res) => {
        let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + user.IdString + '"' + ', "content": "started following you."}')
        };
        setFollowing(true);
      });
    }
  };
  const unfollowClicked = () => {
    var follow = {
      User: loggedInId,
      UnfollowedUser: user.ID,
    };
    axios.put("/api/user-follow/unfollowUser", follow).then((res) => {
      console.log("uspesno");
      //setUser({...user,allFollowers: user.allFollowers - 1})
    });
    setFollowing(false);
  };

  const handleClickOnFollowers = () => {
    setOpenDialogForFollowers(true);
  };

  const handleClickOnFollows = () => {
    setOpenDialogForFollows(true);
  };

  const buttonForUnfollow = (
    <Button
      variant="contained"
      color="default"
      style={{ margin: "auto" }}
      onClick={unfollowClicked}
    >
      Following
    </Button>
  );

  const buttonForFollow = (
    <Button
      variant="contained"
      color="primary"
      style={{ margin: "auto" }}
      onClick={followClicked}
    >
      Follow
    </Button>
  );

  const buttonForRequested = (
    <Button
      variant="contained"
      color="primary"
      style={{
        margin: "auto",
        marginLeft: "30px",
        backgroundColor: "whitesmoke",
        color: "darkgray",
      }}
      onClick={requestedClicked}
    >
      Requested
    </Button>
  );

  const buttonForEditProfile = (
    <Button variant="outlined" color="inherit" style={{ marginLeft: "auto" }}>
      <Link
        to="/accounts/edit/"
        style={{ textDecoration: "none", color: "gray" }}
      >
        Edit profile
      </Link>
    </Button>
  );
  function closeStory() {
    setOpenHighlightsDialog(false);
  }

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
                <MenuItem onClick={handleOpenDialogForBlock}>
                  <Grid container>
                    <Grid item xs={3}></Grid>
                    <Grid item xs={9}>
                      <div style={{ width: "100%" }} style={{ color: "red" }}>
                        Block this user
                      </div>
                    </Grid>
                  </Grid>
                </MenuItem>
                {following && !muted && (
                  <MenuItem onClick={handleOpenDialogForMute}>
                    <Grid container>
                      <Grid item xs={3}></Grid>
                      <Grid item xs={9}>
                        <div style={{ width: "100%" }} style={{ color: "red" }}>
                          Mute this user
                        </div>
                      </Grid>
                    </Grid>
                  </MenuItem>
                )}
                {following && !close && (
                  <MenuItem onClick={handleSetToClose}>
                    <Grid container>
                      <Grid item xs={3}></Grid>
                      <Grid item xs={9}>
                        <div style={{ width: "100%" }} style={{ color: "red" }}>
                          Set to close friends
                        </div>
                      </Grid>
                    </Grid>
                  </MenuItem>
                )}
                {following && muted && (
                  <MenuItem onClick={handleClickUnmute}>
                    <Grid container>
                      <Grid item xs={3}></Grid>
                      <Grid item xs={9}>
                        <div style={{ width: "100%" }} style={{ color: "red" }}>
                          Unmute
                        </div>
                      </Grid>
                    </Grid>
                  </MenuItem>
                )}
                {following && close && (
                  <MenuItem onClick={handleSetToRemoveFromClose}>
                    <Grid container>
                      <Grid item xs={3}></Grid>
                      <Grid item xs={9}>
                        <div style={{ width: "100%" }} style={{ color: "red" }}>
                          Remove from close friends
                        </div>
                      </Grid>
                    </Grid>
                  </MenuItem>
                )}
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

  const handleClickOpen = () => {
    setOpenHighlightsDialog(true);
  };

  const userDetails = (
    <Grid container style={{ marginTop: "3%" }}>
      <Grid container style={{ margin: "auto" }}>
        <Grid item xs={2}></Grid>
        <Grid container item xs={8}>
          <Grid item xs={4}>
            {user !== undefined && user.ProfilePicture !== "" && (
              <img
                src={
                  "http://localhost:8080/api/media/get-profile-picture/" +
                  user.ProfilePicture
                }
                alt="Not founded"
                style={{
                  borderRadius: "50%",
                  border: "1px solid",
                  width: "150px",
                  height: "150px",
                }}
              />
            )}
            {user !== undefined && user.ProfilePicture === "" && (
              <img
                src={avatar}
                alt="Not founded"
                style={{
                  borderRadius: "50%",
                  border: "1px solid",
                  width: "150px",
                  height: "150px",
                }}
              />
            )}
          </Grid>
          {user !== undefined && (
            <Grid item xs={7}>
              <Grid container>
                {user !== undefined && (
                  <>
                    <Grid
                      item
                      xs={3}
                      style={{
                        textAlign: "left",
                      }}
                    >
                      <Typography variant="h6" style={{ margin: "auto" }}>
                        {user.Username} {"  "}
                        {user.VerificationSettings.Verified && (
                          <img
                            src={verification}
                            style={{
                              height: "20px",
                              width: "20px",
                              marginTop: "2%",
                            }}
                          ></img>
                        )}
                      </Typography>
                    </Grid>
                  </>
                )}

                <Grid item xs={3}>
                  {loggedUsername === username && buttonForEditProfile}
                  {requested &&
                    loggedUsername !== username &&
                    buttonForRequested}
                  {following &&
                    loggedUsername !== username &&
                    !requested &&
                    buttonForUnfollow}
                  {!following &&
                    loggedUsername !== username &&
                    !requested &&
                    buttonForFollow}
                </Grid>

                <Grid item xs={2}></Grid>

                <Grid
                  item
                  xs={4}
                  style={{
                    textAlign: "right",
                  }}
                >
                  {loggedUsername !== username && (
                    <>
                      <MoreHorizIcon
                        style={{
                          textAlign: "right",
                          cursor: "pointer",
                        }}
                        aria-controls={open ? "menu-list-grow" : undefined}
                        aria-haspopup="true"
                        ref={anchorRef}
                        onClick={handleToggle}
                      ></MoreHorizIcon>
                      {dropDowMenuForProfile}
                    </>
                  )}

                  {loggedUsername === username && user.ProfileSettings.Public && (
                    <>
                      <Button onClick={clickShowFollowRequests}>
                        View follow request
                      </Button>
                    </>
                  )}
                </Grid>
              </Grid>
              <br></br>
              <Grid container>
                {user !== undefined && (
                  <>
                    <FormLabel>0 posts</FormLabel>
                    <FormLabel
                      style={{ marginLeft: "auto", cursor: "pointer" }}
                      onClick={handleClickOnFollowers}
                    >
                      {allFollowers.length} followers
                    </FormLabel>
                    <FormLabel
                      style={{ marginLeft: "auto", cursor: "pointer" }}
                      onClick={handleClickOnFollows}
                    >
                      {allFollows.length} following
                    </FormLabel>{" "}
                  </>
                )}
              </Grid>
              {user !== undefined && (
                <Grid container style={{ marginTop: "1%" }}>
                  <Typography variant="inherit" align="left">
                    {user.FirstName} {user.LastName}
                  </Typography>
                </Grid>
              )}
              <Grid container>
                {user !== undefined && (
                  <>
                    <Typography style={{ textAlign: "left" }}>
                      {user.Biography}
                    </Typography>
                  </>
                )}
              </Grid>
            </Grid>
          )}
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
      <Grid container style={{ marginTop: "1%" }}></Grid>
      {highlightStories.length !== 0 && !privateProfile && (
        <Grid container style={{ margin: "auto" }}>
          <Grid item xs={2}></Grid>
          <Grid container item xs={8}>
            <Grid item xs={4}>
              <div onClick={handleClickOpen}>
                <div>
                  <img
                    src={highlights}
                    style={{
                      borderRadius: "50%",
                      border: "1px solid",
                      width: "50px",
                      height: "50px",
                      cursor: "pointer",
                    }}
                  />
                </div>
                <div>{"Highlights"}</div>
              </div>
            </Grid>
            <Grid item xs={7}></Grid>
          </Grid>
          <Grid item xs={2}></Grid>
        </Grid>
      )}
    </Grid>
  );

  return (
    <>
      {user !== undefined && (
        <div>
          {<>{userDetails}</>}
          {user !== undefined && loggedUsername === username && (
            <Grid container style={{ marginTop: "2%" }}>
              <Grid item xs={2}></Grid>
              <Grid item xs={8}>
                <Paper>
                  <Tabs
                    value={tabValue}
                    onChange={handleChangeTab}
                    indicatorColor="primary"
                    textColor="inherit"
                  >
                    <Tab
                      label="Posts"
                      icon={<GridOn />}
                      style={{ margin: "auto" }}
                    />
                    <Tab
                      label="Add post"
                      icon={<AddPhotoAlternateOutlined />}
                      style={{ margin: "auto" }}
                    />

                    <Tab
                      label="Saved"
                      icon={<BookmarkBorder />}
                      style={{ margin: "auto" }}
                    />
                    <Tab
                      label="Tagged"
                      icon={<AssignmentIndOutlined />}
                      style={{ margin: "auto" }}
                    />
                  </Tabs>
                </Paper>
              </Grid>
              <Grid item xs={2}></Grid>
            </Grid>
          )}

          {loggedUsername !== username &&
            (!user.ProfileSettings.Public || following) && (
              <Grid container style={{ marginTop: "2%" }}>
                <Grid item xs={2}></Grid>
                <Grid item xs={8}>
                  <Paper>
                    <Tabs
                      value={tabValue}
                      onChange={handleChangeTab}
                      indicatorColor="primary"
                      textColor="inherit"
                    >
                      <Tab
                        label="Posts"
                        icon={<GridOn />}
                        style={{ margin: "auto" }}
                      />
                      <Tab
                        label="Tagged"
                        icon={<AssignmentIndOutlined />}
                        style={{ margin: "auto" }}
                      />
                    </Tabs>
                  </Paper>
                </Grid>
                <Grid item xs={2}></Grid>
              </Grid>
            )}

          <Grid container>
            <Grid item xs={2}></Grid>
            {user !== undefined &&
              user !== null &&
              loggedUsername === user.Username && (
                <Grid item xs={8}>
                  {user !== undefined && user !== null && tabValue === 0 && (
                    <Posts userForProfile={user} username={username}></Posts>
                  )}
                  {user !== undefined && user !== null && tabValue === 1 && (
                    <AddPost setTabValue={setTabValue} />
                  )}
                  {user !== undefined && user !== null && tabValue === 2 && (
                    <Collections></Collections>
                  )}
                  {user !== undefined && user !== null && tabValue === 3 && (
                    <PostWhereUserTagged user={user}></PostWhereUserTagged>
                  )}
                </Grid>
              )}

            {((user !== undefined &&
              user !== null &&
              loggedUsername !== user.Username &&
              !user.ProfileSettings.Public) ||
              (following && loggedUsername !== user.Username)) && (
              <Grid item xs={8}>
                {user !== undefined && user !== null && tabValue === 0 && (
                  <Posts userForProfile={user} username={username}></Posts>
                )}
                {user !== undefined && user !== null && tabValue === 1 && (
                  <PostWhereUserTagged user={user}></PostWhereUserTagged>
                )}
              </Grid>
            )}

            {user !== undefined &&
              user !== null &&
              loggedUsername !== user.Username &&
              user.ProfileSettings.Public &&
              !following && (
                <Grid item xs={8} style={{ marginTop: "2%" }}>
                  {user !== undefined && user !== null && tabValue === 0 && (
                    <Paper style={{ width: "100%", height: "100%" }}>
                      <Typography variant="h5" color="textSecondary">
                        This Account is Private
                      </Typography>
                      <p>Follow to see their photos and videos.</p>
                    </Paper>
                  )}
                </Grid>
              )}

            <Grid item xs={2}></Grid>
          </Grid>
          {user !== undefined && (
            <DialogForBlockUser
              loggedUserId={loggedUserId}
              blockedUserId={user.ID}
              open={openDialogForBlock}
              setOpen={setOpenDialogForBlock}
            ></DialogForBlockUser>
          )}
          {user !== undefined && (
            <DialogForMuteUser
              loggedUserId={loggedUserId}
              muteUserId={user.ID}
              open={openDialogForMute}
              setOpen={setOpenDialogForMute}
            ></DialogForMuteUser>
          )}

          {openDialogForFollowers && (
            <UsersList
              label="Followers"
              users={allFollowers}
              open={openDialogForFollowers}
              setOpen={setOpenDialogForFollowers}
            ></UsersList>
          )}

          {openDialogForFollows && (
            <UsersList
              label="Following"
              users={allFollows}
              open={openDialogForFollows}
              setOpen={setOpenDialogForFollows}
            ></UsersList>
          )}

          {openDialogForFollowRequests && (
            <FollowRequest
              loggedUserId={loggedUserId}
              open={openDialogForFollowRequests}
              setOpen={setOpenDialogForFollowRequests}
            ></FollowRequest>
          )}
        </div>
      )}
      <div>
        {openHighlightsDialog &&
          highlightStories !== undefined &&
          highlightStories !== null &&
          highlightStories.length !== 0 && (
            <Story
              stories={highlightStories}
              onClose={closeStory}
              user={user.IdString}
            ></Story>
          )}
      </div>
    </>
  );
};

export default UserHomePage;
