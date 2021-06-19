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
  const [following,setFollowing] = useState(false);
  const [redirection,setRedirection] = useState(false); 
  const [requested,setRequested] = useState(false); 
  const [privateProfile,setPrivateProfile] =useState();

  const loggedUsername = localStorage.getItem("username");
  const loggedInId = localStorage.getItem("id");
  
  const users=[{ 	Username :"Perica",
                  FirstName :"Perica",
                  LastName:"Peric",
                  DateOfBirth :"krdlkjf",
                  Email :"Peric.peric@gmail.com",
                  PhoneNumber :"0490843",
                  Gender :"Female",
                  Biography :"Jedna vrlo uspesan gospodin",
                  WebSite :"Pericaperic.com"},
  
                { 	Username :"marko",
                    FirstName :"Marko",
                    LastName:"Markovic",
                    DateOfBirth :"krdlkjf",
                    Email :"marko.markovic@gmail.com",
                    PhoneNumber :"0490843",
                    Gender :"Male",
                    Biography :"Jedna vrlo uspesan gospodin",
                    WebSite :"Pericaperic.com"},]

  useEffect(() => {
    console.log(username)
    console.log(loggedUsername)
    //setUser(users.filter(user => user.Username === username)[0])
    setRequested(true)
    setFollowing(false)
    setPrivateProfile(true)
    axios
    .get("/api/user/" + username)
    .then((res) => {
      console.log(res.data);
      setUser(res.data);

      if(res.data.IdString !== loggedInId ){
      axios
      .get("/api/user-follow/checkBlock/" + loggedInId + "/" + res.data.IdString)
      .then((res) => {
        console.log(res.data)
        setRedirection(res.data)
      })
      .catch((error) => {
        alert(error.response.status);
      });

      axios
      .get("/api/user-follow/checkRequested/" + loggedInId + "/" + res.data.IdString)
      .then((res) => {
        console.log(res.data)
        setRequested(res.data)
      })
      .catch((error) => {
        alert(error.response.status);
      });

      axios
      .get("/api/user-follow/checkFollowing/" + loggedInId + "/" + res.data.IdString)
      .then((res) => {
        console.log(res.data)
        setFollowing(res.data)
      })
      .catch((error) => {
        alert(error.response.status);
      });
    }
    })
    .catch((error) => {
      alert(error.response.status);
    });
    
  }, [username, loggedUsername]);

  

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  const requestedClicked= () => {
    setRequested(!requested)

  }
  const followClicked= () => {
    if(privateProfile){
      setRequested(true)
    }
    else{
      setFollowing(true)
    }

  }
  const unfollowClicked= () => {
    setFollowing(false)

  }

  const buttonForUnfollow = (
    <Button variant="contained" color="default" style={{ margin: "auto" }} onClick = {unfollowClicked}>
     Following 
    </Button>
  );

  const buttonForFollow = (
    <Button variant="contained" color="primary" style={{ margin: "auto" }}
    onClick = {followClicked} >
      Follow
    </Button>
  );

  const buttonForRequested = (
    <Button variant="contained" color="primary"  style={{ margin: "auto",marginLeft:"30px"
                                              ,backgroundColor:"whitesmoke",color:"darkgray" }}
                                              onClick={requestedClicked}>
      Requested
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
            {requested && loggedUsername !== username && buttonForRequested}
            {following && loggedUsername !== username && !requested && buttonForUnfollow}
            { !following && loggedUsername !== username && !requested && buttonForFollow}
          
            
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
