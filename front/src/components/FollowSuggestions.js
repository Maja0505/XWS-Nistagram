import { Grid, Paper } from "@material-ui/core";
import FollowSuggestionOne from "./FollowSuggestionOne";
import { useState, useEffect } from "react";

import axios from "axios";

const FollowSuggestions = () => {
  const loggedUserID = localStorage.getItem("id");
  const [followSuggestions, setFollowSuggestions] = useState([]);

  useEffect(() => {
    axios
      .get("/api/user-follow/followSuggestions/" + loggedUserID)
      .then((res) => {
        setFollowSuggestions(res.data);
      }).catch((error) => {
        //console.log(error);
      });
  }, []);

  return (
    <div style={{ marginTop: "5%", marginBottom: "3%" }}>
      {followSuggestions.length !== 0 && (
        <>
          {followSuggestions.map((s, index) => (
            <Grid container style={{ marginTop: "1%" }} key={index}>
              <Grid item xs={4} />
              <Grid item xs={4}>
                <FollowSuggestionOne suggestion={s} />
              </Grid>
              <Grid item xs={4} />
            </Grid>
          ))}
        </>
      )}
    </div>
  );
};

export default FollowSuggestions;
