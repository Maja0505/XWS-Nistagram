import { Grid, Button } from "@material-ui/core";
import { Link } from "react-router-dom";
import { useState, useEffect } from "react";
import Slider from "react-slick";

import axios from "axios";

const ReportedContentOne = ({
  content,
  setReportedContents,
  reportedContents,
}) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const [post, setPost] = useState();
  const [userWichReportedPost, setUserWichReportePost] = useState("");
  const [userWichPosted, setUserWichPosted] = useState("");

  useEffect(() => {
    axios
      .get("/api/post/get-one-post/" + content.Image, authorization)
      .then((res) => {
        if (res.data !== null) {
          console.log(res.data);
          setPost(res.data);
          axios
            .get("/api/user/userid/" + res.data.UserID, authorization)
            .then((res1) => {
              setUserWichPosted(res1.data.Username);
            })
            .catch((error) => {
              //console.log(error);
            });
        }
      })
      .catch((error) => {
        //console.log(error);
      });

    axios
      .get("/api/user/userid/" + content.UserID, authorization)
      .then((res) => {
        setUserWichReportePost(res.data.Username);
      });
  }, [reportedContents]);

  const deletePost = () => {
    axios
      .put(
        "/api/post/delete-post/" + post.ID + "/" + post.UserID,
        {},
        authorization
      )
      .then((res) => {
        deleteReportedContent();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const deleteUser = () => {
    axios
      .put("/api/user/delete/" + post.UserID, {}, authorization)
      .then((res) => {
        deletePost();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const deleteReportedContent = () => {
    axios
      .put(
        "/api/post/delete-reported-content/" +
          content.ID +
          "/" +
          content.UserID,
        {},
        authorization
      )
      .then((res) => {
        deleteFromArray();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const deleteFromArray = () => {
    var array = [...reportedContents];
    var index = array.indexOf(content);
    console.log(array);
    if (index !== -1) {
      array.splice(index, 1);
    }
    setReportedContents(array);
  };

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

  return (
    <div>
      {post !== null && post !== undefined && (
        <Grid container>
          <Grid item xs={6}>
            <Slider {...settings}>
              {post.Media.map((media, index) => (
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
                        height: "200px",
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
                          height: "200px",
                          margin: "auto",
                        }}
                        type="video/mp4"
                      />
                    </video>
                  )}
                </div>
              ))}
            </Slider>
          </Grid>

          <Grid item xs={3} style={{ margin: "auto" }}>
            <Grid container>
              <label style={{ margin: "auto" }}>
                <Link
                  to={"/homePage/" + userWichReportedPost}
                  style={{ textDecoration: "none", color: "black" }}
                >
                  <b>{userWichReportedPost}</b>
                </Link>{" "}
                reported{" "}
                <Link
                  to={"/homePage/" + userWichPosted}
                  style={{ textDecoration: "none", color: "black" }}
                >
                  <b>{userWichPosted}</b>
                </Link>
                's post
              </label>
            </Grid>

            <Grid container>
              <label style={{ marginLeft: "23%" }}>
                {post.LikesCount} likes
              </label>
            </Grid>

            <Grid container>
              <label style={{ marginLeft: "23%" }}>
                {post.DislikesCount} dislikes
              </label>
            </Grid>

            <Grid container>
              <label style={{ marginLeft: "23%" }}>
                {post.CommentsCount} comments
              </label>
            </Grid>

            <Grid container>
              <label style={{ margin: "auto", overflow: "auto" }}>
                {content.Description}
              </label>
            </Grid>
          </Grid>

          <Grid item xs={3} style={{ margin: "auto" }}>
            <Grid container>
              <Button
                color="inherit"
                variant="text"
                style={{ margin: "auto" }}
                onClick={deletePost}
              >
                Delete Post
              </Button>
            </Grid>
            <Grid container style={{ marginTop: "1%" }}>
              <Button
                color="secondary"
                variant="text"
                style={{ margin: "auto" }}
                onClick={deleteUser}
              >
                Delete User
              </Button>
            </Grid>
            <Grid container style={{ marginTop: "1%" }}>
              <Button
                color="primary"
                variant="text"
                style={{ margin: "auto" }}
                onClick={deleteReportedContent}
              >
                Delete Report
              </Button>
            </Grid>
          </Grid>
        </Grid>
      )}
    </div>
  );
};

export default ReportedContentOne;
