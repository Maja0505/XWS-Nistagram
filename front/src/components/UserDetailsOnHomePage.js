import { Grid, Button, Typography, FormLabel } from "@material-ui/core";
import { Link } from "react-router-dom";
import { useState } from "react";
import avatar from "../images/nistagramAvatar.jpg";
import axios from "axios";
import UsersList from "./UsersList";

import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import UserBlockMuteCloseDialog from "./UserBlockMuteCloseDialog";
import FollowRequest from "./FollowRequests";
import ViewCampaignRequestsForInfluencerDialog from "./ViewCampaignRequestsForInfluencerDialog";
import AddCampaignToInfluencerDialog from "./AddCampaignToInfluencerDialog";
import { connect, sendMsg } from "../api/index";

const UserDetailsOnHomePage = ({
  user,
  allFollows,
  allFollowers,
  relationShip,
  setRelationShip,
  setAllFollowers,
}) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const loggedInId = localStorage.getItem("id");
  const loggedInUsername = localStorage.getItem("username");
  const loggedIsAgent = localStorage.getItem("isAgent");

  const [openDialogForFollows, setOpenDialogForFollows] = useState(false);
  const [openDialogForFollowers, setOpenDialogForFollowers] = useState(false);
  const [openDialogForMore, setOpenDialogForMore] = useState(false);
  const [openDialogForFollowRequests, setOpenDialogForFollowRequests] =
    useState(false);
  const [
    openDialogForViewInfluencerCampaignRequests,
    setOpenDialogForViewInfluencerCampaignRequests,
  ] = useState(false);
  const [
    openDialogForAddCamapignToInfluencer,
    setOpenDialogForAddCamapignToInfluencer,
  ] = useState(false);

  const followUser = () => {
    if (user.ProfileSettings.Public) {
      setRelationShip({ ...relationShip, IsRequested: true });
      var follow = {
        User: loggedInId,
        FollowedUser: user.ID,
        Private: true,
      };
      axios
        .post("/api/user-follow/followUser", follow, authorization)
        .then((res) => {
          sendMsg('{"user_from":' +
          '"' +
          loggedInUsername +
          '"' +
          ',"command": 3, "channel": ' +
          '"' +
          user.IdString +
          '"' +
          ', "content": "requested to following you."}')
          console.log("uspesno");
        })
        .catch((error) => {
          console.log(error);
        });
    } else {
      setRelationShip({ ...relationShip, IsFollowing: true });
      setAllFollowers((prevState) => [
        ...prevState,
        { IdString: loggedInId, Username: loggedInUsername },
      ]);
      var follow = {
        User: loggedInId,
        FollowedUser: user.ID,
        Private: false,
      };
      axios
        .post("/api/user-follow/followUser", follow, authorization)
        .then((res) => {
          sendMsg('{"user_from":' +
          '"' +
          loggedInUsername +
          '"' +
          ',"command": 3, "channel": ' +
          '"' +
          user.IdString +
          '"' +
          ', "content": "started following you."}')
          console.log("uspesno");
        })
        .catch((error) => {
          console.log(error);
        });
    }
  };

  const unfollowUser = () => {
    setRelationShip({ ...relationShip, IsFollowing: false });
    deleteFromArray();
    var follow = {
      User: loggedInId,
      UnfollowedUser: user.ID,
    };
    axios
      .put("/api/user-follow/unfollowUser", follow, authorization)
      .then((res) => {
        console.log("uspesno");
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const unsendFollowRequest = () => {
    setRelationShip({ ...relationShip, IsRequested: false });
    var requestDto = {
      User: user.ID,
      UserWitchSendRequest: loggedInId,
    };
    axios
      .put("/api/user-follow/cancelFollowRequest", requestDto, authorization)
      .then((res) => {
        console.log("uspelo");
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const deleteFromArray = () => {
    var array = [...allFollowers];
    let index = allFollowers.findIndex((item) => item.IdString === loggedInId);
    console.log(index);
    if (index !== -1) {
      array.splice(index, 1);
    }
    setAllFollowers(array);
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
      style={{ marginLeft: "5%" }}
      onClick={unfollowUser}
    >
      Following
    </Button>
  );

  const buttonForFollow = (
    <Button
      variant="contained"
      color="primary"
      style={{ marginLeft: "5%" }}
      onClick={followUser}
    >
      Follow
    </Button>
  );

  const buttonForRequested = (
    <Button
      variant="contained"
      color="primary"
      style={{
        marginLeft: "5%",
        backgroundColor: "whitesmoke",
        color: "darkgray",
      }}
      onClick={unsendFollowRequest}
    >
      Requested
    </Button>
  );

  const buttonForEditProfile = (
    <Button variant="outlined" color="inherit" style={{ marginLeft: "5%" }}>
      <Link
        to="/accounts/edit/"
        style={{ textDecoration: "none", color: "gray" }}
      >
        Edit profile
      </Link>
    </Button>
  );

  const buttonForMore = (
    <Button
      variant="text"
      style={{ marginLeft: "2%" }}
      onClick={() => setOpenDialogForMore(true)}
    >
      <MoreHorizIcon></MoreHorizIcon>
    </Button>
  );

  const buttonForViewFollowRequests = (
    <Button
      variant="text"
      style={{ marginLeft: "2%" }}
      onClick={() => setOpenDialogForFollowRequests(true)}
    >
      View follow request
    </Button>
  );

  const buttonForViewCampaignRequest = (
    <Button
      variant="text"
      color="primary"
      style={{ marginLeft: "2%" }}
      onClick={(e) => setOpenDialogForViewInfluencerCampaignRequests(true)}
    >
      Campaign requests
    </Button>
  );

  const buttonForSendCamaignRequest = (
    <Button
      variant="text"
      color="primary"
      style={{ marginLeft: "2%" }}
      onClick={(e) => setOpenDialogForAddCamapignToInfluencer(true)}
    >
      Send campaign
    </Button>
  );

  return (
    <div>
      {allFollows !== null && allFollowers !== null && (
        <Grid container>
          <Grid item xs={4}>
            {user.ProfilePicture !== "" && (
              <img
                src={
                  "http://localhost:8080/api/media/get-profile-picture/" +
                  user.ProfilePicture
                }
                alt="Not founded"
                style={{
                  borderRadius: "50%",
                  border: "1px solid",
                  width: "130px",
                  height: "130px",
                }}
              />
            )}
            {user.ProfilePicture === "" && (
              <img
                src={avatar}
                alt="Not founded"
                style={{
                  borderRadius: "50%",
                  border: "1px solid",
                  width: "130px",
                  height: "130px",
                }}
              />
            )}
          </Grid>

          <Grid item xs={8}>
            {relationShip === null && (
              <Grid container>
                <Typography style={{ fontSize: "25px", fontStyle: "italic" }}>
                  {user.Username}
                </Typography>
                {buttonForEditProfile}
                {user.ProfileSettings.Public && buttonForViewFollowRequests}
                {user.VerificationSettings.Category === 0 &&
                  buttonForViewCampaignRequest}
              </Grid>
            )}

            {relationShip !== null && (
              <Grid container>
                <Typography style={{ fontSize: "25px", fontStyle: "italic" }}>
                  {user.Username}
                </Typography>
                {relationShip.IsFollowing &&
                  !relationShip.IsRequested &&
                  buttonForUnfollow}
                {!relationShip.IsFollowing &&
                  !relationShip.IsRequested &&
                  !relationShip.IsBlocked &&
                  buttonForFollow}
                {relationShip.IsBlocked && (
                  <Typography
                    style={{
                      fontSize: "20px",
                      marginLeft: "3%",
                      color: "red",
                      maxWidth: "200px",
                    }}
                  >
                    {`is blocked, need to unblock\nfor interaction`}
                  </Typography>
                )}
                {relationShip.IsRequested && buttonForRequested}
                {loggedIsAgent === "true" &&
                  user.VerificationSettings.Category === 0 &&
                  !relationShip.IsBlocked &&
                  buttonForSendCamaignRequest}
                {buttonForMore}
              </Grid>
            )}

            <Grid container style={{ marginTop: "3.5%" }}>
              <Grid item xs={6} style={{ textAlign: "left" }}>
                <FormLabel
                  style={{
                    marginRight: "5%",
                    cursor: "pointer",
                    fontSize: "18px",
                  }}
                  onClick={handleClickOnFollowers}
                >
                  <b>{allFollowers.length}</b> followers
                </FormLabel>
                <FormLabel
                  style={{
                    marginLeft: "5%",
                    cursor: "pointer",
                    fontSize: "18px",
                  }}
                  onClick={handleClickOnFollows}
                >
                  <b>{allFollows.length}</b> following
                </FormLabel>
              </Grid>
            </Grid>

            <Grid container style={{ marginTop: "3.5%" }}>
              <Grid item xs={6} style={{ textAlign: "left" }}>
                <label
                  style={{
                    marginRight: "5%",
                    fontSize: "20px",
                  }}
                >
                  <b>
                    {user.FirstName} {user.LastName}
                  </b>
                </label>
                <br></br>
                <div style={{ maxHeight: "150px", overflow: "auto" }}>
                  <label>{user.Biography}</label>
                </div>
                <a href={user.WebSite}>{user.WebSite}</a>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
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
          loggedUserId={loggedInId}
          open={openDialogForFollowRequests}
          setOpen={setOpenDialogForFollowRequests}
          setAllFollowers={setAllFollowers}
        ></FollowRequest>
      )}

      {openDialogForMore && (
        <UserBlockMuteCloseDialog
          open={openDialogForMore}
          setOpen={setOpenDialogForMore}
          user={user}
          relationShip={relationShip}
          setRelationShip={setRelationShip}
          allFollowers={allFollowers}
          setAllFollowers={setAllFollowers}
        ></UserBlockMuteCloseDialog>
      )}

      {openDialogForViewInfluencerCampaignRequests && (
        <ViewCampaignRequestsForInfluencerDialog
          label={"Camapign requests"}
          open={openDialogForViewInfluencerCampaignRequests}
          setOpen={setOpenDialogForViewInfluencerCampaignRequests}
        ></ViewCampaignRequestsForInfluencerDialog>
      )}

      {openDialogForAddCamapignToInfluencer && (
        <AddCampaignToInfluencerDialog
          label={"Send campaign request to influencer"}
          influencer={user.IdString}
          agent={loggedInId}
          open={openDialogForAddCamapignToInfluencer}
          setOpen={setOpenDialogForAddCamapignToInfluencer}
        ></AddCampaignToInfluencerDialog>
      )}
    </div>
  );
};

export default UserDetailsOnHomePage;
