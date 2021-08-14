import { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";

import { Grid, CircularProgress } from "@material-ui/core";

import UserDetailsOnHomePage from "./UserDetailsOnHomePage";
import PostDetailsOnHomePage from "./PostDetailsOnHomePage";
import UserHighlightsOnHomePage from "./UserHighlightsOnHomePage";

const UserHomePage = () => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const { username } = useParams();
  const loggedInId = localStorage.getItem("id");
  const logedInUsername = localStorage.getItem("username");

  const [relationShip, setRelationShip] = useState(null);
  const [user, setUser] = useState(null);
  const [allFollows, setAllFollows] = useState(null);
  const [allFollowers, setAllFollowers] = useState(null);
  const [highlightStories, setHighlightStories] = useState(null);

  useEffect(() => {
    console.log(username);
    console.log(loggedInId);

    setAllFollowers(null);

    axios
      .get("/api/user/" + username, authorization)
      .then((res) => {
        setUser(res.data);
        console.log(res.data);
        
        if (logedInUsername !== username) {
          console.log('usao u if')
          axios
            .get(
              "/api/user-follow/relationship/" +
                loggedInId +
                "/" +
                res.data.IdString,
              authorization
            )
            .then((res1) => {
              console.log(res1.data)
              setRelationShip(res1.data);
            })
            .catch((error) => {
              console.log(error);
            });
        }

        axios
          .get(
            "/api/post/story/all-highlights/" + res.data.IdString,
            authorization
          )
          .then((res) => {
            if (res.data) {
              setHighlightStories(res.data);
            }
          })
          .catch((error) => {
            console.log(error);
          });

        axios
          .get(
            "/api/user-follow/allFollows/" + res.data.IdString,
            authorization
          )
          .then((res) => {
            if (res.data) {
              setAllFollows(res.data);
            } else {
              setAllFollows([]);
            }
          })
          .catch((error) => {
            setAllFollows([]);
          });

        axios
          .get(
            "/api/user-follow/allFollowers/" + res.data.IdString,
            authorization
          )
          .then((res) => {
            if (res.data) {
              setAllFollowers(res.data);
            } else {
              setAllFollowers([]);
            }
          })
          .catch((error) => {
            setAllFollowers([]);
          });
      })
      .catch((error) => {
        console.log(error);
      });
  }, [username]);

  return (
    <div>
      {(user === null || allFollowers === null || allFollows === null) && (
        <div style={{ margin: "auto", marginTop: "20%" }}>
          <CircularProgress disableShrink />
        </div>
      )}
      {user !== null && allFollowers !== null && allFollows !== null && (
        <>
          <Grid container style={{ marginTop: "5%" }}>
            <Grid item xs={2} />
            <Grid item xs={8}>
              <UserDetailsOnHomePage
                user={user}
                allFollows={allFollows}
                allFollowers={allFollowers}
                relationShip={relationShip}
                setRelationShip={setRelationShip}
                setAllFollowers={setAllFollowers}
              />
            </Grid>
            <Grid item xs={2} />
          </Grid>

          {relationShip && !relationShip.IsBlocked && (
            <>
              {(relationShip.IsFollowing || !user.ProfileSettings.Public) && (
                <UserHighlightsOnHomePage
                  highlightStories={highlightStories}
                  userId={user.IdString}
                />
              )}

              <PostDetailsOnHomePage
                username={username}
                following={relationShip.IsFollowing}
                user={user}
              />
            </>
          )}

          {!relationShip && (
            <>
              <UserHighlightsOnHomePage
                highlightStories={highlightStories}
                userId={user.IdString}
              />
              <PostDetailsOnHomePage
                username={username}
                following={false}
                user={user}
              />
            </>
          )}
        </>
      )}
    </div>
  );
};

export default UserHomePage;
