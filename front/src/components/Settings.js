import React, { useEffect, useState } from "react";
import { Tabs, Tab, Grid } from "@material-ui/core";
import Box from "@material-ui/core/Box";
import ChangePasswordPage from "./ChangePasswordPage.js";
import { useHistory } from "react-router-dom";
import ProfilePage from "./ProfilePage";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Snackbar from '@material-ui/core/Snackbar';
import { Alert } from "@material-ui/lab";
import ProfilePrivacy from "./ProfilePrivacy.js";
import PushNotificationPage from "./PushNotificationPage.js";
import VerificationRequest from "./VerificationRequest";
import axios from "axios";



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
  },{
    key: 4,
    id: 4,
    label: "Request verification",
  },

];

const Settings = () => {

  const [user, setUser] = useState({});
  const username = localStorage.getItem("username");

  const [tabs] = useState(tabList);
  const [value, setValue] = useState(0);


  const [open, setOpen] = useState(false)
  const [message,setMessage] = useState('')

  const [selectedValue, setSelectedValue] = useState("male");
  const [load, setLoad] = useState(false);
  const [userCopy, setUserCopy] = useState({});

  const [accountPrivacy,setAccountPrivacy] = useState(false)
  const [messageRequest,setMessageRequest] = useState(false)
  const [allowTags,setAllowTags] = useState(false)
  const [profileSettings,setProfileSettings] = useState({})
  const [pushNotification, setPushNotification] = useState({})


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
    return <ProfilePage user={user} setUser={setUser} selectedValue={selectedValue} setSelectedValue={setSelectedValue} userCopy={userCopy} setUserCopy={setUserCopy} load={load}></ProfilePage>
  }else if(value === 1){
    return <ChangePasswordPage setOpen={setOpen} setMessage={setMessage}></ChangePasswordPage>
  }else if(value === 2){
    return <ProfilePrivacy profileSettings={profileSettings} setProfileSettings={setProfileSettings} load={load}></ProfilePrivacy>
  }else if(value === 3){
    return <PushNotificationPage pushNotification={pushNotification} setPushNotification={setPushNotification} load={load}></PushNotificationPage>
  }else if(value === 4){
    return <VerificationRequest user={user} setOpen={setOpen} setMessage={setMessage}></VerificationRequest>
  }
}

const handleUrlForTab = (value) => {
  var route =''
  if (value === 0) {
    route = '/accounts/edit/'
  }else if(value === 1){
    route = '/accounts/password/change/'
  }else if(value === 2){
    route = '/accounts/privacy/'
  }else if(value === 3){
    route = '/accounts/notification/'
  }else if(value === 4){
    route = '/accounts/verification/'
  }
  history.push(route)
  
}

useEffect(() => {
 
  var urls = window.location.href.split('accounts')
  if(urls[1] === '/edit/'){
    setValue(0)
  }else if(urls[1] === '/password/change/'){
    setValue(1)
  }else if(urls[1] === '/privacy/'){
    setValue(2)
  }else if(urls[1] === '/notification/'){
    setValue(3)
  }else if(urls[1] === '/verification/'){
    setValue(4)
  }
  axios.get("/api/user/" + username).then((res) => {
    setUser(res.data);
    res.data.Gender === 0 ? setSelectedValue("male") : setSelectedValue("female");
    setLoad(true);
    setUserCopy(res.data);

    setProfileSettings(res.data.ProfileSettings)
    setPushNotification(res.data.NotificationSettings)

  }).catch((error) => {
    //console.log(error);
  });

}, []);

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
               {TabChanged}
            </Box>
          </Grid>

          <Grid item xs={2} />
        </Grid>
      </Grid>
    <Snackbar
        open={open}
        autoHideDuration={2000} 
        onClose={handleClose}
        TransitionComponent='TransitionRight'
        message={message}
        
      >
        {message === 'Successful changed password' || message === 'Successfully sent verification request' && <Alert onClose={handleClose} severity="success">
          {message}
        </Alert>}
     </Snackbar>
    </div>
  );
};

export default Settings;
