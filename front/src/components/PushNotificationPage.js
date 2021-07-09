import { Grid, Button, TextField, Divider } from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import React, { useState, useEffect } from "react";
import axios from "axios";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";

const PushNotificationPage = ({pushNotification, setPushNotification,load}) => {
  const username = localStorage.getItem("username");


  const HandleOnChangeLikesNotification = (value) => {
 
      if(pushNotification.LikesNotification !== value){
        axios.put("/api/user/" + username + "/like-notification/" + value).then((res) => {
            setPushNotification({...pushNotification, LikeNotification: value === 'true'})
          });
      }
    
  };

  const HandleOnChangeCommentNotification = (value) => {
 
    if(pushNotification.CommentNotification !== value){
      axios.put("/api/user/" + username + "/comment-notification/" + value).then((res) => {
        setPushNotification({...pushNotification, CommentNotification: value === 'true'})

        });
    }
  
};

const HandleOnChangeMessageRequestNotification = (value) => {
 
    if(pushNotification.MessageRequestNotification !== value){
      axios.put("/api/user/" + username + "/message-request-notification/" + value).then((res) => {
        setPushNotification({...pushNotification, MessageRequestNotification: value === 'true'})

        });
    }
  
};

const HandleOnChangeMessageNotification = (value) => {
 
    if(pushNotification.MessageNotification !== value){
      axios.put("/api/user/" + username + "/message-notification/" + value).then((res) => {
        setPushNotification({...pushNotification, MessageNotification: value === 'true'})
        });
    }
  
};

const HandleOnChangeFollowRequestNotification = (value) => {
 
    if(pushNotification.FollowRequestNotification !== value){
      axios.put("/api/user/" + username + "/follow-request-notification/" + value).then((res) => {
        setPushNotification({...pushNotification, FollowRequestNotification: value === 'true'})
        });
    }
  
};

const HandleOnChangeFollowNotification = (value) => {
 
    if(pushNotification.FollowNotification !== value){
      axios.put("/api/user/" + username + "/follow-notification/" + value).then((res) => {
        setPushNotification({...pushNotification, FollowNotification: value === 'true'})
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
            <RadioGroup aria-label="likes" value={pushNotification.LikeNotification === true} name="likes" onClick={(e) => HandleOnChangeLikesNotification(e.target.value)}>
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
            <RadioGroup aria-label="comment" name="comment" value={pushNotification.CommentNotification === true} onClick={(e)=> HandleOnChangeCommentNotification(e.target.value)}>
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
            <RadioGroup aria-label="message-request" name="message-request" value={pushNotification.MessageRequestNotification === true} onClick={(e)=> HandleOnChangeMessageRequestNotification(e.target.value)}>
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
            <RadioGroup aria-label="message" name="message" value={pushNotification.MessageNotification === true} onClick={(e)=> HandleOnChangeMessageNotification(e.target.value)}>
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
            <RadioGroup aria-label="follow-request" name="follow-request" value={pushNotification.FollowRequestNotification === true} onClick={(e)=> HandleOnChangeFollowRequestNotification(e.target.value)}>
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
            <RadioGroup aria-label="follow" name="follow" value={pushNotification.FollowNotification === true} onClick={(e)=> HandleOnChangeFollowNotification(e.target.value)}>
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
