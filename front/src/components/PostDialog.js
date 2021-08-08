import Avatar from "@material-ui/core/Avatar";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import {
  Grid,
  Paper,
  Divider,
  Grow,
  Popper,
  MenuItem,
  MenuList,
  InputBase,
} from "@material-ui/core";
import { Button } from "@material-ui/core";
import {
  BookmarkBorderRounded,
  BookmarkRounded,
  SendRounded,
  PersonRounded,
} from "@material-ui/icons";
import ThumbUpAltOutlinedIcon from "@material-ui/icons/ThumbUpAltOutlined";
import ThumbUpAltIcon from "@material-ui/icons/ThumbUpAlt";
import ThumbDownAltOutlinedIcon from "@material-ui/icons/ThumbDownAltOutlined";
import ThumbDownIcon from "@material-ui/icons/ThumbDown";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useState, useEffect, useRef } from "react";
import CommentsForPost from "./CommentsForPost.js";
import { Link } from "react-router-dom";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";
import DialogForReport from "./DialogForReport";
import DialogForSaveToFavorites from "./DialogForSaveToFavorites";
import UsersList from "./UsersList";
import Picker from "emoji-picker-react";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import SendContentDialog from "./SendContentDialog.js";

import avatar from "../images/nistagramAvatar.jpg";

