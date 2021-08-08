import {
  Dialog,
  DialogTitle,
  DialogContent,
  Grid,
  Button,
} from "@material-ui/core";
import { useEffect, useState } from "react";
import CamapignRequestForInfluencerOne from "./CamapignRequestForInfluencerOne";

import axios from "axios";

const ViewCampaignRequestsForInfluencerDialog = ({ label, open, setOpen }) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const [campaignRequests, setCampaignRequests] = useState([]);
  const loggedUserId = localStorage.getItem("id");

  const handleClose = () => {
    setOpen(false);
  };

  useEffect(() => {
    axios
      .get("/api/agent/get-campaign-requests/" + loggedUserId, authorization)
      .then((res) => {
        setCampaignRequests(res.data);
      })
      .catch((error) => {
        //console.log(error);
      });
  }, []);

  return (
    <div>
      <Dialog
        onClose={handleClose}
        aria-labelledby="customized-dialog-title"
        open={open}
      >
        <DialogTitle
          id="customized-dialog-title"
          onClose={handleClose}
          style={{ textAlign: "center", width: 400 }}
        >
          <h3>{label}</h3>
        </DialogTitle>
        <DialogContent style={{ width: 400, height: 400 }}>
          {campaignRequests !== undefined &&
            campaignRequests !== null &&
            campaignRequests.map((c, index) => (
              <Grid container key={index}>
                <CamapignRequestForInfluencerOne campaignId={c.CampaignID} />
              </Grid>
            ))}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default ViewCampaignRequestsForInfluencerDialog;
