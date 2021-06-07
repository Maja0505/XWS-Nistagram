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

import {
  GridOn,
  BookmarkBorder,
  AssignmentIndOutlined,
} from "@material-ui/icons";

const UserHomePage = () => {
  const [user, setUser] = useState();
  const [tabValue, setTabValue] = useState(0);

  useEffect(() => {
    axios.get("/user/pera").then((res) => {
      setUser(res.data);
    });
  }, []);

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  return (
    <div>
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
                <Typography variant="h6" style={{ margin: "auto" }}>
                  {user.Username}
                </Typography>
              )}

              <Button
                variant="outlined"
                color="inherit"
                style={{ marginLeft: "auto" }}
              >
                Edit profile
              </Button>
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
          </Grid>
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
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
        <Grid item xs={8}></Grid>
        <Grid item xs={2}></Grid>
      </Grid>
    </div>
  );
};

export default UserHomePage;
