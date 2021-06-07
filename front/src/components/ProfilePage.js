import { makeStyles, withStyles } from "@material-ui/core/styles";
import React, { useState, useEffect } from "react";
import { AppBar, Tabs, Tab, Grid, Button, TextField } from "@material-ui/core";
import cloneDeep from "lodash/cloneDeep";
import Box from "@material-ui/core/Box";
import Avatar from "@material-ui/core/Avatar";
import { deepOrange, deepPurple } from "@material-ui/core/colors";
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

const ProfilePage = () => {
  const classes = useStyles();
  const [selectedValue, setSelectedValue] = React.useState("a");
  const handleChange = (event) => {
    setSelectedValue(event.target.value);
    setUser({ ...user, Gender: selectedValue });
  };
  const [user, setUser] = useState({});
  const [userCopy, setUserCopy] = useState({});

  useEffect(() => {
    axios.get("http://localhost:8080/user/peraaa").then((res) => {
      setUser(res.data);
      setUserCopy(res.data);
      res.data.Gender === 0
        ? setSelectedValue("male")
        : setSelectedValue("female");
    });
  }, []);

  const handleClickSubmit = () => {
      var userDto = {
        Username:user.Username,
        FirstName:user.FirstName,
        LastName:user.LastName,
        DateOfBirth: user.DateOfBirth,
        Email:user.Email,
        PhoneNumber:user.PhoneNumber,
        Gender:user.Gender,
        Biography:user.Biography,
        WebSite:user.WebSite
      }
      axios.put("http://localhost:8080/update/peraaa", userDto)
       .then((res) => {
           console.log('uspesno')
       })
  }

  return (
    <Grid container item xs={9} style={{ height: 600 }}>
      <Grid container item xs={12}>
        <Grid item xs={2}>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            <Avatar className={classes.orange}>N</Avatar>
          </Grid>

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
              <p style={{ textAlign: "left", margin: 0 }}>profile_name</p>
              <p style={{ textAlign: "left", margin: 0 }}>
                {" "}
                <Button style={{ fontSize: 12 }} color="primary">
                  Change profile photo
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
                onChange={(e) => setUser({ ...user, Username: e.target.value })}
              />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField
                fullWidth
                variant="outlined"
                size="small"
                value={user.WebSite}
                onChange={(e) => setUser({ ...user, WebSite: e.target.value })}
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
                defaultValue="top"
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
              <Button onClick={handleClickSubmit}>Submit</Button>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default ProfilePage;