import React, { useState, useEffect } from "react";
import { Grid, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import ButtonBase from "@material-ui/core/ButtonBase";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ModeCommentIcon from "@material-ui/icons/ModeComment";
import { Redirect } from "react-router-dom";
import PostDialog from "./PostDialog";
import ThumbUpAltIcon from "@material-ui/icons/ThumbUpAlt";
import ThumbDownIcon from "@material-ui/icons/ThumbDown";
import axios from "axios";
import { useParams } from "react-router-dom";

const images = [
  {
    url: "https://besthqwallpapers.com/Uploads/5-12-2020/148667/thumb2-mount-fuji-4k-two-swans-autumn-stratovolcano.jpg",
    title: "Breakfast",
  },
  {
    url: "https://besthqwallpapers.com/Uploads/5-12-2020/148667/thumb2-mount-fuji-4k-two-swans-autumn-stratovolcano.jpg",
    title: "Burgers",
  },
  {
    url: "https://besthqwallpapers.com/Uploads/5-12-2020/148667/thumb2-mount-fuji-4k-two-swans-autumn-stratovolcano.jpg",
    title: "Camera",
  },
  {
    url: "https://besthqwallpapers.com/Uploads/5-12-2020/148667/thumb2-mount-fuji-4k-two-swans-autumn-stratovolcano.jpg",
    title: "Breakfast2",
  },
];

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
      width: "100% !important", // Overrides inline-style
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
  /*imageMarked: {
    height: 3,
    width: 18,
    backgroundColor: theme.palette.common.white,
    position: "absolute",
    bottom: -2,
    left: "calc(50% - 9px)",
    transition: theme.transitions.create("opacity"),
  },*/
}));

const PostsForCollection = () => {
  const classes = useStyles();
  const [redirection, setRedirectiton] = useState(false);
  const [posts, setPosts] = useState([]);
  const [postID, setPostID] = useState({});
  const loggedUserId = localStorage.getItem("id");
  const username = localStorage.getItem("username");
  const { collection } = useParams();

  const handleClickImage = (post) => {
    setRedirectiton(true);
    setPostID(post.ID);
  };

  useEffect(() => {
    if (collection === "allPosts") {
      axios
        .get("/api/post/get-favourite-posts/" + loggedUserId)
        .then((res) => {
          setPosts(res.data);
        })
        .catch((error) => {});
    } else {
      axios
        .get(
          "/api/post/get-posts-from-collection/" +
            loggedUserId +
            "/" +
            collection
        )
        .then((res) => {
          setPosts(res.data);
        })
        .catch((error) => {});
    }
  }, []);

  const getImage = (image) => {
    axios.get("/api/media/get-media-image/" + image).then((res) => {
      return res.data;
    }).catch((error) => {
      //console.log(error);
    });
  };

  return (
    <Grid container>
      <Grid item xs={2}></Grid>
      <Grid item xs={8}>
        {redirection === true && <Redirect to={"/dialog/" + postID}></Redirect>}
        <Grid container style={{ margin: "auto" }}>
          {collection === "allPosts" && <h3>All Posts</h3>}
          {collection !== "allPosts" && <h3>{collection}</h3>}
        </Grid>
        <Grid container>
          {posts !== null &&
            posts.map((post) => (
              <Grid item xs={4} style={{ margin: "auto" }} key={post.ID}>
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
                  {post.Image.substring(
                    post.Image.length - 3,
                    post.Image.length
                  ) === "jpg" && (
                    <img
                      width="100%"
                      height="100%"
                      src={`http://localhost:8080/api/media/get-media-image/${post.Image}`}
                    />
                  )}
                  {post.Image.substring(
                    post.Image.length - 3,
                    post.Image.length
                  ) !== "jpg" && (
                    <video width="100%" height="100%" controls>
                      <source
                        src={`http://localhost:8080/api/media/get-video/${post.Image}`}
                        type="video/mp4"
                      />
                    </video>
                  )}
                  <span className={classes.imageBackdrop} />
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
          {images.length % 3 === 1 && (
            <>
              <Grid item xs={4} /> <Grid item xs={4} />
            </>
          )}
          {images.length % 3 === 2 && <Grid item xs={4} />}
        </Grid>
      </Grid>
      <Grid item xs={2}></Grid>
    </Grid>
  );
};

export default PostsForCollection;
