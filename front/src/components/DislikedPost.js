import { Grid, Typography, Paper } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import ButtonBase from "@material-ui/core/ButtonBase";
import ModeCommentIcon from "@material-ui/icons/ModeComment";
import { Redirect } from "react-router-dom";
import ThumbUpAltIcon from "@material-ui/icons/ThumbUpAlt";
import ThumbDownIcon from "@material-ui/icons/ThumbDown";
import PhotoLibraryOutlinedIcon from "@material-ui/icons/PhotoLibraryOutlined";
import { useEffect, useState } from "react";

import axios from "axios";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
    flexWrap: "wrap",
    minWidth: 300,
    width: "100%",
    marginTop: "3%",
  },
  image: {
    position: "relative",
    height: 250,
    [theme.breakpoints.down("xs")]: {
      width: "100% !important",
      height: 100,
    },
    "&:hover, &$focusVisible": {
      zIndex: 1,
      "& $imageBackdrop": {
        opacity: 0.6,
      },
      "& $imageMarked": {
        opacity: 0,
      },
      "& $imageTitle": {
        display: "flex",
      },
    },
  },
  focusVisible: {},
  imageButton: {
    position: "absolute",
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    color: theme.palette.common.white,
  },
  imageSrc: {
    position: "absolute",
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
    backgroundSize: "cover",
    backgroundPosition: "center 40%",
  },
  imageBackdrop: {
    position: "absolute",
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
    backgroundColor: theme.palette.common.black,
    opacity: 0.2,
    transition: theme.transitions.create("opacity"),
  },
  imageTitle: {
    position: "relative",
    display: "none",
  },
  imageButton2: {
    position: "absolute",
    top: 0,
    marginTop: "5%",
    marginLeft: "70%",
    color: theme.palette.common.white,
  },
  imageTitle2: {
    position: "relative",
    display: "flex",
  },
}));

const LikedPost = () => {
  const classes = useStyles();
  const [redirection, setRedirectiton] = useState(false);
  const [postID, setPostID] = useState({});
  const [posts, setPosts] = useState([]);
  const loggedUserId = localStorage.getItem("id");

  useEffect(() => {
    axios
      .get("/api/post/get-disliked-posts-for-user/" + loggedUserId)
      .then((res) => {
        if (res.data) {
          setPosts(res.data);
        } else {
          setPosts([]);
        }
      })
      .catch((error) => {
        //alert(error.response.status);
      });
  }, [loggedUserId]);

  const handleClickImage = (post) => {
    setRedirectiton(true);
    setPostID(post.ID);
  };

  return (
    <div className={classes.root}>
      {posts !== undefined && posts !== null && (
        <>
          {redirection === true && (
            <Redirect to={"/dialog/" + postID}></Redirect>
          )}
          <Grid container>
            {posts.length === 0 && (
              <Paper style={{ width: "100%", height: "100%" }}>
                <Typography variant="h5" color="textSecondary">
                  No one disliked post
                </Typography>
              </Paper>
            )}
            {posts !== null &&
              posts.map((post) => (
                <Grid
                  item
                  xs={4}
                  style={{ margin: "auto", marginTop: "2%" }}
                  key={post.ID}
                >
                  <ButtonBase
                    focusRipple
                    key={post.ID}
                    className={classes.image}
                    focusVisibleClassName={classes.focusVisible}
                    style={{
                      width: "95%",
                    }}
                    onClick={() => handleClickImage(post)}
                  >
                    {post.Media[0].substring(
                      post.Media[0].length - 3,
                      post.Media[0].length
                    ) === "jpg" && (
                      <img
                        width="100%"
                        height="100%"
                        src={`http://localhost:8080/api/media/get-media-image/${post.Media[0]}`}
                      />
                    )}
                    {post.Media[0].substring(
                      post.Media[0].length - 3,
                      post.Media[0].length
                    ) !== "jpg" && (
                      <video width="100%" height="100%" controls>
                        <source
                          src={`http://localhost:8080/api/media/get-video/${post.Media[0]}`}
                          type="video/mp4"
                        />
                      </video>
                    )}

                    <span className={classes.imageBackdrop} />

                    {post.Media.length !== 1 && (
                      <span className={classes.imageButton2}>
                        <Typography
                          component="span"
                          variant="subtitle1"
                          color="inherit"
                          className={classes.imageTitle2}
                          width="100%"
                          height="30%"
                        >
                          <Grid container style={{ margin: "auto" }}>
                            <Grid item xs={10}></Grid>
                            <Grid item xs={2}>
                              <PhotoLibraryOutlinedIcon></PhotoLibraryOutlinedIcon>
                            </Grid>
                          </Grid>
                        </Typography>
                      </span>
                    )}

                    <span className={classes.imageButton}>
                      <Typography
                        component="span"
                        variant="subtitle1"
                        color="inherit"
                        className={classes.imageTitle}
                        style={{ padding: 0, width: "50%" }}
                        width="30%"
                      >
                        <Grid container>
                          <Grid item xs={1}>
                            <ThumbUpAltIcon></ThumbUpAltIcon>
                          </Grid>
                          <Grid item xs={3}>
                            {post.LikesCount}
                          </Grid>
                          <Grid item xs={1}>
                            <ThumbDownIcon></ThumbDownIcon>
                          </Grid>
                          <Grid item xs={3}>
                            {post.DislikesCount}
                          </Grid>
                          <Grid item xs={1}>
                            <ModeCommentIcon></ModeCommentIcon>
                          </Grid>
                          <Grid item xs={3}>
                            {post.CommentsCount}
                          </Grid>
                        </Grid>
                      </Typography>
                    </span>
                  </ButtonBase>
                </Grid>
              ))}
            {posts.length % 3 === 1 && (
              <>
                <Grid item xs={4} /> <Grid item xs={4} />
              </>
            )}
            {posts.length % 3 === 2 && <Grid item xs={4} />}
          </Grid>
        </>
      )}
    </div>
  );
};

export default LikedPost;
