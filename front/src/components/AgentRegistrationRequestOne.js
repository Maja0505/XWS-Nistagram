import { Paper, Grid, Button } from "@material-ui/core";

import { useState } from "react";

import axios from "axios";

const AgentRegistrationRequestOne = ({
  request,
  setAgentRequests,
  agentRequests,
}) => {
  const [webSiteOK, setWebSiteOK] = useState(false);
  const [clickedOnCheck, setClickedOnCheck] = useState(false);

  function isValidURL(string) {
    var res = string.match(
      /(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g
    );
    return res !== null;
  }

  const checkWebSite = () => {
    setClickedOnCheck(true);
    setWebSiteOK(isValidURL(request.WebSite));
  };

  const approveRequest = () => {
    axios
      .put(
        "/api/user/agent-registration-request/update-to-approved/" +
          request.Username,
        {}
      )
      .then((res) => {
        deleteFromArray();
      }).catch((error) => {
        //console.log(error);
      });;
  };

  const rejectRequest = () => {
    axios
      .put(
        "/api/user/agent-registration-request/delete/" + request.Username,
        {}
      )
      .then((res) => {
        deleteFromArray();
      }).catch((error) => {
        //console.log(error);
      });;
  };

  const deleteFromArray = () => {
    var array = [...agentRequests];
    var index = array.indexOf(request);
    console.log(array);
    if (index !== -1) {
      array.splice(index, 1);
    }
    setAgentRequests(array);
  };

  return (
    <div>
      {request !== null && request !== undefined && (
        <Paper>
          <Grid container style={{ marginBottom: "1%" }}></Grid>
          <Grid container>
            <Grid item xs={1}></Grid>
            <Grid
              item
              xs={7}
              style={{ textAlign: "left", overflowWrap: "anywhere" }}
            >
              <h6 style={{ textAlign: "left" }}>User ID : {request.ID}</h6>
              <label>
                Username : <b>{request.Username}</b>
              </label>
              <br></br>
              <label>
                First name : <b>{request.FirstName}</b>
              </label>
              <br></br>
              <label>
                Last name : <b>{request.LastName}</b>
              </label>
              <br></br>
              <label>
                Email : <b>{request.Email}</b>
              </label>
              <br></br>
              <label>
                Web site : <a href={request.WebSite}>{request.WebSite}</a>
              </label>
              <br></br>
              <label>
                Phone number : <b>{request.PhoneNumber}</b>
              </label>
              <br></br>
              <br></br>
            </Grid>
            <Grid item xs={4} style={{ margin: "auto" }}>
              <Button
                variant="text"
                color="primary"
                style={{ margin: "auto" }}
                onClick={approveRequest}
              >
                Approve
              </Button>
              <Button
                variant="text"
                color="secondary"
                style={{ margin: "auto" }}
                onClick={rejectRequest}
              >
                Reject
              </Button>
              {clickedOnCheck === false && (
                <Button
                  variant="text"
                  color="inherit"
                  style={{ margin: "auto" }}
                  onClick={checkWebSite}
                >
                  Check web site
                </Button>
              )}
              {clickedOnCheck === true && webSiteOK === false && (
                <>
                  <br></br>
                  <label>Web site isn't ok</label>
                </>
              )}
              {clickedOnCheck === true && webSiteOK === true && (
                <>
                  <br></br>
                  <label>Web site is ok</label>
                </>
              )}
            </Grid>
          </Grid>
          <Grid container style={{ marginTop: "1%" }}></Grid>
        </Paper>
      )}
    </div>
  );
};

export default AgentRegistrationRequestOne;
