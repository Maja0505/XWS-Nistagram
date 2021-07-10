import { useState, useEffect } from "react";
import { Grid, Paper, Typography } from "@material-ui/core";
import PostsForHashTag from "./PostsForHashTag.js";

import axios from "axios";

const PostWhereUserTagged = ({ user }) => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    axios.get("/api/post/get-all-by-tag/@" + user.Username).then((res) => {
      if (res.data !== null) {
        setPosts(res.data);
      }
    }).catch((error) => {
      //console.log(error);
    });
  }, []);

  return (
    <div>
      <Grid container>
        <PostsForHashTag posts={posts}></PostsForHashTag>
        {posts.length === 0 && (
          <Paper style={{ width: "100%", height: "100%" }}>
            <Typography variant="h5" color="textSecondary">
              No one has tagged you yet
            </Typography>
          </Paper>
        )}
      </Grid>
    </div>
  );
};

export default PostWhereUserTagged;
