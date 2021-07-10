import { Grid, Paper, Avatar, Button } from "@material-ui/core";
import { useState, useEffect } from "react";

import { Link } from "react-router-dom";

import avatar from "../images/nistagramAvatar.jpg";

import axios from "axios";
const FollowSuggestionOne = ({ suggestion }) => {
  const [user, setUser] = useState();
  const [allFollowers, setAllFollowers] = useState([]);
  const [buttonState, setButtonState] = useState("not_following");
  const loggedUserId = localStorage.getItem("id");
  const [requested, setRequested] = useState(false);

  useEffect(() => {
    axios.get("/api/user/" + suggestion.Username).then((res) => {
      console.log(res.data);
      setUser(res.data);
    });

    axios
      .get("/api/user-follow/allFollowers/" + suggestion.IdString)
      .then((res) => {
        if (res.data) {
          console.log(res.data);
          setAllFollowers(res.data);
        } else {
          setAllFollowers([]);
        }
      })
      .catch((error) => {
        setAllFollowers([]);
      });

    axios
      .get(
        "/api/user-follow/checkRequested/" +
          loggedUserId +
          "/" +
          suggestion.IdString
      )
      .then((res) => {
        setRequested(res.data);
      }).catch((error) => {
        //console.log(error);
      });
  }, []);

  const followClicked = () => {
    if (user.ProfileSettings.Public === false) {
      setButtonState("following");
      setAllFollowers([
        ...allFollowers,
        { IdString: user.IdString, Username: user.Username },
      ]);
      var follow = {
        User: loggedUserId,
        FollowedUser: user.ID,
        Private: false,
      };
      axios.post("/api/user-follow/followUser", follow).then((res) => {
        console.log("uspesno");
      });
    } else {
      var follow = {
        User: loggedUserId,
        FollowedUser: user.ID,
        Private: true,
      };
      axios.post("/api/user-follow/followUser", follow).then((res) => {
        console.log("uspesno");
      }).catch((error) => {
        //console.log(error);
      });
      setButtonState("requested");
    }
  };

  const unfollowClicked = () => {
    setButtonState("not_following");
    var array = [...allFollowers];
    array.pop();
    setAllFollowers(array);
    var follow = {
      User: loggedUserId,
      UnfollowedUser: user.ID,
    };
    axios.put("/api/user-follow/unfollowUser", follow).then((res) => {
      console.log("uspesno");
    }).catch((error) => {
      //console.log(error);
    });
  };

  const requestedClicked = () => {
    setButtonState("not_following");
    setRequested(false);
    var requestDto = {
      User: user.ID,
      UserWitchSendRequest: loggedUserId,
    };
    axios
      .put("/api/user-follow/cancelFollowRequest", requestDto)
      .then((res) => {
        console.log("uspelo");
      }).catch((error) => {
        //console.log(error);
      });
  };

  const buttonForUnfollow = (
    <Button variant="contained" color="default" onClick={unfollowClicked}>
      Following
    </Button>
  );

  const buttonForFollow = (
    <Button variant="contained" color="primary" onClick={followClicked}>
      Follow
    </Button>
  );

  const buttonForRequested = (
    <Button variant="text" color="inherit" onClick={requestedClicked}>
      Requested
    </Button>
  );

  return (
    <div>
      {user !== null && user !== undefined && (
        <Paper elevation={15} style={{ marginTop: "1%" }}>
          <Grid container style={{ marginBottom: "3%" }}></Grid>
          <Grid container>
            <Grid item xs={3} style={{ margin: "auto" }}>
              {user.profilePicture !== null &&
              user.profilePicture !== undefined ? (
                <Avatar
                  alt="N"
                  src={
                    user.profilePicture === ""
                      ? avatar
                      : "http://localhost:8080/api/media/get-profile-picture/" +
                        user.profilePicture
                  }
                  style={{
                    border: "0.5px solid",
                    margin: "auto",
                    width: "50px",
                    height: "50px",
                    borderColor: "#b9b9b9",
                  }}
                ></Avatar>
              ) : (
                <Avatar
                  alt="N"
                  src={avatar}
                  style={{
                    border: "0.5px solid",
                    margin: "auto",
                    width: "50px",
                    height: "50px",
                    borderColor: "#b9b9b9",
                  }}
                ></Avatar>
              )}
            </Grid>

            <Grid container item xs={6}>
              <Grid container>
                <Grid item xs={12} style={{ textAlign: "left" }}>
                  <label>
                    <Link
                      to={"/homePage/" + user.Username}
                      style={{ textDecoration: "none", color: "black" }}
                    >
                      <b>{user.Username}</b>
                    </Link>
                  </label>
                </Grid>
              </Grid>

              <Grid container>
                <Grid
                  item
                  xs={12}
                  style={{ textAlign: "left", marginLeft: "1%" }}
                >
                  <label>
                    {user.FirstName} {user.LastName}
                  </label>
                </Grid>
              </Grid>

              <Grid container>
                <Grid
                  item
                  xs={12}
                  style={{ textAlign: "left", marginLeft: "1%" }}
                >
                  <label>Followed by {allFollowers.length} users</label>
                </Grid>
              </Grid>
            </Grid>

            <Grid item xs={3} style={{ margin: "auto" }}>
              {buttonState === "not_following" &&
                requested === false &&
                buttonForFollow}
              {buttonState === "following" &&
                requested === false &&
                buttonForUnfollow}
              {(buttonState === "requested" || requested === true) &&
                buttonForRequested}
            </Grid>
          </Grid>
          <Grid container style={{ marginTop: "3%" }}></Grid>
        </Paper>
      )}
    </div>
  );
};

export default FollowSuggestionOne;
