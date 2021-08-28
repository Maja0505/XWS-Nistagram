// api/index.js
import React, { useEffect,useState } from 'react'
import { useSnackbar  } from 'notistack';
import SnackbarUtils from '../components/SnackbarUtils';
import { useHistory } from "react-router-dom";

var userid = localStorage.getItem("id");
var socket = new WebSocket("ws://localhost:8080/ws/msg/connect-ws/" + userid);

const handleClick = (text) => {
    SnackbarUtils.info(text);
};


let connect = cb => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
    //sendMsg('{"command": 0, "channel": '+ '"' + userid + '"' + '}');
  };

  socket.onmessage = msg => {
    console.log(msg);
    cb(msg);
    var p1 = JSON.parse(msg.data);
    var p2 = JSON.parse(p1.content)
    console.log(p2.content)
    console.log(p2.channel)
    
    if(p2.content === "tagged you in a post." && p2.channel === userid){
      localStorage.setItem('invisibleNotification',false)
      handleClick(p2.user_from + " " + p2.content)
    }
    if((p2.channel === userid && (p2.content === "started following you." || p2.content === "requested to following you."))){
        localStorage.setItem('invisibleNotification',false)
        handleClick(p2.user_from + " " + p2.content)

   }
   if((p2.channel === userid && (p2.content === "liked your photo." || p2.content === "disliked your photo." ))){

     if(window.location.href.split('/')[window.location.href.split('/').length-1] !== p2.post_id){
      localStorage.setItem('invisibleNotification',false)
      handleClick(p2.user_from + " " + p2.content)
     }
   
     
   

   }
    if(p2.content === "commented your post:" && p2.channel === userid){
      localStorage.setItem('invisibleNotification',false)
      handleClick(p2.user_from + " " + p2.content + " " + p2.comment)
    }

    if(p2.content === "sent you a message." && p2.channel === userid){
      localStorage.setItem('invisibleNotification',false)
      handleClick(p2.user_from + " " + p2.content)
    }
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
