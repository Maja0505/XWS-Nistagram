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
          }
    
          
          socket.onmessage = event => {
            var p1 = JSON.parse(event.data);
            var p2 = JSON.parse(p1.content)
            if(p2.channel !== userid){
                localStorage.setItem('invisibleNotification',false)
                handleClick(p2.user_who_follow + " " + p2.content)

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