const PostDialog = () => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };
  const loggedUserId = localStorage.getItem("id");
  const loggedUsername = localStorage.getItem("username");

  const [newComment, setNewComment] = useState("");

  const { post } = useParams();
  const [user, setUser] = useState();

  const [imagePost, setImagePost] = useState();
  const [descriptionArray, setDescriptionArray] = useState([]);
  const [commentsForPost, setCommentsForPost] = useState([]);
  const [postIsLiked, setPostIsLiked] = useState(false);
  const [postIsDisliked, setPostIsDisliked] = useState(false);
  const [open, setOpen] = useState(false);
  const [openDialogForReport, setOpenDialogForReport] = useState(false);
  const anchorRef = useRef(null);
  const [saveToFavoritesDialog, setSaveToFavoritesDialog] = useState(false);
  const [postSavedToFavourites, setPostSavedToFavourites] = useState(false);
  const [openDialogForLikes, setOpenDialogForLikes] = useState(false);
  const [openDialogForDislikes, setOpenDialogForDislikes] = useState(false);
  const [likers, setLikers] = useState([]);
  const [dislikers, setDislikers] = useState([]);
  const [openPicker, setOpenPicker] = useState(false);
  const [location, setLocation] = useState();
  const [taggedUsers, setTaggedUsers] = useState([]);
  const [openDialogForTaggedUsers, setOpenDialogForTaggedUsers] =
    useState(false);
  const [openSendContentDialog, setOpenSendContentDialog] = useState(false);

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event) => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
  };

  function handleListKeyDown(event) {
    if (event.key === "Tab") {
      event.preventDefault();
      setOpen(false);
    }
  }

  const prevOpen = useRef(open);

  const handleClickPostComment = () => {
    var comment = {
      ID: "0433e24a-d6d2-46f7-9111-33152b10d846",
      PostID: post,
      UserID: loggedUserId,
      CreatedAt: "2018-12-10T13:49:51.141Z",
      Content: newComment,
    };
    axios
      .post("/api/post/add-comment", comment, authorization)
      .then((res) => {
        console.log("upisan komentar");
        let socket = new WebSocket(
          "ws://localhost:8080/api/notification/chat/" + loggedUserId
        );
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send(
            '{"user_who_follow":' +
              '"' +
              loggedUsername +
              '"' +
              ',"command": 2, "channel": ' +
              '"' +
              imagePost.UserID +
              '"' +
              ', "content": "commented your post:"' +
              ', "media": "' +
              imagePost.Media[0] +
              '"' +
              ', "comment": "' +
              newComment +
              '"}' +
              ', "post_id": "' +
              imagePost.ID +
              '"}'
          );
        };
        setNewComment("");
        axios
          .get("/api/post/get-comments-for-post/" + post, authorization)
          .then((res) => {
            setCommentsForPost(res.data);
          })
          .catch((error) => {
            //console.log(error);
          });
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;

    axios.get("/api/post/get-one-post/" + post).then((res) => {
      setImagePost(res.data);
      makeDescriptionFromPost(res.data.Description);
      axios
        .get("/api/user/userid/" + res.data.UserID, authorization)
        .then((res1) => {
          console.log(res1.data);
          setUser(res1.data);
        })
        .catch((error) => {
          //console.log(error);
        });

      console.log(res.data);
    });
    var like = {
      PostID: post,
      UserID: loggedUserId,
    };
    console.log(like);
    axios
      .put("/api/post/like-exists", like, authorization)
      .then((res) => {
        if (res.data == true && res.status === 201) {
          setPostIsLiked(true);
        } else if (res.data == false && res.status === 201) {
          setPostIsLiked(false);
        }
      })
      .catch((error) => {
        //console.log(error);
      });

    axios
      .put("/api/post/dislike-exists", like, authorization)
      .then((res) => {
        if (res.data == true && res.status === 201) {
          setPostIsDisliked(true);
        } else if (res.data == false && res.status === 201) {
          setPostIsDisliked(false);
        }
      })
      .catch((error) => {
        //console.log(error);
      });

    axios
      .get(
        "/api/post/post-exists-in-favourites/" + loggedUserId + "/" + post,
        authorization
      )
      .then((res) => {
        if (res.data == true) {
          setPostSavedToFavourites(true);
        } else if (res.data == false) {
          setPostSavedToFavourites(false);
        }
      })
      .catch((error) => {
        //console.log(error);
      });

    axios
      .get("/api/post/get-comments-for-post/" + post, authorization)
      .then((res) => {
        setCommentsForPost(res.data);
      })
      .catch((error) => {});

    axios
      .get("/api/post/get-location-for-post/" + post, authorization)
      .then((res) => {
        setLocation(res.data.Location);
      })
      .catch((error) => {});

    axios
      .get("/api/post/get-users-tagged-on-post/" + post, authorization)
      .then((res) => {
        if (res.data !== null) {
          setTaggedUsers(res.data);
        }
      })
      .catch((error) => {});
  }, [open]);

  const HandleClickLike = () => {
    var like = {
      PostID: post,
      UserID: loggedUserId,
    };
    axios
      .post("/api/post/like-post", like, authorization)
      .then((res) => {
        if (postIsLiked) {
          if (postIsDisliked) {
            setImagePost({
              ...imagePost,
              LikesCount: Number(imagePost.LikesCount) + Number(1),
              DislikesCount: Number(imagePost.DislikesCount) - Number(1),
            });
            setPostIsDisliked(false);
          } else {
            setImagePost({
              ...imagePost,
              LikesCount: Number(imagePost.LikesCount) - Number(1),
            });
          }
          setPostIsLiked(false);
        } else {
          if (postIsDisliked) {
            setImagePost({
              ...imagePost,
              LikesCount: Number(imagePost.LikesCount) + Number(1),
              DislikesCount: Number(imagePost.DislikesCount) - Number(1),
            });
            setPostIsDisliked(false);
          } else {
            setImagePost({
              ...imagePost,
              LikesCount: Number(imagePost.LikesCount) + Number(1),
            });
          }
          let socket = new WebSocket(
            "ws://localhost:8080/api/notification/chat/" + loggedUserId
          );
          socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send(
              '{"user_who_follow":' +
                '"' +
                loggedUsername +
                '"' +
                ',"command": 2, "channel": ' +
                '"' +
                imagePost.UserID +
                '"' +
                ', "content": "liked your photo."' +
                ', "media": "' +
                imagePost.Media[0] +
                '"' +
                ', "post_id": "' +
                imagePost.ID +
                '"}'
            );
          };
          setPostIsLiked(true);
        }
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const HandleClickDislike = () => {
    var dislike = {
      PostID: post,
      UserID: loggedUserId,
    };
    axios
      .post("/api/post/dislike-post", dislike, authorization)
      .then((res) => {
        if (postIsDisliked) {
          if (postIsLiked) {
            setImagePost({
              ...imagePost,
              DislikesCount: Number(imagePost.DislikesCount) + Number(1),
              LikesCount: Number(imagePost.LikesCount) - Number(1),
            });
            setPostIsLiked(false);
          } else {
            setImagePost({
              ...imagePost,
              DislikesCount: Number(imagePost.DislikesCount) - Number(1),
            });
          }
          setPostIsDisliked(false);
        } else {
          if (postIsLiked) {
            setImagePost({
              ...imagePost,
              DislikesCount: Number(imagePost.DislikesCount) + Number(1),
              LikesCount: Number(imagePost.LikesCount) - Number(1),
            });
            setPostIsLiked(false);
          } else {
            setImagePost({
              ...imagePost,
              DislikesCount: Number(imagePost.DislikesCount) + Number(1),
            });
          }
          let socket = new WebSocket(
            "ws://localhost:8080/api/notification/chat/" + loggedUserId
          );
          socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send(
              '{"user_who_follow":' +
                '"' +
                loggedUsername +
                '"' +
                ',"command": 2, "channel": ' +
                '"' +
                imagePost.UserID +
                '"' +
                ', "content": "disliked your photo."' +
                ', "media": "' +
                imagePost.Media[0] +
                '"' +
                ', "post_id": "' +
                imagePost.ID +
                '"}'
            );
          };
          setPostIsDisliked(true);
        }
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const handleClickAllLikes = () => {
    axios
      .get("/api/post/get-users-who-liked-post/" + post, authorization)
      .then((res) => {
        setLikers(res.data);
        setOpenDialogForLikes(true);
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const handleClickAllDislikes = () => {
    axios
      .get("/api/post/get-users-who-disliked-post/" + post, authorization)
      .then((res) => {
        setDislikers(res.data);
        setOpenDialogForDislikes(true);
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const handleOpenDialogForReport = () => {
    setOpenDialogForReport(true);
    setOpen((prevOpen) => !prevOpen);
  };

  const openSaveToFavoritesDialog = () => {
    setSaveToFavoritesDialog(true);
  };

  const makeDescriptionFromPost = async (text) => {
    console.log(text);
    var resultDescription = [];
    var listOfWords = text.split("#");
    if (listOfWords.length > 0) {
      resultDescription.push(listOfWords[0]);
      for (var i = 1; i < listOfWords.length; i++) {
        var listOfWordsStartWithHash = listOfWords[i].split(" ");
        resultDescription.push("#" + listOfWordsStartWithHash[0]);
        for (var j = 1; j < listOfWordsStartWithHash.length; j++) {
          resultDescription.push(" " + listOfWordsStartWithHash[j]);
        }
      }
    }
    setDescriptionArray(resultDescription);
    console.log(resultDescription);
  };

  const dropDowMenuForPost = (
    <Popper
      open={open}
      anchorEl={anchorRef.current}
      role={undefined}
      transition
      disablePortal
      style={{ width: "15%", zIndex: "1" }}
    >
      {({ TransitionProps, placement }) => (
        <Grow
          {...TransitionProps}
          style={{
            transformOrigin:
              placement === "bottom" ? "center top" : "center bottom",
          }}
        >
          <Paper>
            <ClickAwayListener onClickAway={handleClose}>
              <MenuList
                autoFocusItem={open}
                id="menu-list-grow"
                onKeyDown={handleListKeyDown}
              >
                <MenuItem onClick={handleOpenDialogForReport}>
                  <Grid container>
                    <Grid item xs={3}></Grid>
                    <Grid item xs={9}>
                      <div style={{ width: "100%" }} style={{ color: "red" }}>
                        Report
                      </div>
                    </Grid>
                  </Grid>
                </MenuItem>
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

  const addToComment = (event, emojiObject) => {
    setNewComment(newComment + emojiObject.emoji);
  };

  const Emojis = (
    <>
      {openPicker && (
        <Grid container>
          <Grid item xs={7}></Grid>
          <Grid item xs={5}>
            <Picker onEmojiClick={addToComment} />
          </Grid>
        </Grid>
      )}
    </>
  );

  const handleClickOpenPicker = () => {
    if (openPicker) {
      setOpenPicker(false);
    } else {
      setOpenPicker(true);
    }
  };

  const HandleClickOpenSendContentDialog = () => {
    setOpenSendContentDialog(true);
  };

  function SampleNextArrow(props) {
    const { className, style, onClick } = props;
    return (
      <div
        className={className}
        style={{ ...style, zIndex: 1, right: 0, width: 30, height: 30 }}
        onClick={onClick}
      />
    );
  }

  function SamplePrevArrow(props) {
    const { className, style, onClick } = props;
    return (
      <div
        className={className}
        style={{ ...style, zIndex: 1, left: 0 }}
        onClick={onClick}
      />
    );
  }

  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    nextArrow: <SampleNextArrow />,
    prevArrow: <SamplePrevArrow />,
  };

  return (
    <div>
      <Grid container>
        <Grid item xs={2}></Grid>
        <Grid item xs={8}>
          <Paper
            style={{
              width: "100%",
              height: "600px",
              marginTop: "5%",
            }}
            variant="outlined"
          >
            <Grid container style={{ width: "100%", height: "100%" }}>
              {imagePost !== undefined && imagePost !== null && (
                <Grid item xs={7}>
                  <Slider {...settings}>
                    {imagePost.Media.map((media, index) => (
                      <div
                        style={{ width: "100%", height: "100%" }}
                        key={index}
                      >
                        {media.substring(media.length - 3, media.length) ===
                          "jpg" && (
                          <img
                            src={
                              "http://localhost:8080/api/media/get-media-image/" +
                              media
                            }
                            style={{ width: "100%", height: "600px" }}
                          />
                        )}
                        {media.substring(media.length - 3, media.length) !==
                          "jpg" && (
                          <video
                            width="100%"
                            height="100%"
                            style={{ marginTop: "25%" }}
                            controls
                          >
                            <source
                              src={
                                "http://localhost:8080/api/media/get-video/" +
                                media
                              }
                              style={{ width: "100%", height: "100%" }}
                              type="video/mp4"
                            />
                          </video>
                        )}
                        {imagePost.IsCampaign === true &&
                          imagePost.Links[index] !== "" && (
                            <Button
                              variant="contained"
                              color="primary"
                              style={{ width: "100%", height: "40px" }}
                            >
                              <a
                                href={imagePost.Links[index]}
                                style={{
                                  textDecoration: "none",
                                  color: "white",
                                }}
                              >
                                Visit on web site
                              </a>
                            </Button>
                          )}
                      </div>
                    ))}
                  </Slider>
                </Grid>
              )}

              <Grid item xs={5}>
                <Grid container>
                  <Grid
                    item
                    xs={3}
                    style={{
                      margin: "auto",
                      marginTop: "2%",
                      textAlign: "center",
                    }}
                  >
                    {user !== undefined &&
                    user !== null &&
                    user.ProfilePicture !== "" ? (
                      <Avatar
                        alt="N"
                        style={{ margin: "auto", border: "1px solid black" }}
                        src={
                          "http://localhost:8080/api/media/get-profile-picture/" +
                          user.ProfilePicture
                        }
                      ></Avatar>
                    ) : (
                      <Avatar
                        style={{ margin: "auto", border: "1px solid black" }}
                        alt="N"
                        src={avatar}
                      ></Avatar>
                    )}
                  </Grid>
                  <Grid
                    item
                    xs={7}
                    style={{ margin: "auto", textAlign: "left" }}
                  >
                    {user !== undefined && user !== null && (
                      <Link
                        to={"/homePage/" + user.Username}
                        style={{ textDecoration: "none", color: "black" }}
                      >
                        <b>{user.Username}</b>
                      </Link>
                    )}
                    <br />
                    {location && (
                      <Link
                        to={"/explore/locations/" + location + "/"}
                        style={{ textDecoration: "none", color: "gray" }}
                      >
                        {location ? location : ""}
                      </Link>
                    )}
                  </Grid>
                  <Grid item xs={2}>
                    {imagePost && loggedUserId !== imagePost.UserID && (
                      <MoreHorizIcon
                        style={{
                          marginTop: "50%",
                          textAlign: "right",
                          cursor: "pointer",
                        }}
                        aria-controls={open ? "menu-list-grow" : undefined}
                        aria-haspopup="true"
                        ref={anchorRef}
                        onClick={handleToggle}
                      ></MoreHorizIcon>
                    )}
                    {dropDowMenuForPost}
                  </Grid>
                </Grid>

                <Grid
                  container
                  style={{
                    height: "10%",
                    overflow: "auto",
                    paddingLeft: "8%",
                    paddingRight: "5%",
                    paddingTop: "1%",
                  }}
                >
                  {imagePost && imagePost.Description.length !== 0 && (
                    <label style={{ textAlign: "left" }}>
                      {descriptionArray.map((word, index) =>
                        word.charAt(0) === "#" ? (
                          <Link
                            to={`/explore/tags/${word.substring(1)}/`}
                            style={{ textDecoration: "none" }}
                            key={index}
                          >
                            {word}
                          </Link>
                        ) : (
                          `${word}`
                        )
                      )}
                    </label>
                  )}
                </Grid>

                <Grid container style={{ height: "335px", overflow: "auto" }}>
                  <CommentsForPost comments={commentsForPost}></CommentsForPost>
                </Grid>

                <Grid container style={{ height: "25%" }}>
                  <Grid container style={{ height: "70%" }}>
                    <Grid container style={{ height: "30%" }}>
                      <Grid item xs={2}>
                        <Divider />
                        {postIsLiked ? (
                          <ThumbUpAltIcon
                            onClick={HandleClickLike}
                            fontSize="large"
                            style={{ cursor: "pointer" }}
                          ></ThumbUpAltIcon>
                        ) : (
                          <ThumbUpAltOutlinedIcon
                            onClick={HandleClickLike}
                            fontSize="large"
                            style={{ cursor: "pointer" }}
                          ></ThumbUpAltOutlinedIcon>
                        )}
                      </Grid>
                      <Grid item xs={2}>
                        <Divider />

                        {postIsDisliked ? (
                          <ThumbDownIcon
                            onClick={HandleClickDislike}
                            fontSize="large"
                            style={{ cursor: "pointer" }}
                          ></ThumbDownIcon>
                        ) : (
                          <ThumbDownAltOutlinedIcon
                            onClick={HandleClickDislike}
                            fontSize="large"
                            style={{ cursor: "pointer" }}
                          ></ThumbDownAltOutlinedIcon>
                        )}
                      </Grid>

                      <Grid item xs={2}>
                        <Divider />
                        <SendRounded
                          fontSize="large"
                          style={{ cursor: "pointer" }}
                          onClick={HandleClickOpenSendContentDialog}
                        ></SendRounded>
                      </Grid>
                      <Grid item xs={2}>
                        <Divider />
                      </Grid>
                      <Grid item xs={2}>
                        <Divider />
                        {taggedUsers.length !== 0 && (
                          <PersonRounded
                            fontSize="large"
                            style={{ margin: "auto", cursor: "pointer" }}
                            onClick={() => setOpenDialogForTaggedUsers(true)}
                          />
                        )}
                      </Grid>
                      <Grid item xs={2}>
                        <Divider />
                        {postSavedToFavourites ? (
                          <BookmarkRounded
                            onClick={openSaveToFavoritesDialog}
                            style={{ margin: "auto", cursor: "pointer" }}
                            fontSize="large"
                          ></BookmarkRounded>
                        ) : (
                          <BookmarkBorderRounded
                            onClick={openSaveToFavoritesDialog}
                            style={{ margin: "auto", cursor: "pointer" }}
                            fontSize="large"
                          ></BookmarkBorderRounded>
                        )}
                      </Grid>
                    </Grid>
                    <Grid container style={{ height: "70%" }}>
                      <Grid item xs={5}>
                        {imagePost !== undefined && imagePost !== null && (
                          <h5
                            onClick={handleClickAllLikes}
                            style={{ cursor: "pointer" }}
                          >
                            {imagePost.LikesCount} Likes
                          </h5>
                        )}
                      </Grid>
                      <Grid item xs={5}>
                        {imagePost !== undefined && imagePost !== null && (
                          <h5
                            onClick={handleClickAllDislikes}
                            style={{ cursor: "pointer" }}
                          >
                            {imagePost.DislikesCount} Dislikes
                          </h5>
                        )}
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
                        onClick={handleClickOpenPicker}
                      ></SentimentSatisfiedRoundedIcon>
                    </Grid>
                    <Grid item xs={7}>
                      <Divider />
                      <InputBase
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
          </Paper>
          {Emojis}
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
      <Grid container style={{ marginBottom: "2%" }} />

      {openDialogForReport && (
        <DialogForReport
          loggedUserId={loggedUserId}
          post={post}
          open={openDialogForReport}
          setOpen={setOpenDialogForReport}
        ></DialogForReport>
      )}

      {saveToFavoritesDialog && (
        <DialogForSaveToFavorites
          loggedUserId={loggedUserId}
          post={post}
          open={saveToFavoritesDialog}
          setOpen={setSaveToFavoritesDialog}
          saved={postSavedToFavourites}
          setSaved={setPostSavedToFavourites}
        ></DialogForSaveToFavorites>
      )}

      {openSendContentDialog && (
        <SendContentDialog
          open={openSendContentDialog}
          setOpen={setOpenSendContentDialog}
          userForPost={imagePost.UserID}
          postId={imagePost.ID}
        ></SendContentDialog>
      )}

      {openDialogForLikes && (
        <UsersList
          label="People who like post"
          users={likers}
          open={openDialogForLikes}
          setOpen={setOpenDialogForLikes}
        ></UsersList>
      )}

      {openDialogForDislikes && (
        <UsersList
          label="People who dislike post"
          users={dislikers}
          open={openDialogForDislikes}
          setOpen={setOpenDialogForDislikes}
        ></UsersList>
      )}

      {openDialogForTaggedUsers && (
        <UsersList
          label="People tagged on post"
          users={taggedUsers}
          open={openDialogForTaggedUsers}
          setOpen={setOpenDialogForTaggedUsers}
        ></UsersList>
      )}
    </div>
  );
};

export default PostDialog;
