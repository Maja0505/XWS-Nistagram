import React, { useEffect, useState } from "react";
import { Tabs, Tab, Grid } from "@material-ui/core";
import Box from "@material-ui/core/Box";
import ChangePasswordPage from "./ChangePasswordPage.js";
import { useHistory } from "react-router-dom";
import ProfilePage from "./ProfilePage";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Snackbar from '@material-ui/core/Snackbar';
import { Alert } from "@material-ui/lab";


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


  const [open, setOpen] = useState(false)
  const [message,setMessage] = useState('')

  const handleClick = () => {
    setOpen(true);
  };

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };
  const handleTabChange = (event, value) => {
    setValue(value);
    handleUrlForTab(value)
  };
  let history = useHistory();

const TabChanged = () => {
  if (value === 0) {
    return  <ProfilePage></ProfilePage>
  }else if(value === 1){
    return <ChangePasswordPage setOpen={setOpen} setMessage={setMessage}></ChangePasswordPage>
  }
}

const handleUrlForTab = (value) => {
  var route =''
  if (value === 0) {
    route = '/accounts/edit/'
  }else if(value === 1){
    route = '/accounts/password/change/'
  }
  history.push(route)
}

useEffect(() => {
 
  var urls = window.location.href.split('accounts')
  if(urls[1] === '/edit/'){
    setValue(0)
  }else if(urls[1] === '/password/change/'){
    setValue(1)
  }

}, []);

  return (
    <div>
    <Router>

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
               {TabChanged}
            </Box>
          </Grid>

          <Grid item xs={2} />
        </Grid>
      </Grid>
    </Router>
    <Snackbar
        open={open}
        autoHideDuration={2000} 
        onClose={handleClose}
        TransitionComponent='TransitionRight'
        message={message}
        
      >
        {message === 'Successful changed password' && <Alert onClose={handleClose} severity="success">
          {message}
        </Alert>}
     </Snackbar>
    </div>
  );
};

export default Settings;
