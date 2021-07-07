import { Grid, Paper, Tabs, Tab } from "@material-ui/core";

import { useState } from "react";

import AdminVerificationRequestCard from "./AdminVerificationRequestCard";
import ReportedContents from "./ReportedContents.js";

const AdminHomePage = () => {
  const [tabValue, setTabValue] = useState(0);

  const handleChangeTab = (event, newValue) => {
    setTabValue(newValue);
  };

  return (
    <div>
      <Grid container style={{ marginTop: "5%" }}>
        <Grid item xs={2}></Grid>
        <Grid container item xs={8}>
          <Grid item xs={12}>
            <Paper>
              <Tabs
                value={tabValue}
                onChange={handleChangeTab}
                indicatorColor="primary"
                textColor="inherit"
              >
                <Tab label="Verification requests" style={{ margin: "auto" }} />
                <Tab label="Reported posts" style={{ margin: "auto" }} />
              </Tabs>
            </Paper>
          </Grid>

          <Grid item xs={10} style={{ margin: "auto" }}>
            {tabValue === 0 && <AdminVerificationRequestCard />}
            {tabValue === 1 && <ReportedContents />}
          </Grid>
        </Grid>
        <Grid item xs={2}></Grid>
      </Grid>
    </div>
  );
};

export default AdminHomePage;
