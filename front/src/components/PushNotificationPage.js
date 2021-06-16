import { Grid, Button, TextField, Divider } from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import React, { useState, useEffect } from "react";
import axios from "axios";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";

const PushNotificationPage = ({user}) => {
  const username = localStorage.getItem("username");
  const [accountPrivacy, setAccountPrivacy] = useState(false);
  const [messageRequest, setMessageRequest] = useState(false);
  const [allowTags, setAllowTags] = useState(false);
  const [likesNotification, setLikesNotification] = useState(false)
  const [commentNotification, setCommentNotification] = useState(false)
  const [messageRequestNotification, setMessageRequestNotification] = useState(false)
  const [messageNotification, setMessageNotification] = useState(false)
  const [followRequestNotification, setFollowRequestNotification] = useState(false)
  const [followNotification, setFollowNotification] = useState(false)
  const [load, setLoad] = useState(false)


  useEffect(() => {
    if(user.NotificationSettings !== undefined) {
        setLikesNotification(user.NotificationSettings.LikeNotification);
        setCommentNotification(user.NotificationSettings.CommentNotification);
        setMessageRequestNotification(user.NotificationSettings.MessageRequestNotification)
        setMessageNotification(user.NotificationSettings.MessageNotification)
        setFollowRequestNotification(user.NotificationSettings.FollowRequestNotification)
        setFollowNotification(user.NotificationSettings.FollowNotification)
        setLoad(true)
    }
  }, [user]);


  const HandleOnChangeLikesNotification = (value) => {
 
      if(likesNotification !== value){
        axios.put("/api/user/" + username + "/like-notification/" + value).then((res) => {
            setLikesNotification(value === 'true');
          });
      }
    
  };

  const HandleOnChangeCommentNotification = (value) => {
 
    if(commentNotification !== value){
      axios.put("/api/user/" + username + "/comment-notification/" + value).then((res) => {
          setCommentNotification(value === 'true');
        });
    }
  
};

const HandleOnChangeMessageRequestNotification = (value) => {
 
    if(messageRequestNotification !== value){
      axios.put("/api/user/" + username + "/message-request-notification/" + value).then((res) => {
          setMessageRequestNotification(value === 'true');
        });
    }
  
};

const HandleOnChangeMessageNotification = (value) => {
 
    if(messageNotification !== value){
      axios.put("/api/user/" + username + "/message-notification/" + value).then((res) => {
          setMessageNotification(value === 'true');
        });
    }
  
};

const HandleOnChangeFollowRequestNotification = (value) => {
 
    if(followRequestNotification !== value){
      axios.put("/api/user/" + username + "/follow-request-notification/" + value).then((res) => {
          setFollowRequestNotification(value === 'true');
        });
    }
  
};

const HandleOnChangeFollowNotification = (value) => {
 
    if(followNotification !== value){
      axios.put("/api/user/" + username + "/follow-notification/" + value).then((res) => {
          setFollowNotification(value === 'true');
        });
    }
  
};
  return (
    <Grid container item xs={9} style={{ height: 600 }}>
     <Grid item xs={1}></Grid>
     {load && <Grid container item xs={10}>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>
              Likes notification
            </p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="likes" value={likesNotification} name="likes" onClick={(e) => HandleOnChangeLikesNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
            johnappleseed liked your photo.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Comments notification</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="comment" name="comment" value={commentNotification} onClick={(e)=> HandleOnChangeCommentNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
            johnappleseed commented: "Nice shot!"
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Messages request notification</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="message-request" name="message-request" value={messageRequestNotification} onClick={(e)=> HandleOnChangeMessageRequestNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
            johnappleseed wants to send you a message.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Messages notification</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="message" name="message" value={messageNotification} onClick={(e)=> HandleOnChangeMessageNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
            johnappleseed sent you a message.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>
              Follow request notification
            </p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="follow-request" name="follow-request" value={followRequestNotification} onClick={(e)=> HandleOnChangeFollowRequestNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
                johnappleseed requested to follow you.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "20%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>
             Follow notification
            </p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <RadioGroup aria-label="follow" name="follow" value={followNotification} onClick={(e)=> HandleOnChangeFollowNotification(e.target.value)}>
              <FormControlLabel value={false} control={<Radio />} label="Off" />
              <FormControlLabel
                value={true}
                control={<Radio />}
                label="From Everyone"
              />
            </RadioGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
            johnappleseed started following you.
            </p>
          </Grid>
          <Divider />
        </Grid>
      </Grid>}
      <Grid item xs={1}></Grid>
    </Grid>
  );
};

export default PushNotificationPage;
