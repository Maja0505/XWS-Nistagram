import React, { useEffect } from 'react'
import { useSnackbar  } from 'notistack';
import SnackbarUtils from './SnackbarUtils';

const Snackbar = () => {
    const { enqueueSnackbar,closeSnackbar } = useSnackbar();

    const handleClick = (text) => {
        SnackbarUtils.info(text);
      };

    useEffect(() => {
        SnackbarUtils.setSnackBar(enqueueSnackbar,closeSnackbar)
        let userid = localStorage.getItem("id")
        let username = localStorage.getItem("username")
        if(userid){
            let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + userid)
            socket.onopen = () => {
              console.log("Successfully Connected");
              socket.send('{"command": 0, "channel": ' + '"' + userid + '"' + '}')
            };
      
            socket.onclose = event => {
              console.log("Socket Closed Connection: ", event);
              let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + userid)

          }
    
          
          socket.onmessage = event => {
            var p1 = JSON.parse(event.data);
            var p2 = JSON.parse(p1.content)
            console.log("channel " + p2.channel + " loggeduser " + userid)
            if(p2.content === "tagged you in a post." && p2.channel === userid){
              localStorage.setItem('invisibleNotification',false)
              handleClick(p2.user_who_follow + " " + p2.content)
            }
            if((p2.channel !== userid && (p2.content === "started following you." || p2.content === "requested to following you.")) || (p2.channel === userid && (p2.content === "liked your photo." || p2.content === "disliked your photo." ))){
                localStorage.setItem('invisibleNotification',false)
                handleClick(p2.user_who_follow + " " + p2.content)

           }
            if(p2.content === "commented your post:" && p2.channel === userid){
              localStorage.setItem('invisibleNotification',false)
              handleClick(p2.user_who_follow + " " + p2.content + " " + p2.comment)
            }
            };
      };
    

    }, [])

    return (
        <div>
            
        </div>
    )
}

export default Snackbar
