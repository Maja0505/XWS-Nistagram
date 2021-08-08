import {
  Dialog,
  DialogTitle,
  DialogContent,
  Grid,
  Button,
} from "@material-ui/core";
import { useEffect, useState } from "react";

import axios from "axios";

const AddCampaignToInfluencerDialog = ({
  label,
  influencer,
  agent,
  open,
  setOpen,
}) => {
  const [agentCamapigns, setAgentCamapigns] = useState([]);

  const now = new Date();

  const handleClose = () => {
    setOpen(false);
  };

  useEffect(() => {
    axios.get("/api/agent/get-campaigns-for-user/" + agent).then((res) => {
      console.log(res.data);
      setAgentCamapigns(res.data);
    }).catch((error) => {
      //console.log(error);
    });
  }, []);

  const createCamapignRequest = (campaignId) => {
    axios
      .post("/api/agent/create-campaign-request", {
        CampaignID: campaignId,
        UserID: influencer,
      })
      .then((res) => {
        alert("Success create camapign request !");
      }).catch((error) => {
        //console.log(error);
      });
  };

  const comaringTimes = (endDate) => {
    console.log(now);
    endDate.setHours(endDate.getHours() - 2);
    console.log(endDate);
  };

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
          {agentCamapigns !== undefined &&
            agentCamapigns !== null &&
            agentCamapigns.map((c, index) => (
              <Grid container key={index}>
                <>
                  <Grid
                    item
                    xs={8}
                    style={{ margin: "auto", textAlign: "left" }}
                  >
                    {c.ID}
                  </Grid>
                  <Grid
                    item
                    xs={4}
                    style={{ margin: "auto", textAlign: "left" }}
                  >
                    <Button
                      color="inherit"
                      onClick={() => createCamapignRequest(c.ID)}
                    >
                      Send request
                    </Button>
                  </Grid>
                </>
              </Grid>
            ))}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default AddCampaignToInfluencerDialog;
