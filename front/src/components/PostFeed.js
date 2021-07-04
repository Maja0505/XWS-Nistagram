import {
  Grid,
  Paper,
  Avatar,
  Box,
  FormLabel,
  Button,
  InputBase,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import avatar from "../images/nistagramAvatar.jpg";
import { useState, useEffect } from "react";
import {
  MoreHoriz,
  ThumbUpOutlined,
  ThumbDownAltOutlined,
  SendRounded,
  BookmarkBorderRounded,
  SentimentSatisfiedRounded,
} from "@material-ui/icons";

import Slider from "react-slick";

import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

import axios from "axios";

const PostFeed = ({ feed }) => {
  const [username, setUsername] = useState();
  const [profileImage, setProfileImage] = useState();
  const [descriptionArray, setDescriptionArray] = useState([]);

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
    makeDescriptionFromPost(feed.Description);
    axios
      .get("/api/user/find-username-and-profile-picture/" + feed.UserID)
      .then((res) => {
        setUsername(res.data.Username);
        setProfileImage(res.data.ProfilePicture);
      });
  }, [feed]);

  return (
    <div>
      {feed !== undefined && feed !== null && (
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
                        : "http://localhost:8080/api/user/get-image/" +
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
                {"ovde lokacija ako je ima"}
              </Grid>
              <Grid item xs={2} style={{ margin: "auto" }}>
                <MoreHoriz />
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
            <div>
              <Slider {...settings}>
                {feed.Media.map((media, index) => (
                  <div style={{ width: "100%", height: "100%" }} key={index}>
                    {media.substring(media.length - 3, media.length) ===
                      "jpg" && (
                      <img
                        src={
                          "http://localhost:8080/api/post/get-image/" + media
                        }
                        style={{ width: "100%", height: "100%" }}
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
                            "http://localhost:8080/api/post/get-image/" + media
                          }
                          style={{ width: "100%", height: "100%" }}
                          type="video/mp4"
                        />
                      </video>
                    )}
                  </div>
                ))}
              </Slider>
            </div>
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
                  <ThumbUpOutlined
                    fontSize="large"
                    style={{ margin: "auto" }}
                  />
                </Grid>
                <Grid item xs={3}>
                  <ThumbDownAltOutlined
                    fontSize="large"
                    style={{ margin: "auto" }}
                  />
                </Grid>
                <Grid item xs={3}>
                  <SendRounded fontSize="large" style={{ margin: "auto" }} />
                </Grid>
              </Grid>
              <Grid item xs={4} />
              <Grid container item xs={4}>
                <Grid item xs={2} />
                <Grid item xs={3} />
                <Grid item xs={3} />
                <Grid item xs={3}>
                  <BookmarkBorderRounded
                    fontSize="large"
                    style={{ margin: "auto" }}
                  />
                </Grid>
              </Grid>
            </Grid>

            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel style={{ fontSize: "14px" }}>
                <b>{feed.LikesCount} likes</b>
              </FormLabel>
            </Grid>
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel style={{ fontSize: "14px" }}>
                <b>{feed.DislikesCount} dislikes</b>
              </FormLabel>
            </Grid>
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              {feed.Description.length !== 0 && (
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
            {feed.CommentsCount > 0 && (
              <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
                <FormLabel style={{ fontSize: "13px" }}>
                  View all {feed.CommentsCount} comments
                </FormLabel>
              </Grid>
            )}
            <Grid container style={{ marginTop: "1%", marginLeft: "4.4%" }}>
              <FormLabel style={{ fontSize: "13px" }}>
                POSTED {feed.CreatedAt.split("-")[2].substring(0, 2)}
                {"."}
                {feed.CreatedAt.split("-")[1]}
                {"."}
                {feed.CreatedAt.split("-")[0]}
                {". AT "}
                {feed.CreatedAt.split("T")[1].substring(0, 5)}
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
                />
              </Grid>
              <Grid item xs={9} style={{ margin: "auto" }}>
                <InputBase
                  placeholder="Add a comment..."
                  inputProps={{ "aria-label": "naked" }}
                  style={{ width: "100%", textAlign: "left" }}
                />
              </Grid>
              <Grid item xs={1} style={{ margin: "auto", textAlign: "right" }}>
                <Button variant="text" color="primary">
                  POST
                </Button>
              </Grid>
            </Grid>
          </Paper>
        </>
      )}
    </div>
  );
};

export default PostFeed;
