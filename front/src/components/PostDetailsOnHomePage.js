import { Grid, Typography, Paper, Tabs, Tab } from "@material-ui/core";

import { useState, useEffect } from "react";
import Posts from "./Posts";
import LocalMallOutlinedIcon from "@material-ui/icons/LocalMallOutlined";
import Collections from "./Collections.js";
import {
  GridOn,
  BookmarkBorder,
  AssignmentIndOutlined,
  AddPhotoAlternateOutlined,
} from "@material-ui/icons";
import AddPost from "./AddPost";
import PostWhereUserTagged from "./PostWhereUserTagged.js";
import AddCampaign from "./AddCampaign";

const PostDetailsOnHomePage = ({ user, username, following }) => {
  const loggedUsername = localStorage.getItem("username");
  const [tabValue, setTabValue] = useState(0);

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  return (
    <div>
      {loggedUsername === username && (
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

                {user.IsAgent && (
                  <Tab
                    label="Add campaign"
                    icon={<LocalMallOutlinedIcon />}
                    style={{ margin: "auto" }}
                  />
                )}
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
              {user !== undefined && user !== null && tabValue === 4 && (
                <AddCampaign setTabValue={setTabValue} />
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
    </div>
  );
};

export default PostDetailsOnHomePage;
