import { Grid } from "@material-ui/core";

import { useState, useEffect } from "react";
import axios from "axios";

import PostFeed from "./PostFeed.js";
import StoryFeeds from "./StoryFeeds.js";

const StartPage = () => {
  const loggedUserId = localStorage.getItem("id");

  const [feeds, setFeeds] = useState([]);

  useEffect(() => {
    axios
      .get("/api/post/get-all-post-feeds-for-user/" + loggedUserId)
      .then((res) => {
        console.log(res.data);
        setFeeds(res.data);
      }).catch((error) => {
        
      }
    );
  }, []);

  return (
    <div>
      {loggedUserId !== null && loggedUserId !== undefined && (
        <div>
          <Grid container style={{ marginTop: "1%" }}>
            <Grid item xs={3} />
            <Grid item xs={6}>
              <StoryFeeds />
            </Grid>
            <Grid item xs={3} />
          </Grid>
          {feeds !== null && feeds !== undefined && (
            <>
              <Grid container style={{ marginTop: "1%" }}></Grid>
              {feeds.map((f, index) => (
                <Grid container style={{ marginTop: "3%" }} key={index}>
                  <Grid item xs={3} />
                  <Grid item xs={6}>
                    <PostFeed feed={f} />
                  </Grid>
                  <Grid item xs={3} />
                </Grid>
              ))}
            </>
          )}

          <Grid container style={{ marginTop: "5%" }}></Grid>
        </div>
      )}
    </div>
  );
};

export default StartPage;
