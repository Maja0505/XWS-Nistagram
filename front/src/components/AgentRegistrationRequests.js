import { Paper, Grid, Button } from "@material-ui/core";

import { useEffect, useState } from "react";

import axios from "axios";

import AgentRegistrationRequestOne from "./AgentRegistrationRequestOne";

const AgentRegistrationRequests = () => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };
  const [agentRequests, setAgentRequests] = useState([]);

  useEffect(() => {
    axios.get("/api/user/agent-registration-request/get-all",authorization).then((res) => {
      if (res.data !== null) {
        console.log(res.data);
        setAgentRequests(res.data);
      }
    }).catch((error) => {
      //console.log(error);
    });
  }, []);

  return (
    <div>
      <Grid container style={{ marginTop: "3%" }}></Grid>
      {agentRequests.map((r, index) => (
        <Grid container key={index} style={{ marginTop: "2%" }}>
          <Grid item xs={2} />
          <Grid item xs={8}>
            <AgentRegistrationRequestOne
              request={r}
              agentRequests={agentRequests}
              setAgentRequests={setAgentRequests}
            />
          </Grid>
          <Grid item xs={2} />
        </Grid>
      ))}
      {agentRequests !== null && agentRequests !== undefined && (
        <>
          {agentRequests.length === 0 && (
            <Paper>
              <Grid container style={{ marginBottom: "2%" }}></Grid>
              <Grid container>
                <p style={{ margin: "auto" }}>NO AGENT REGISTRATION REQUESTS</p>
              </Grid>
              <Grid container style={{ marginTop: "2%" }}></Grid>
            </Paper>
          )}
        </>
      )}
    </div>
  );
};

export default AgentRegistrationRequests;
