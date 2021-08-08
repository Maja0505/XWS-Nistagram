import React, { useState,useEffect } from "react";
import { Grid, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import ButtonBase from "@material-ui/core/ButtonBase";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ModeCommentIcon from "@material-ui/icons/ModeComment";
import { Redirect } from "react-router-dom";
import PostDialog from "./PostDialog";
import ThumbUpAltIcon from '@material-ui/icons/ThumbUpAlt';
import ThumbDownIcon from '@material-ui/icons/ThumbDown';
import axios from "axios";
import Background from "../images/background.jpg"



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
    display: "flex",
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

const Collections = ({userForProfile}) => {
  const classes = useStyles();
  const [redirection, setRedirectiton] = useState(false);
  const [collections, setCollections] = useState([]);
  const [collectionName, setCollectionName] = useState({})

  const loggedUserId = localStorage.getItem("id")
  const username = localStorage.getItem("username")

  const handleClickImage = (collection) => {
    setRedirectiton(true);
    setCollectionName(collection)
  };

  useEffect(() => {
    axios.get('/api/post/get-collections-for-user/' + loggedUserId)
      .then((res)=>
      {
        setCollections(res.data)
      }).catch((error) => {

      })
  }, [userForProfile])

  
  return (
    <div className={classes.root}>
      {redirection === true && <Redirect to={"/homePage/" + username + "/collection/" + collectionName}></Redirect>}
     



      <Grid container>
      <Grid item xs={4} style={{ margin: "auto", marginTop: "2%" }} key="allPosts">
      <ButtonBase
              focusRipple
              key="allPosts"
              className={classes.image}
              focusVisibleClassName={classes.focusVisible}
              style={{
                width: "95%",
              }}
              onClick={() => handleClickImage("allPosts")}
            >
              <span
                className={classes.imageSrc}
                style={{
                  backgroundImage: `url(${Background}) `,
                }}
              />
              <span className={classes.imageBackdrop} />
              <span className={classes.imageButton}>
                <Typography
                  component="span"
                  variant="subtitle1"
                  color="inherit"
                  className={classes.imageTitle}
                  style={{ padding: 0,  width:'100%',textAlign:"center"}}
                  width="100%"
                >
                  <Grid container >
                      <Grid item xs={2}></Grid>
                      <Grid item xs={8}>
                         <h2>All Posts</h2>

                      </Grid>

                    <Grid item xs={2}></Grid>

                  </Grid>
                </Typography>
              </span>
            </ButtonBase>
        </Grid>
        {collections !== null && collections.map((collection) => (
          <Grid item xs={4} style={{ margin: "auto", marginTop: "2%" }} key={collection}>
            <ButtonBase
              focusRipple
              key={collection}
              className={classes.image}
              focusVisibleClassName={classes.focusVisible}
              style={{
                width: "95%",
              }}
              onClick={() => handleClickImage(collection)}
            >
              <span
                className={classes.imageSrc}
                style={{
                  backgroundImage: `url(${Background}) `,
                }}
              />
              <span className={classes.imageBackdrop} />
              <span className={classes.imageButton}>
                <Typography
                  component="span"
                  variant="subtitle1"
                  color="inherit"
                  className={classes.imageTitle}
                  style={{ padding: 0,  width:'100%',textAlign:"center"}}
                  width="100%"
                >
                  <Grid container >
                      <Grid item xs={2}></Grid>
                      <Grid item xs={8}>
                         <h2>{collection}</h2>

                      </Grid>

                    <Grid item xs={2}></Grid>

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

export default Collections;
