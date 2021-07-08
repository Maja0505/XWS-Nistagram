import { Grid, Paper, Tabs, Tab } from "@material-ui/core";

import { useState } from "react";

import { Redirect } from "react-router-dom";

import RegistrationAgent from "./RegistrationAgent";
import RegistrationUser from "./RegistrationUser";

const RegistartionPage = () => {
  const [tabValue, setTabValue] = useState(0);

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  const [redirection, setRedirection] = useState(false);

  return (
    <div>
      {redirection === true && <Redirect to="/login" />}

      <Grid container style={{ marginTop: "5%" }}>
        <Grid item xs={3}></Grid>
        <Grid item xs={6}>
          <Paper>
            <Tabs
              value={tabValue}
              onChange={handleChangeTab}
              indicatorColor="primary"
              textColor="inherit"
            >
              <Tab label="User Registration" style={{ margin: "auto" }} />
              <Tab label="Agent registration" style={{ margin: "auto" }} />
            </Tabs>
          </Paper>
        </Grid>
        <Grid item xs={3}></Grid>
      </Grid>

      {tabValue === 0 && <RegistrationUser setRedirection={setRedirection} />}
      {tabValue === 1 && <RegistrationAgent setRedirection={setRedirection} />}
      <Grid container style={{ marginBottom: "2%" }}></Grid>
    </div>
  );
};

export default RegistartionPage;
