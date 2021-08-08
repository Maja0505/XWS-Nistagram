import { Grid, Button, TextField, Divider } from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import React, { useState, useEffect } from "react";
import axios from "axios";
import DialogForListBlockUser from "./DialogForListBlockUser";
import DialogForListMuteUser from "./DialogForListMuteUser";

const ProfilePrivacy = ({ profileSettings, setProfileSettings, load }) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const username = localStorage.getItem("username");
  const loggedUserId = localStorage.getItem("id");

  const [openDialogForListBlockUsers, setOpenDialogForListBlockUsers] =
    useState(false);
  const [openDialogForListMuteUsers, setOpenDialogForListMuteUsers] =
    useState(false);

  const handleClickViewAllBlock = () => {
    setOpenDialogForListBlockUsers(true);
  };

  const handleClickViewAllMute = () => {
    setOpenDialogForListMuteUsers(true);
  };

  useEffect(() => {}, [
    profileSettings.Public,
    profileSettings.MessageRequest,
    profileSettings.AllowTags,
  ]);

  const HandleOnChangeAccountPrivacy = () => {
    if (profileSettings.Public) {
      console.log(profileSettings.Public);
      axios
        .put(
          "/api/user/" + username + "/public-profile/false",
          {},
          authorization
        )
        .then((res) => {
          setProfileSettings({ ...profileSettings, Public: false });
        });
    } else {
      console.log(profileSettings.Public);

      axios
        .put(
          "/api/user/" + username + "/public-profile/true",
          {},
          authorization
        )
        .then((res) => {
          setProfileSettings({ ...profileSettings, Public: true });
        })
        .catch((error) => {
          //console.log(error);
        });
    }
  };

  const HandleOnChangeMessageRequest = () => {
    if (profileSettings.MessageRequest) {
      axios
        .put(
          "/api/user/" + username + "/message-request/false",
          {},
          authorization
        )
        .then((res) => {
          setProfileSettings({ ...profileSettings, MessageRequest: false });
        });
    } else {
      axios
        .put(
          "/api/user/" + username + "/message-request/true",
          {},
          authorization
        )
        .then((res) => {
          setProfileSettings({ ...profileSettings, MessageRequest: true });
        })
        .catch((error) => {
          //console.log(error);
        });
    }
  };

  const HandleOnChangeAllowTags = () => {
    if (profileSettings.AllowTags) {
      axios
        .put("/api/user/" + username + "/allow-tags/false", {}, authorization)
        .then((res) => {
          setProfileSettings({ ...profileSettings, AllowTags: false });
        });
    } else {
      axios
        .put("/api/user/" + username + "/allow-tags/true", {}, authorization)
        .then((res) => {
          setProfileSettings({ ...profileSettings, AllowTags: true });
        })
        .catch((error) => {
          //console.log(error);
        });
    }
  };
  return (
    <Grid container item xs={9} style={{ height: 600 }}>
      <Grid item xs={1}></Grid>
      {load && profileSettings !== undefined && (
        <Grid container item xs={10}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 25, textAlign: "left" }}>Account Privacy</p>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Checkbox
                      name="checkedF"
                      checked={profileSettings.Public === true}
                      onChange={HandleOnChangeAccountPrivacy}
                    />
                  }
                  label="Private Account"
                  style={{ fontSize: 15, fontWeight: "bold" }}
                />
              </FormGroup>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 13, textAlign: "left" }}>
                When your account is private, only people you approve can see
                your photos and videos on Instagram. Your existing followers
                won't be affected.
              </p>
            </Grid>
            <Divider />
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 25, textAlign: "left" }}>Message Request</p>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Checkbox
                      name="checkedF"
                      checked={profileSettings.MessageRequest === true}
                      onChange={HandleOnChangeMessageRequest}
                    />
                  }
                  label="Allow message request"
                  style={{ fontSize: 15, fontWeight: "bold" }}
                />
              </FormGroup>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 13, textAlign: "left" }}>
                Let people send you message request.
              </p>
            </Grid>
            <Divider />
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 25, textAlign: "left" }}>Tag settings</p>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Checkbox
                      name="checkedF"
                      checked={profileSettings.AllowTags === true}
                      onChange={HandleOnChangeAllowTags}
                    />
                  }
                  label="Allow tags"
                  style={{ fontSize: 15, fontWeight: "bold" }}
                />
              </FormGroup>
            </Grid>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 13, textAlign: "left" }}>
                Let people tag you on posts,stories and albums.
              </p>
            </Grid>
            <Divider />
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <Grid style={{ height: "30%", width: "100%" }}>
              <p style={{ fontSize: 25, textAlign: "left" }}>Connections</p>
            </Grid>
            <Grid container style={{ height: "30%", width: "100%" }}>
              <Grid container style={{ height: "50%", width: "100%" }}>
                <Grid item xs={3}>
                  All blocked users
                </Grid>
                <Grid item xs={3}>
                  <Button onClick={handleClickViewAllBlock}>View all</Button>
                </Grid>
              </Grid>
              <Grid container style={{ height: "50%", width: "100%" }}>
                <Grid item xs={3}>
                  All muted users
                </Grid>
                <Grid item xs={3}>
                  <Button onClick={handleClickViewAllMute}>View all</Button>
                </Grid>
              </Grid>
            </Grid>
            <Divider />
          </Grid>
        </Grid>
      )}
      <Grid item xs={1}></Grid>
      {openDialogForListBlockUsers && (
        <DialogForListBlockUser
          loggedUserId={loggedUserId}
          open={openDialogForListBlockUsers}
          setOpen={setOpenDialogForListBlockUsers}
        ></DialogForListBlockUser>
      )}
      {openDialogForListMuteUsers && (
        <DialogForListMuteUser
          loggedUserId={loggedUserId}
          open={openDialogForListMuteUsers}
          setOpen={setOpenDialogForListMuteUsers}
        ></DialogForListMuteUser>
      )}
    </Grid>
  );
};

export default ProfilePrivacy;
