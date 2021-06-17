import { Grid, Paper, Tabs, Tab } from "@material-ui/core";

import { useParams, Redirect } from "react-router-dom";
import { useState } from "react";
import { ThumbUpAltOutlined, ThumbDownAltOutlined } from "@material-ui/icons";

import LikedPost from "./LikedPost.js";
import DislikedPost from "./DislikedPost.js";

const LikedDislikedPost = () => {
  const { username } = useParams();
  const [tabValue, setTabValue] = useState(0);
  const loggedUsername = localStorage.getItem("username");

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  return (
    <div>
      {username !== undefined &&
        loggedUsername !== undefined &&
        username !== loggedUsername && (
          <Redirect to={"/homePage/" + `${loggedUsername}`}></Redirect>
        )}
      <Grid container style={{ marginTop: "4%" }}>
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
                label="Posts You've liked"
                icon={<ThumbUpAltOutlined />}
                style={{ margin: "auto" }}
              />
              <Tab
                label="Posts You've Disliked"
                icon={<ThumbDownAltOutlined />}
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
          {loggedUsername !== undefined &&
            loggedUsername !== null &&
            tabValue === 0 && <LikedPost></LikedPost>}
          {loggedUsername !== undefined &&
            loggedUsername !== null &&
            tabValue === 1 && <DislikedPost></DislikedPost>}
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
    </div>
  );
};

export default LikedDislikedPost;
