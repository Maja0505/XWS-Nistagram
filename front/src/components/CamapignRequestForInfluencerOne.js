import { Grid, Button } from "@material-ui/core";
import { useEffect, useState } from "react";

import axios from "axios";
const CamapignRequestForInfluencerOne = ({ campaignId }) => {
  const [campaign, setCampaign] = useState();
  const loggedUserId = localStorage.getItem("id");

  useEffect(() => {
    axios.get("/api/agent/get-campaign/" + campaignId).then((res) => {
      setCampaign(res.data);
    });
  }, []);

  const acceptRequest = () => {
    axios
      .post("/api/agent/add-influencer", {
        ID: campaignId,
        UserID: campaign.UserID,
        InfluencerID: loggedUserId,
      })
      .then((res) => {
        alert("Success accept camapign request !");
      }).catch((error) => {
        //console.log(error);
      });;
  };

  return (
    <div>
      {campaign !== null && campaign !== undefined && (
        <Grid container>
          <Grid item xs={8} style={{ margin: "auto", textAlign: "left" }}>
            <label style={{ fontSize: "10px" }}>
              Camapign ID : {campaignId}
            </label>
            <br></br>
            <label style={{ fontSize: "10px" }}>
              Agent ID : {campaign.UserID}
            </label>
            <br></br>
          </Grid>
          <Grid item xs={4} style={{ margin: "auto", textAlign: "left" }}>
            <Button color="inherit" onClick={acceptRequest}>
              Accept request
            </Button>
          </Grid>
        </Grid>
      )}
    </div>
  );
};

export default CamapignRequestForInfluencerOne;
