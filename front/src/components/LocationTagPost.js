import { useParams } from "react-router-dom";
import { RoomRounded } from "@material-ui/icons";
import { Grid, Typography } from "@material-ui/core";
import PostsForHashTag from "./PostsForHashTag.js";

import { useEffect, useState } from "react";

import axios from "axios";

const LocationTagPost = () => {
  const { location } = useParams();

  const [posts, setPosts] = useState();

  useEffect(() => {
    axios.get("/api/post/get-all-by-location/" + location).then((res) => {
      setPosts(res.data);
    }).catch((error) => {
      //console.log(error);
    });
  }, [location]);

  return (
    <div>
      <Grid container style={{ marginTop: "3%" }}>
        <Grid item xs={2} />
        <Grid container item xs={8}>
          <Grid item xs={3}>
            {posts !== undefined && posts !== null && (
              <>
                {posts[0].Media[0].substring(
                  posts[0].Media[0].length - 3,
                  posts[0].Media[0].length
                ) === "jpg" && (
                  <img
                    src={
                      "http://localhost:8080/api/media/get-media-image/" +
                      posts[0].Media[0]
                    }
                    alt="Not founded"
                    style={{
                      borderRadius: "50%",
                      border: "0px solid",
                      width: "150px",
                      height: "150px",
                    }}
                  />
                )}
                {posts[0].Media[0].substring(
                  posts[0].Media[0].length - 3,
                  posts[0].Media[0].length
                ) === "mp4" && (
                  <video
                    controls
                    style={{
                      borderRadius: "50%",
                      border: "0px solid",
                      width: "150px",
                      height: "150px",
                    }}
                  >
                    <source
                      src={`http://localhost:8080/api/media/get-video/${posts[0].Media[0]}`}
                      type="video/mp4"
                    />
                  </video>
                )}
              </>
            )}
          </Grid>
          <Grid item xs={3} style={{ textAlign: "left", margin: "auto" }}>
            <Typography variant="h4" color="textSecondary">
              <RoomRounded fontSize="large" />
              {location}
            </Typography>
            <Typography variant="h6" color="textPrimary">
              <b>
                {posts !== undefined && posts !== null ? posts.length : `0`}
              </b>{" "}
              posts
            </Typography>
          </Grid>
          <Grid item xs={6} />
        </Grid>
        <Grid item xs={2} />
      </Grid>

      <Grid container style={{ marginTop: "5%" }}>
        <Grid item xs={2}></Grid>
        <Grid item xs={8}>
          <PostsForHashTag posts={posts}></PostsForHashTag>
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
    </div>
  );
};

export default LocationTagPost;
