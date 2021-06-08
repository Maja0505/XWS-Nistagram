import { makeStyles } from "@material-ui/core/styles";
import { deepOrange } from "@material-ui/core/colors";
import Avatar from "@material-ui/core/Avatar";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import InputBase from '@material-ui/core/InputBase';
import { Grid,Paper,Divider  } from "@material-ui/core";
import {Button } from "@material-ui/core";
import BookmarkBorderSharpIcon from '@material-ui/icons/BookmarkBorderSharp';
import FavoriteBorderSharpIcon from '@material-ui/icons/FavoriteBorderSharp';
import SendOutlinedIcon from '@material-ui/icons/SendOutlined';
import ChatBubbleOutlineIcon from '@material-ui/icons/ChatBubbleOutline';
import { useParams } from "react-router-dom";
import axios from "axios";
import React, { useState,useEffect } from "react";




const useStyles = makeStyles((theme) => ({
  orange: {
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
    marginLeft: "auto",
  },
  margin: {
    margin: theme.spacing(1),
  },
}));



const ProfileDialog = () => {
  const classes = useStyles();
  const [open, setOpen] = React.useState(true);
  const username = localStorage.getItem("username");
  const { post } = useParams();
  const [imagePost, setImagePost] = useState()


  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleClickPost = () => {

  }

  useEffect(() => {
    axios.get('/api/post/get-one-post/' + post)
      .then((res)=>
      {
        setImagePost(res.data)
      })
  }, [])

  return (
    <div>
      <Paper
        style={{ width: "60%", height: 500, marginLeft:'auto',marginRight:'auto',marginTop:'5%'}}
        variant="outlined"
      >
        <Grid container style={{ width: "100%", height: "100%" }}>
          {imagePost !== undefined && imagePost !== null && <Grid item xs={7}>
            <img
              src={"http://localhost:8080/api/post/get-image/" + imagePost.Image}
              style={{ width: "100%", height: "100%" }}
            />
          </Grid>}
          
          <Grid item xs={5}>
            <Grid container style={{ height: "15%" }}>
              <Grid item xs={3}>
                <Avatar
                  className={classes.orange}
                  style={{ margin: "auto", marginTop: "25%" }}
                >
                  N
                </Avatar>
              </Grid>
              <Grid item xs={7}>
                <h4 style={{ marginTop: "10%", textAlign: "left" }}>
                  {username}
                </h4>
              </Grid>
              <Grid item xs={2}>
                <MoreHorizIcon
                  style={{ marginTop: "50%", textAlign: "right" }}
                ></MoreHorizIcon>
              </Grid>
            </Grid>
            <Grid container style={{ height: "60%" }}></Grid>
            <Grid container style={{ height: "25%" }}>
              <Grid container style={{ height: "60%" }}>
                <Grid container style={{ height: "40%" }}>
                    <Grid item xs={2}>
                        <Divider />
                        <FavoriteBorderSharpIcon  fontSize="large"></FavoriteBorderSharpIcon>
                    </Grid>
                    <Grid item xs={2}>
                         <Divider />
                        <ChatBubbleOutlineIcon  fontSize="large"></ChatBubbleOutlineIcon>
                    </Grid>
                    <Grid item xs={2}>
                        <Divider />
                        <SendOutlinedIcon  fontSize="large"></SendOutlinedIcon>
                    </Grid>
                    <Grid item xs={6}>
                    <Divider />
                        <BookmarkBorderSharpIcon style={{marginLeft:"80%"}} fontSize="large"></BookmarkBorderSharpIcon>
                    </Grid>

                </Grid>
                <Grid container style={{ height: "60%" }}></Grid>
              </Grid>
              <Grid container style={{ height: "40%" }}>
                <Grid item xs={3}>
                <Divider />

                  <SentimentSatisfiedRoundedIcon
                    style={{ margin: "auto", marginTop: "5%" }}
                    fontSize="large"
                  ></SentimentSatisfiedRoundedIcon>
                </Grid>
                <Grid item xs={7}>
                <Divider />

                  <InputBase
                    className={classes.margin}
                    placeholder="Add a comment..."
                    inputProps={{ "aria-label": "naked" }}
                  />
                </Grid>
                <Grid item xs={2}>
                <Divider />

                    <Button  style={{ margin: "auto", marginTop: "10%" }} color="primary" onClick={handleClickPost}>Post</Button>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
    </div>
  );
};

export default ProfileDialog;
