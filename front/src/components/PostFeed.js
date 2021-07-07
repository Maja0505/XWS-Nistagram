import {
  Grid,
  Paper,
  Avatar,
  Box,
  FormLabel,
  Button,
  InputBase,
  Grow,
  Popper,
  MenuItem,
  MenuList,
  ClickAwayListener,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import avatar from "../images/nistagramAvatar.jpg";
import { useState, useEffect, useRef } from "react";
import {
  MoreHoriz,
  ThumbUpOutlined,
  ThumbUp,
  ThumbDownOutlined,
  ThumbDown,
  SendRounded,
  BookmarkBorderRounded,
  BookmarkRounded,
  SentimentSatisfiedRounded,
  PersonRounded,
} from "@material-ui/icons";
import UsersList from "./UsersList";
import CommentsForFeeds from "./CommentsForFeeds.js";
import DialogForReport from "./DialogForReport";
import DialogForSaveToFavorites from "./DialogForSaveToFavorites";
import Slider from "react-slick";
import Picker from "emoji-picker-react";

import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

import axios from "axios";

const PostFeed = ({ feed }) => {
  const [username, setUsername] = useState();
  const [profileImage, setProfileImage] = useState();
  const [descriptionArray, setDescriptionArray] = useState([]);
  const loggedUserId = localStorage.getItem("id");
  const loggedUsername = localStorage.getItem("username");

  const [isLiked, setIsLiked] = useState(false);
  const [isDisliked, setIsDisliked] = useState(false);
  const [isSaved, setIsSaved] = useState(false);
  const [copyOfFeed, setCopyOfFeed] = useState();
  const [newComment, setNewComment] = useState("");
  const [showComments, setShowComments] = useState(false);
  const [comments, setComments] = useState([]);
  const [locationForFeed, setLocationForFeed] = useState();

  const [openDialogForLikes, setOpenDialogForLikes] = useState(false);
  const [likers, setLikers] = useState([]);
  const [openDialogForDislikes, setOpenDialogForDislikes] = useState(false);
  const [dislikers, setDislikers] = useState([]);
  const [taggedUsers, setTaggedUsers] = useState([]);

  const [openPicker, setOpenPicker] = useState(false);
  const [openDialogForReport, setOpenDialogForReport] = useState(false);
  const [saveToFavoritesDialog, setSaveToFavoritesDialog] = useState(false);
  const [openDialogForTaggedUsers, setOpenDialogForTaggedUsers] =
    useState(false);

  const [open, setOpen] = useState(false);
  const anchorRef = useRef(null);

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleOpenDialogForReport = () => {
    setOpenDialogForReport(true);
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

  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    nextArrow: <SampleNextArrow />,
    prevArrow: <SamplePrevArrow />,
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

  const makeDescriptionFromPost = (text) => {
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
  };

  useEffect(() => {
    setCopyOfFeed(feed);

    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;

    axios
      .put("/api/post/like-exists", { PostID: feed.ID, UserID: loggedUserId })
      .then((res) => {
        setIsLiked(res.data);
      });

    axios
      .put("/api/post/dislike-exists", {
        PostID: feed.ID,
        UserID: loggedUserId,
      })
      .then((res) => {
        setIsDisliked(res.data);
      });

    axios
      .get(
        "/api/post/post-exists-in-favourites/" + loggedUserId + "/" + feed.ID
      )
      .then((res) => {
        setIsSaved(res.data);
      });

    makeDescriptionFromPost(feed.Description);
    axios
      .get("/api/user/find-username-and-profile-picture/" + feed.UserID)
      .then((res) => {
        setUsername(res.data.Username);
        setProfileImage(res.data.ProfilePicture);
      });

    axios.get("/api/post/get-location-for-post/" + feed.ID).then((res) => {
      setLocationForFeed(res.data.Location);
    });

    axios.get("/api/post/get-users-tagged-on-post/" + feed.ID).then((res) => {
      if (res.data !== null) {
        setTaggedUsers(res.data);
      }
    });
  }, [feed]);

  const likePost = () => {
    axios
      .post("/api/post/like-post", { PostID: feed.ID, UserID: loggedUserId })
      .then((res) => {
        if (!isLiked) {
          if (isDisliked) {
            setIsDisliked(!isDisliked);
            setCopyOfFeed({
              ...copyOfFeed,
              DislikesCount: Number(copyOfFeed.DislikesCount) - Number(1),
              LikesCount: Number(copyOfFeed.LikesCount) + Number(1),
            });
          } else {
            setCopyOfFeed({
              ...copyOfFeed,
              LikesCount: Number(copyOfFeed.LikesCount) + Number(1),
            });
          }
          let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
          socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + feed.UserID + '"' + ', "content": "liked your photo."' + ', "post_id": "' + feed.ID + '"}')
          };
        } else {
          if (isDisliked) {
            setIsDisliked(!isDisliked);
            setCopyOfFeed({
              ...copyOfFeed,
              DislikesCount: Number(copyOfFeed.DislikesCount) - Number(1),
              LikesCount: Number(copyOfFeed.LikesCount) - Number(1),
            });
          } else {
            setCopyOfFeed({
              ...copyOfFeed,
              LikesCount: Number(copyOfFeed.LikesCount) - Number(1),
            });
          }
        }
        setIsLiked(!isLiked);
      });
  };

  const dislikePost = () => {
    axios
      .post("/api/post/dislike-post", { PostID: feed.ID, UserID: loggedUserId })
      .then((res) => {
        if (!isDisliked) {
          if (isLiked) {
            setIsLiked(!isLiked);
            setCopyOfFeed({
              ...copyOfFeed,
              LikesCount: Number(copyOfFeed.LikesCount) - Number(1),
              DislikesCount: Number(copyOfFeed.DislikesCount) + Number(1),
            });
          } else {
            setCopyOfFeed({
              ...copyOfFeed,
              DislikesCount: Number(copyOfFeed.DislikesCount) + Number(1),
            });
          }
          let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
          socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + feed.UserID + '"' + ', "content": "disliked your photo."' + ', "post_id": "' + feed.ID + '"}')
          };
        } else {
          if (isLiked) {
            setIsLiked(!isLiked);
            setCopyOfFeed({
              ...copyOfFeed,
              LikesCount: Number(copyOfFeed.LikesCount) - Number(1),
              DislikesCount: Number(copyOfFeed.DislikesCount) - Number(1),
            });
          } else {
            setCopyOfFeed({
              ...copyOfFeed,
              DislikesCount: Number(copyOfFeed.DislikesCount) - Number(1),
            });
          }
        }
        setIsDisliked(!isDisliked);
      });
  };

  const addComment = () => {
    axios
      .post("/api/post/add-comment", {
        PostID: feed.ID,
        UserID: loggedUserId,
        Content: newComment,
      })
      .then((res) => {
        let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + feed.UserID + '"' + ', "content": "commented your post:"' + ', "post_id": "' + feed.ID + '"' + ', "comment": "' + newComment + '"}')
        };
        axios.get("/api/post/get-comments-for-post/" + feed.ID).then((res) => {
          setComments(res.data);
        });
        setCopyOfFeed({
          ...copyOfFeed,
          CommentsCount: Number(copyOfFeed.CommentsCount) + Number(1),
        });
        setNewComment("");
      });
  };

  const viewComments = () => {
    axios.get("/api/post/get-comments-for-post/" + feed.ID).then((res) => {
      setComments(res.data);
      setShowComments(true);
    });
  };

  const addToComment = (event, emojiObject) => {
    setNewComment(newComment + emojiObject.emoji);
  };

  const getUsersWhoLikedPost = () => {
    axios.get("/api/post/get-users-who-liked-post/" + feed.ID).then((res) => {
      setLikers(res.data);
      setOpenDialogForLikes(true);
    });
  };

  const getUsersWhoDislikedPost = () => {
    axios
      .get("/api/post/get-users-who-disliked-post/" + feed.ID)
      .then((res) => {
        setDislikers(res.data);
        setOpenDialogForDislikes(true);
      });
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

  const Emojis = (
    <>
      {openPicker && (
        <Grid container>
          <Grid item xs={5}>
            <Picker onEmojiClick={addToComment} />
          </Grid>
          <Grid item xs={7}></Grid>
        </Grid>
      )}
    </>
  );

  return (
    <div>
      {copyOfFeed !== undefined && copyOfFeed !== null && (
        <>
          <Paper
            style={{
              border: "solid 0.05px",
              borderRadius: "0%",
              borderColor: "#b9b9b9",
              height: "60px",
              borderBottom: "0px",
            }}
          >
            <Grid container style={{ height: "100%" }}>
              <Grid item xs={2} style={{ margin: "auto" }}>
                {profileImage !== null && profileImage !== undefined && (
                  <Avatar
                    alt="N"
                    src={
                      profileImage === ""
                        ? avatar
                        : "http://localhost:8080/api/media/get-profile-picture/" +
                          profileImage
                    }
                    style={{
                      border: "0.5px solid",
                      margin: "auto",
                      width: "33px",
                      height: "33px",
                      borderColor: "#b9b9b9",
                    }}
                  ></Avatar>
                )}
              </Grid>
              <Grid item xs={8} style={{ textAlign: "left", margin: "auto" }}>
                <Link
                  to={"/homePage/" + username}
                  style={{ textDecoration: "none", color: "black" }}
                >
                  <b>{username}</b>
                  <br />
                </Link>
                {locationForFeed && (
                  <Link
                    to={"/explore/locations/" + locationForFeed + "/"}
                    style={{ textDecoration: "none", color: "black" }}
                  >
                    {locationForFeed ? locationForFeed : ""}
                  </Link>
                )}
              </Grid>
              <Grid item xs={2} style={{ margin: "auto" }}>
                <MoreHoriz
                  style={{ cursor: "pointer" }}
                  aria-controls={open ? "menu-list-grow" : undefined}
                  aria-haspopup="true"
                  ref={anchorRef}
                  onClick={handleToggle}
                />
                {dropDowMenuForPost}
              </Grid>
            </Grid>
          </Paper>

          <Box
            style={{
              border: "0.2px solid",
              borderColor: "#b9b9b9",
              borderTop: "0px",
              borderBottom: "0px",
            }}
          >
            <Slider {...settings}>
              {copyOfFeed.Media.map((media, index) => (
                <div style={{ width: "100%", height: "100%" }} key={index}>
                  {media.substring(media.length - 3, media.length) ===
                    "jpg" && (
                    <img
                      src={
                        "http://localhost:8080/api/media/get-media-image/" +
                        media
                      }
                      style={{
                        width: "100%",
                        height: copyOfFeed.Media.length > 1 ? "470px" : "100%",
                      }}
                    />
                  )}
                  {media.substring(media.length - 3, media.length) !==
                    "jpg" && (
                    <video width="100%" height="100%" controls>
                      <source
                        src={
                          "http://localhost:8080/api/media/get-video/" + media
                        }
                        style={{
                          width: "100%",
                          height: "100%",
                          margin: "auto",
                        }}
                        type="video/mp4"
                      />
                    </video>
                  )}
                </div>
              ))}
            </Slider>
          </Box>

          <Paper
            style={{
              border: "solid 0.05px",
              borderRadius: "0%",
              borderColor: "#b9b9b9",
              borderBottom: "0px",
            }}
          >
            <Grid container style={{ marginTop: "2%" }}>
              <Grid container item xs={4}>
                <Grid item xs={1} />
                <Grid item xs={3}>
                  {isLiked ? (
                    <ThumbUp
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={likePost}
                    />
                  ) : (
                    <ThumbUpOutlined
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={likePost}
                    />
                  )}
                </Grid>
                <Grid item xs={3}>
                  {isDisliked ? (
                    <ThumbDown
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={dislikePost}
                    />
                  ) : (
                    <ThumbDownOutlined
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={dislikePost}
                    />
                  )}
                </Grid>
                <Grid item xs={3}>
                  <SendRounded
                    fontSize="large"
                    style={{ margin: "auto", cursor: "pointer" }}
                  />
                </Grid>
              </Grid>
              <Grid item xs={4} />
              <Grid container item xs={4}>
                <Grid item xs={2} />
                <Grid item xs={3} />
                <Grid item xs={3}>
                  {taggedUsers.length !== 0 && (
                    <PersonRounded
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={() => setOpenDialogForTaggedUsers(true)}
                    />
                  )}
                </Grid>
                <Grid item xs={3}>
                  {isSaved ? (
                    <BookmarkRounded
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={() => setSaveToFavoritesDialog(true)}
                    />
                  ) : (
                    <BookmarkBorderRounded
                      fontSize="large"
                      style={{ margin: "auto", cursor: "pointer" }}
                      onClick={() => setSaveToFavoritesDialog(true)}
                    />
                  )}
                </Grid>
              </Grid>
            </Grid>

            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel
                style={{ fontSize: "14px", cursor: "pointer" }}
                onClick={getUsersWhoLikedPost}
              >
                <b>{copyOfFeed.LikesCount} likes</b>
              </FormLabel>
            </Grid>
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel
                style={{ fontSize: "14px", cursor: "pointer" }}
                onClick={getUsersWhoDislikedPost}
              >
                <b>{copyOfFeed.DislikesCount} dislikes</b>
              </FormLabel>
            </Grid>
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              {copyOfFeed.Description.length !== 0 && (
                <label style={{ textAlign: "left" }}>
                  <Link
                    to={"/homePage/" + username}
                    style={{ textDecoration: "none", color: "black" }}
                  >
                    <b>{username}</b>
                  </Link>
                  {"  "}
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
            {copyOfFeed.CommentsCount > 0 && (
              <div>
                <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
                  <FormLabel
                    style={{ fontSize: "13px" }}
                    style={{ cursor: "pointer" }}
                    onClick={viewComments}
                  >
                    View all {copyOfFeed.CommentsCount} comments
                  </FormLabel>
                </Grid>
                <>{showComments && <CommentsForFeeds comments={comments} />}</>
              </div>
            )}
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel style={{ fontSize: "13px" }}>
                POSTED {copyOfFeed.CreatedAt.split("-")[2].substring(0, 2)}
                {"."}
                {copyOfFeed.CreatedAt.split("-")[1]}
                {"."}
                {copyOfFeed.CreatedAt.split("-")[0]}
                {". AT "}
                {copyOfFeed.CreatedAt.split("T")[1].substring(0, 5)}
                {" H"}
              </FormLabel>
            </Grid>
            <Grid container style={{ marginTop: "1.5%" }}></Grid>
          </Paper>

          <Paper
            style={{
              border: "solid 0.05px",
              borderRadius: "0%",
              borderColor: "#b9b9b9",
            }}
          >
            <Grid container style={{ height: "50px" }}>
              <Grid item xs={1} style={{ margin: "auto" }}>
                <SentimentSatisfiedRounded
                  style={{
                    margin: "auto",
                    marginTop: "5%",
                    cursor: "pointer",
                  }}
                  fontSize="large"
                  onClick={() => setOpenPicker(!openPicker)}
                />
              </Grid>
              <Grid item xs={9} style={{ margin: "auto" }}>
                <InputBase
                  placeholder="Add a comment..."
                  inputProps={{ "aria-label": "naked" }}
                  style={{ width: "100%", textAlign: "left" }}
                  value={newComment}
                  onChange={(e) => setNewComment(e.target.value)}
                />
              </Grid>
              <Grid item xs={1} style={{ margin: "auto", textAlign: "right" }}>
                <Button
                  variant="text"
                  color="primary"
                  onClick={addComment}
                  disabled={newComment !== "" ? false : true}
                >
                  POST
                </Button>
              </Grid>
              {Emojis}
            </Grid>
          </Paper>
        </>
      )}

      {saveToFavoritesDialog && (
        <DialogForSaveToFavorites
          loggedUserId={loggedUserId}
          post={feed.ID}
          open={saveToFavoritesDialog}
          setOpen={setSaveToFavoritesDialog}
          saved={isSaved}
          setSaved={setIsSaved}
        ></DialogForSaveToFavorites>
      )}

      {openDialogForReport && (
        <DialogForReport
          loggedUserId={loggedUserId}
          post={feed.ID}
          open={openDialogForReport}
          setOpen={setOpenDialogForReport}
        ></DialogForReport>
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

export default PostFeed;
