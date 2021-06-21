import { useParams } from "react-router-dom";
import { Grid, Typography } from "@material-ui/core";
import { useEffect, useState } from "react";

import PostsForHashTag from "./PostsForHashTag";

import axios from "axios";

const HashTagPost = () => {
  const { tag } = useParams();
  const [posts, setPosts] = useState();

  useEffect(() => {
    axios.get("/api/post/get-all-by-tag/" + tag).then((res) => {
      console.log(res.data);
      setPosts(res.data);
    });
  }, []);

  return (
    <div>
      <Grid container style={{ marginTop: "3%" }}>
        <Grid item xs={2} />
        <Grid container item xs={8}>
          <Grid item xs={3}>
            {posts !== undefined && posts !== null && (
              <>
                {posts[0].Image.substring(
                  posts[0].Image.length - 3,
                  posts[0].Image.length
                ) === "jpg" && (
                  <img
                    src={
                      "http://localhost:8080/api/post/get-image/" +
                      posts[0].Image
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
                {posts[0].Image.substring(
                  posts[0].Image.length - 3,
                  posts[0].Image.length
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
                      src={`http://localhost:8080/api/post/video-get/${posts[0].Image}`}
                      type="video/mp4"
                    />
                  </video>
                )}
              </>
            )}
          </Grid>
          <Grid item xs={3} style={{ textAlign: "left", margin: "auto" }}>
            <Typography variant="h4" color="textSecondary">
              #{tag}
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

export default HashTagPost;
