import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import avatar from "../images/nistagramAvatar.jpg";

import { Avatar, Grid } from "@material-ui/core";

const CommentOneForFeed = ({ comment }) => {
  const [profilePicture, setProfilePicture] = useState("");
  const [username, setUsername] = useState();

  useEffect(() => {
    axios
      .get("/api/user/find-username-and-profile-picture/" + comment.UserID)
      .then((res) => {
        setUsername(res.data.Username);
        setProfilePicture(res.data.ProfilePicture);
      }).catch((error) => {
        //console.log(error);
      });;
  }, []);

  return (
    <div>
      {comment !== undefined && comment !== null && (
        <Grid container style={{ marginTop: "1%" }}>
          <Grid item xs={1}>
            {profilePicture !== null && profilePicture !== undefined && (
              <Avatar
                alt="N"
                src={
                  profilePicture === ""
                    ? avatar
                    : "http://localhost:8080/api/media/get-profile-picture/" +
                      profilePicture
                }
                style={{
                  border: "0.5px solid",
                  margin: "auto",
                  width: "25px",
                  height: "25px",
                  borderColor: "#b9b9b9",
                }}
              ></Avatar>
            )}
          </Grid>
          <Grid item xs={11}>
            <label>
              <Link
                to={"/homePage/" + username}
                style={{ textDecoration: "none", color: "black" }}
              >
                <b>{username}</b>
              </Link>{" "}
              {comment.Content}
            </label>
            <br />
          </Grid>
        </Grid>
      )}
    </div>
  );
};

export default CommentOneForFeed;
