import React, { useState } from "react";
import { Tabs, Tab, Grid } from "@material-ui/core";
import Box from "@material-ui/core/Box";
import ProfilePage from './ProfilePage.js'

const tabList = [
  {
    key: 0,
    id: 0,
    label: "Edit Profile",
  },
  {
    key: 1,
    id: 1,
    label: "Change Password",
  },
  {
    key: 2,
    id: 2,
    label: "Privacy and Security",
  },
  {
    key: 3,
    id: 3,
    label: "Push Notification",
  },
  {
    key: 4,
    id: 4,
    label: "Edit Profile",
  },
];

const Settings = () => {
  const [tabs] = useState(tabList);
  const [value, setValue] = useState(0);
  const handleTabChange = (event, value) => {
    setValue(value);
  };

  return (
    <div>
      <Grid container>
        <Grid container style={{ marginTop: "2%" }}>
          <Grid item xs={2} />
          <Grid container item xs={8}>
            <Grid item xs={3}>
              <Box border={1}>
                <Tabs
                  style={{ height: 700 }}
                  value={value}
                  onChange={handleTabChange}
                  indicatorColor="primary"
                  textColor="primary"
                  orientation="vertical"
                  variant="scrollable"
                  scrollButtons="auto"
                >
                  {tabs.map((tab) => (
                    <Tab
                      key={tab.key.toString()}
                      value={tab.id}
                      label={tab.label}
                    />
                  ))}
                </Tabs>
              </Box>
            </Grid>
            <Box border={1} style={{ width: 600,height: 700 }}>
                <ProfilePage></ProfilePage>
            </Box>
          </Grid>

          <Grid item xs={2} />
        </Grid>
      </Grid>
    </div>
  );
};

export default Settings;
