import { makeStyles } from "@material-ui/core/styles";
import { deepOrange } from "@material-ui/core/colors";
import Avatar from "@material-ui/core/Avatar";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import InputBase from "@material-ui/core/InputBase";
import { Grid, Paper, Divider, List } from "@material-ui/core";
import { Button } from "@material-ui/core";
import BookmarkBorderSharpIcon from "@material-ui/icons/BookmarkBorderSharp";
import FavoriteBorderSharpIcon from "@material-ui/icons/FavoriteBorderSharp";
import SendOutlinedIcon from "@material-ui/icons/SendOutlined";
import ChatBubbleOutlineIcon from "@material-ui/icons/ChatBubbleOutline";
import ThumbUpAltOutlinedIcon from "@material-ui/icons/ThumbUpAltOutlined";
import ThumbUpAltIcon from '@material-ui/icons/ThumbUpAlt';
import ThumbDownAltOutlinedIcon from "@material-ui/icons/ThumbDownAltOutlined";
import ThumbDownIcon from '@material-ui/icons/ThumbDown';
import { useParams } from "react-router-dom";
import axios from "axios";
import React, { useState, useEffect } from "react";
import CommentsForPost from "./CommentsForPost";

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
  const [newComment, setNewComment] = useState("");

  const { post } = useParams();
  const [imagePost, setImagePost] = useState();
  const [commentsForPost, setCommentsForPost] = useState([]);
  const [postIsLiked, setPostIsLiked] = useState(false)
  const [postIsDisliked, setPostIsDisliked] = useState(false)

  const handleClickPostComment = () => {
    var comment = {
      ID: "0433e24a-d6d2-46f7-9111-33152b10d846",
      PostID: post,
      UserID: loggedUserId,
      CreatedAt: "2018-12-10T13:49:51.141Z",
      Content: newComment,
    };
    axios.post("/api/post/add-comment", comment).then((res) => {
      console.log("upisan komentar");
      setNewComment("");
      axios.get("/api/post/get-comments-for-post/" + post).then((res) => {
        setCommentsForPost(res.data);
      });
    });
  };

  useEffect(() => {
    axios.get("/api/post/get-one-post/" + post).then((res) => {
      setImagePost(res.data);
      console.log(res.data);
    });
    var like = {
      PostID : post,
      UserID : loggedUserId
    }
    console.log(like)
    axios.put("/api/post/like-exists",like)
      .then((res)=>{
        if(res.data == true && res.status === 201){
          setPostIsLiked(true)
        }else if(res.data == false && res.status === 201){
          setPostIsLiked(false)
        }
      })

    axios.put("/api/post/dislike-exists",like)
      .then((res)=>{
        if(res.data == true && res.status === 201){
          setPostIsDisliked(true)
        }else if(res.data == false && res.status === 201){
          setPostIsDisliked(false)
        }
      })
    
  }, []);

  const HandleClickLike = () => {
    var like = {
      PostID: post,
      UserID: loggedUserId,
    };
    axios.post("/api/post/like-post",like).then((res) => {
      if(postIsLiked){
        setPostIsLiked(false)

      }else{
        setPostIsLiked(true)
      }
    });
  };

  const HandleClickDislike = () => {
    var dislike = {
      PostID: post,
      UserID: loggedUserId,
    };
    axios.post("/api/post/dislike-post", dislike).then((res) => {
      if(postIsDisliked){
        setPostIsDisliked(false)

      }else{
        setPostIsDisliked(true)
      }
    });
  };

  const handleClickAllLikes = () => {
    axios.get("/api/post/get-users-who-liked-post/" + post).then((res) => {
      console.log(res.data);
    });
  };

  const handleClickAllDislikes = () => {
    axios.get("/api/post/get-users-who-disliked-post/" + post).then((res) => {
      console.log(res.data);
    });
  };

  return (
    <div>
      <Grid container>
        <Grid item xs={2}></Grid>
        <Grid item xs={8}>
       <Paper
        style={{
          width: "100%",
          height: 600,
          marginTop: "5%",
        }}
        variant="outlined"
      >
        <Grid container style={{ width: "100%", height: "100%" }}>
          {imagePost !== undefined && imagePost !== null && (
            <Grid item xs={7}>
              <img
                src={
                  "http://localhost:8080/api/post/get-image/" + imagePost.Image
                }
                style={{ width: "100%", height: "100%" }}
              />
            </Grid>
          )}

          <Grid item xs={5} >
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
                  style={{
                    marginTop: "50%",
                    textAlign: "right",
                    cursor: "pointer",
                  }}
                ></MoreHorizIcon>
              </Grid>
            </Grid>
            <Grid container style={{ height: "60%", overflow: "auto" }}>
              <CommentsForPost
                commentsForPost={commentsForPost}
                setCommentsForPost={setCommentsForPost}
              ></CommentsForPost>
            </Grid>
            <Grid container style={{ height: "25%" }}>
              <Grid container style={{ height: "70%" }}>
                <Grid container style={{ height: "30%" }}>
                  <Grid item xs={2}>
                    <Divider />
                    {postIsLiked ? 
                    <ThumbUpAltIcon
                    onClick={HandleClickLike}
                    fontSize="large"
                    style={{ cursor: "pointer" ,color:"blue"}}
                  ></ThumbUpAltIcon>
                  :
                    <ThumbUpAltOutlinedIcon
                      onClick={HandleClickLike}
                      fontSize="large"
                      style={{ cursor: "pointer" ,color:"blue"}}
                    ></ThumbUpAltOutlinedIcon>}
                  </Grid>
                  <Grid item xs={2}>
                    <Divider />
                  
                   {postIsDisliked ? 
                   <ThumbDownIcon
                   onClick={HandleClickDislike}
                   fontSize="large"
                   style={{ cursor: "pointer",color:"blue" }}
                  ></ThumbDownIcon>
                   :
                   <ThumbDownAltOutlinedIcon
                      onClick={HandleClickDislike}
                      fontSize="large"
                      style={{ cursor: "pointer",color:"blue" }}
                    ></ThumbDownAltOutlinedIcon>
                    }

                  </Grid>

                  <Grid item xs={2}>
                    <Divider />
                    <SendOutlinedIcon
                      fontSize="large"
                      style={{ cursor: "pointer"}}
                    ></SendOutlinedIcon>
                  </Grid>
                  <Grid item xs={6}>
                    <Divider />
                    <BookmarkBorderSharpIcon
                      style={{ marginLeft: "80%", cursor: "pointer" }}
                      fontSize="large"
                    ></BookmarkBorderSharpIcon>
                  </Grid>
                </Grid>
                <Grid container style={{ height: "70%" }}>
                  <Grid item xs={5}>
                    <h5
                      onClick={handleClickAllLikes}
                      style={{ cursor: "pointer" }}
                    >
                      Likes
                    </h5>
                  </Grid>
                  <Grid item xs={5}>
                    <h5
                      onClick={handleClickAllDislikes}
                      style={{ cursor: "pointer" }}
                    >
                      Dislikes
                    </h5>
                  </Grid>
                  <Grid item xs={2}></Grid>
                </Grid>
              </Grid>
              <Grid container style={{ height: "30%" }}>
                <Grid item xs={3}>
                  <Divider />

                  <SentimentSatisfiedRoundedIcon
                    style={{
                      margin: "auto",
                      marginTop: "5%",
                      cursor: "pointer",
                    }}
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
                  <Button
                    disabled={newComment === ""}
                    style={{ margin: "auto", marginTop: "10%" }}
                    color="primary"
                    onClick={handleClickPostComment}
                  >
                    Post
                  </Button>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Paper></Grid>
        <Grid item xs={2}></Grid>
      </Grid>

    </div>
  );
};

export default PostDialog;
