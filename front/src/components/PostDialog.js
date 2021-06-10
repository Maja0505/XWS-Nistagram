import { makeStyles } from "@material-ui/core/styles";
import { deepOrange } from "@material-ui/core/colors";
import Avatar from "@material-ui/core/Avatar";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import InputBase from '@material-ui/core/InputBase';
import { Grid,Paper,Divider, List  } from "@material-ui/core";
import {Button } from "@material-ui/core";
import BookmarkBorderSharpIcon from '@material-ui/icons/BookmarkBorderSharp';
import FavoriteBorderSharpIcon from '@material-ui/icons/FavoriteBorderSharp';
import SendOutlinedIcon from '@material-ui/icons/SendOutlined';
import ChatBubbleOutlineIcon from '@material-ui/icons/ChatBubbleOutline';
import ThumbUpAltOutlinedIcon from '@material-ui/icons/ThumbUpAltOutlined';
import ThumbDownAltOutlinedIcon from '@material-ui/icons/ThumbDownAltOutlined';
import { useParams } from "react-router-dom";
import axios from "axios";
import React, { useState,useEffect } from "react";
import CommentsForPost from './CommentsForPost'




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



const PostDialog = () => {
  const classes = useStyles();
  const username = localStorage.getItem("username");
  const loggedUserId = localStorage.getItem("id");
  const [newComment,setNewComment] = useState('');

  const { post } = useParams();
  const [imagePost, setImagePost] = useState()
  const [commentsForPost,setCommentsForPost] = useState([])


  const handleClickPostComment = () => {
    var comment = {
      ID:'0433e24a-d6d2-46f7-9111-33152b10d846',
      PostID: post,
      UserID:loggedUserId,
      CreatedAt: '2018-12-10T13:49:51.141Z',
      Content: newComment
    }
    axios.post('/api/post/add-comment', comment)
      .then((res) => {
        console.log('upisan komentar')
        setNewComment('')
        axios.get('/api/post/get-comments-for-post/' + post)
        .then((res) => {
            setCommentsForPost(res.data)
            
        })
      })
  }

  useEffect(() => {
    axios.get('/api/post/get-one-post/' + post)
      .then((res)=>
      {
        setImagePost(res.data)
        console.log(res.data)
      })
      axios.get('/api/post/get-one-post/' + post)
      .then((res)=>
      {
        setImagePost(res.data)
        console.log(res.data)
      })
  }, [])

  const HandleClickLike = () => {
    var like = {
      PostID: post.ID,
      UserID: loggedUserId
    }
    axios.post('/api/post/like-post',like)
      .then((res) => {
        console.log("uspeloo")
      })
  }

  const HandleClickDislike = () => {
    var dislike = {
      PostID: post.ID,
      UserID: loggedUserId
    }
    axios.post('/api/post/dislike-post',dislike)
      .then((res) => {
        console.log("uspeloo")
      })
  }

  return (
    <div>
      <Paper
        style={{ width: "60%", height: '50%',maxHeight:510, marginLeft:'auto',marginRight:'auto',marginTop:'5%'}}
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
                  style={{ marginTop: "50%", textAlign: "right",cursor: 'pointer'}}
                ></MoreHorizIcon>
              </Grid>
            </Grid>
            <Grid container style={{maxHeight: '60%',  overflow: 'auto'
          }}><CommentsForPost commentsForPost={commentsForPost} setCommentsForPost={setCommentsForPost}></CommentsForPost></Grid>
            <Grid container style={{ height: "25%" }}>
              <Grid container style={{ height: "60%" }}>
                <Grid container style={{ height: "40%" }}>
                    <Grid item xs={2}>
                        <Divider />
                        <ThumbUpAltOutlinedIcon onClick={HandleClickLike} fontSize="large" style={{cursor: 'pointer'}}></ThumbUpAltOutlinedIcon>
                    </Grid>
                    <Grid item xs={2}>
                        <Divider />
                        <ThumbDownAltOutlinedIcon  onClick={HandleClickDislike} fontSize="large" style={{cursor: 'pointer'}}></ThumbDownAltOutlinedIcon>
                    </Grid>

                    <Grid item xs={2}>
                        <Divider />
                        <SendOutlinedIcon  fontSize="large" style={{cursor: 'pointer'}}></SendOutlinedIcon>
                    </Grid>
                    <Grid item xs={6}>
                    <Divider />
                        <BookmarkBorderSharpIcon style={{marginLeft:"80%",cursor: 'pointer'}} fontSize="large" ></BookmarkBorderSharpIcon>
                    </Grid>

                </Grid>
                <Grid container style={{ height: "60%" }}></Grid>
              </Grid>
              <Grid container style={{ height: "40%" }}>
                <Grid item xs={3}>
                <Divider />

                  <SentimentSatisfiedRoundedIcon
                    style={{ margin: "auto", marginTop: "5%",cursor: 'pointer'}}
                    fontSize="large"
                  ></SentimentSatisfiedRoundedIcon>
                </Grid>
                <Grid item xs={7}>
                <Divider />

                  <InputBase
                    className={classes.margin}
                    placeholder="Add a comment..."
                    inputProps={{ "aria-label": "naked" }}
                    value={newComment}
                    onChange={(e) => setNewComment(e.target.value)}
                  />
                </Grid>
                <Grid item xs={2}>
                <Divider />
                    <Button disabled={newComment === ''} style={{ margin: "auto", marginTop: "10%" }} color="primary" onClick={handleClickPostComment}>Post</Button>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
    </div>
  );
};

export default PostDialog;
