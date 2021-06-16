import {
  Grid,
  Typography,
  Button,
  FormLabel,
  Paper,
  Tabs,
  Tab,
} from "@material-ui/core";

import { useState, useEffect } from "react";
import axios from "axios";
import avatar from "../images/nistagramAvatar.jpg";
import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import Posts from "./Posts";
import verification from "../images/verification.png"

import {
  GridOn,
  BookmarkBorder,
  AssignmentIndOutlined,
} from "@material-ui/icons";

const UserHomePage = () => {
  const [user, setUser] = useState();
  const [tabValue, setTabValue] = useState(0);
  const { username } = useParams();

  const loggedUsername = localStorage.getItem("username");

  useEffect(() => {
    axios
      .get("/api/user/" + username)
      .then((res) => {
        console.log(res.data);
        setUser(res.data);
      })
      .catch((error) => {
        alert(error.response.status);
      });
  }, [username, loggedUsername]);

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  const buttonForUnFollow = (
    <Button variant="contained" color="default" style={{ margin: "auto" }}>
      Unfollow
    </Button>
  );

  const buttonForFollow = (
    <Button variant="contained" color="primary" style={{ margin: "auto" }}>
      Follow
    </Button>
  );

  const buttonForEditProfile = (
    <Button variant="outlined" color="inherit" style={{ marginLeft: "auto" }}>
      <Link to="/accounts/edit/" style={{ textDecoration: "none", color: "gray" }}>
        Edit profile
      </Link>
    </Button>
  );

  const userDetails = (
    <Grid container style={{ marginTop: "3%" }}>
      <Grid item xs={2}></Grid>
      <Grid container item xs={8}>
        <Grid item xs={4}>
          <img
            src={avatar}
            alt="Not founded"
            style={{
              borderRadius: "50%",
              border: "1px solid",
              width: "40%",
            }}
          />
        </Grid>
        <Grid item xs={7}>
          <Grid container>
            {user !== undefined && (
              <>
              <Grid item xs={8}>
              <Typography variant="h6" style={{ margin: "auto" }}>
                {user.Username} {"  "}
                {user.VerificationSettings.Verified && <img src={verification} style={{height:"20px", width:"20px", marginTop:"2%"}}></img>}
              </Typography>
              </Grid>
              <Grid item xs={3}>
               
                </Grid>
                </>
            )}

            {loggedUsername === username && buttonForEditProfile}
            {loggedUsername !== username && buttonForFollow}
          </Grid>
          <br></br>
          <Grid container>
            {user !== undefined && (
              <>
                <FormLabel>{user.NumberOfPosts} posts</FormLabel>
                <FormLabel style={{ marginLeft: "auto" }}>
                  {user.Followers} followers
                </FormLabel>
                <FormLabel style={{ marginLeft: "auto" }}>
                  {user.Following} following
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
                <Typography>{user.Biography}</Typography>
              </>
            )}
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={2}></Grid>
    </Grid>
  );

  return (
    <div>
      {userDetails}
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
              <Tab label="Posts" icon={<GridOn />} style={{ margin: "auto" }} />
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

      <Grid container>
        <Grid item xs={2}></Grid>
        <Grid item xs={8}>
          {user !== undefined && user !== null && <Posts userForProfile={user}></Posts>}
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
    </div>
  );
};

export default UserHomePage;
