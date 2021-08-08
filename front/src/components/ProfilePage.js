import { makeStyles } from "@material-ui/core/styles";
import { useState, useEffect } from "react";
import { Grid, Button, TextField } from "@material-ui/core";
import Avatar from "@material-ui/core/Avatar";
import { deepOrange } from "@material-ui/core/colors";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import axios from "axios";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: "center",
  },
  orange: {
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
    marginLeft: "auto",
  },
}));

const ProfilePage = ({
  user,
  setUser,
  selectedValue,
  setSelectedValue,
  userCopy,
  setUserCopy,
  load,
}) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const classes = useStyles();
  const username = localStorage.getItem("username");
  const [selectedFile, setSelectedFile] = useState();
  const [image, setImage] = useState();
  const loggedUserId = localStorage.getItem("id");

  const handleChange = (event) => {
    setSelectedValue(event.target.value);
    setUser({ ...user, Gender: selectedValue });
  };

  const ChangeProfileImage = (event) => {
    setSelectedFile(null);

    console.log(user);
    setUser({ ...user, ProfilePicture: loggedUserId });
    console.log(userCopy);

    var formData = new FormData();
    console.log(event.target.files[0]);
    var file = event.target.files[0];
    formData.append("myFile", file);
    const reader = new FileReader();
    var url = reader.readAsDataURL(file);
    reader.onloadend = function (e) {
      setSelectedFile(reader.result);
    }.bind(this);

    setImage(formData);
  };

  const handleClickSubmit = () => {
    if (user.ProfilePicture !== userCopy.ProfilePicture) {
      axios
        .post("/api/media/upload-profile-picture/" + loggedUserId, image, {
          headers: { "Content-Type": "multipart/form-data" },
        })
        .then((res) => {
          console.log("Uspesno upload-ovao sliku");
        })
        .catch((error) => {
          //console.log(error);
        });
    }

    var userDto = {
      Username: user.Username,
      FirstName: user.FirstName,
      LastName: user.LastName,
      DateOfBirth: user.DateOfBirth,
      Email: user.Email,
      PhoneNumber: user.PhoneNumber,
      Gender: selectedValue === "female" ? 1 : 0,
      Biography: user.Biography,
      WebSite: user.WebSite,
      ProfilePicture: user.ProfilePicture + ".jpg",
    };
    axios
      .put("/api/user/update/" + user.Username, userDto, authorization)
      .then((res) => {
        setUserCopy({
          ...userCopy,
          FirstName: user.FirstName,
          Username: user.Username,
          WebSite: user.WebSite,
          Biography: user.Biography,
          Email: user.Email,
          PhoneNumber: user.PhoneNumber,
          Gender: user.Gender,
          ProfilePicture: user.ProfilePicture + ".jpg",
        });
        setUser({ ...user, ProfilePicture: loggedUserId + ".jpg" });
        localStorage.setItem("username", user.Username);
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  return (
    <Grid container item xs={9} style={{ height: 600 }}>
      {load && (
        <Grid container item xs={12}>
          <Grid item xs={2}>
            {userCopy !== undefined &&
              userCopy !== null &&
              userCopy.ProfilePicture !== undefined && (
                <Grid item style={{ height: "12%", textAlign: "right" }}>
                  {selectedFile && (
                    <Avatar
                      alt="N"
                      src={selectedFile}
                      style={{ marginLeft: "auto" }}
                    />
                  )}
                  {!selectedFile && userCopy.ProfilePicture !== "" && (
                    <Avatar
                      alt="N"
                      src={
                        "http://localhost:8080/api/media/get-profile-picture/" +
                        userCopy.ProfilePicture
                      }
                      style={{ marginLeft: "auto" }}
                    />
                  )}
                  {!selectedFile && userCopy.ProfilePicture === "" && (
                    <Avatar style={{ marginLeft: "auto" }}>N</Avatar>
                  )}
                </Grid>
              )}

            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Name
            </Grid>
            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Username
            </Grid>
            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Website
            </Grid>
            <Grid item style={{ height: "25%", textAlign: "right" }}>
              Bio
            </Grid>
            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Email
            </Grid>
            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Phone number
            </Grid>
            <Grid item style={{ height: "12%", textAlign: "right" }}>
              Gender
            </Grid>
          </Grid>
          <Grid container item xs={10}>
            <Grid item xs={1}></Grid>
            <Grid item xs={11}>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <p style={{ textAlign: "left", margin: 0, fontSize: 20 }}>
                  {userCopy.Username}
                </p>{" "}
                <p style={{ textAlign: "left" }}>
                  <Button
                    style={{ fontSize: 12 }}
                    color="primary"
                    component="label"
                  >
                    Change profile photo
                    <input
                      hidden
                      accept="image/*"
                      multiple
                      type="file"
                      name="myFile"
                      onChange={(event) => ChangeProfileImage(event)}
                      style={{ display: "none" }}
                    />
                  </Button>
                </p>
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  value={user.FirstName}
                  onChange={(e) =>
                    setUser({ ...user, FirstName: e.target.value })
                  }
                />
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  value={user.Username}
                  onChange={(e) =>
                    setUser({ ...user, Username: e.target.value })
                  }
                />
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  value={user.WebSite}
                  onChange={(e) =>
                    setUser({ ...user, WebSite: e.target.value })
                  }
                />
              </Grid>
              <Grid item xs={12} style={{ height: "25%", textAlign: "right" }}>
                <TextField
                  multiline
                  rows={4}
                  variant="outlined"
                  fullWidth
                  value={user.Biography}
                  onChange={(e) =>
                    setUser({ ...user, Biography: e.target.value })
                  }
                />
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  value={user.Email}
                  onChange={(e) => setUser({ ...user, Email: e.target.value })}
                />
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  value={user.PhoneNumber}
                  onChange={(e) =>
                    setUser({ ...user, PhoneNumber: e.target.value })
                  }
                />
              </Grid>
              <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
                <RadioGroup
                  row
                  aria-label="position"
                  name="position"
                  value={selectedValue}
                  onClick={handleChange}
                >
                  <FormControlLabel
                    value="male"
                    control={<Radio color="primary" />}
                    label="Male"
                  />
                  <FormControlLabel
                    value="female"
                    control={<Radio color="primary" />}
                    label="Female"
                  />
                </RadioGroup>
              </Grid>
              <Grid item style={{ height: "12%", textAlign: "left" }}>
                <Button
                  disabled={
                    (user.FirstName !== userCopy.FirstName ||
                      user.Username !== userCopy.Username ||
                      user.WebSite !== userCopy.WebSite ||
                      user.Biography !== userCopy.Biography ||
                      user.Email !== userCopy.Email ||
                      user.PhoneNumber !== userCopy.PhoneNumber ||
                      user.Gender !== userCopy.Gender ||
                      user.ProfilePicture !== userCopy.ProfilePicture) &&
                    user.Username !== ""
                      ? false
                      : true
                  }
                  onClick={handleClickSubmit}
                  color="primary"
                  variant="contained"
                >
                  Submit
                </Button>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      )}
    </Grid>
  );
};

export default ProfilePage;
