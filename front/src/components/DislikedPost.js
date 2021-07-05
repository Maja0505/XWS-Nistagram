import { makeStyles } from "@material-ui/core/styles";
import { ThumbUpAlt, ModeComment, ThumbDown } from "@material-ui/icons";
import { Grid, Typography, ButtonBase } from "@material-ui/core";

import { Redirect } from "react-router-dom";

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
}));

const LikedPost = () => {
  const classes = useStyles();
  const [images, setImages] = useState([]);
  const loggedUserId = localStorage.getItem("id");
  const [redirection, setRedirectiton] = useState(false);
  const [postID, setPostID] = useState({});

  useEffect(() => {
    axios
      .get("/api/post/get-disliked-posts-for-user/" + loggedUserId)
      .then((res) => {
        if (res.data) {
          setImages(res.data);
        } else {
          setImages([]);
        }
      })
      .catch((error) => {
        alert(error.response.status);
      });
  }, [loggedUserId]);

  const handleClickImage = (image) => {
    setRedirectiton(true);
    setPostID(image.ID);
  };

  return (
    <div className={classes.root}>
      {redirection === true && <Redirect to={"/dialog/" + postID}></Redirect>}
      <Grid container>
        {images !== null && images.length === 0 && (
          <Grid item style={{ margin: "auto" }}>
            <Typography variant="h5" color="textSecondary">
              No disliked posts
            </Typography>
          </Grid>
        )}
        {images !== null &&
          images.length !== 0 &&
          images.map((image) => (
            <Grid item xs={4} style={{ margin: "auto", marginTop: "2%" }}>
              <ButtonBase
                focusRipple
                className={classes.image}
                focusVisibleClassName={classes.focusVisible}
                style={{
                  width: "95%",
                }}
                onClick={() => handleClickImage(image)}
              >
                {image.Image.substring(
                  image.Image.length - 3,
                  image.Image.length
                ) === "jpg" && (
                  <img
                    width="100%"
                    height="100%"
                    src={`http://localhost:8080/api/media/get-media-image/${image.Image}`}
                  />
                )}
                {image.Image.substring(
                  image.Image.length - 3,
                  image.Image.length
                ) !== "jpg" && (
                  <video width="100%" height="100%" controls>
                    <source
                      src={`http://localhost:8080/api/media/get-video/${image.Image}`}
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
                        <ThumbUpAlt></ThumbUpAlt>
                      </Grid>
                      <Grid item xs={3}>
                        {image.LikesCount}
                      </Grid>
                      <Grid item xs={1}>
                        <ThumbDown></ThumbDown>
                      </Grid>
                      <Grid item xs={3}>
                        {image.DislikesCount}
                      </Grid>
                      <Grid item xs={1}>
                        <ModeComment></ModeComment>
                      </Grid>
                      <Grid item xs={3}>
                        {image.CommentsCount}
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
    </div>
  );
};

export default LikedPost;
