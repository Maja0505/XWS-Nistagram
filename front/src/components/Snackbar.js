import React, { useEffect, useState } from "react";
import { useSnackbar } from "notistack";
import SnackbarUtils from "./SnackbarUtils";
import { useHistory } from "react-router-dom";

const Snackbar = () => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const { enqueueSnackbar, closeSnackbar } = useSnackbar();
  const [urls, setUrls] = useState([]);

  const handleClick = (text) => {
    SnackbarUtils.info(text);
  };

  useEffect(() => {
    setUrls();

    SnackbarUtils.setSnackBar(enqueueSnackbar, closeSnackbar);
    let userid = localStorage.getItem("id");
    let username = localStorage.getItem("username");

    if (userid) {
      let socket = new WebSocket(
        "ws://localhost:8080/api/notification/chat/" + userid
      );
      socket.onopen = () => {
        console.log("Successfully Connected");
        socket.send('{"command": 0, "channel": ' + '"' + userid + '"' + "}");
      };

      socket.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
        let socket = new WebSocket(
          "ws://localhost:8080/api/notification/chat/" + userid
        );
      };

      socket.onmessage = (event) => {
        var p1 = JSON.parse(event.data);
        var p2 = JSON.parse(p1.content);
        if (p2.content === "tagged you in a post." && p2.channel === userid) {
          localStorage.setItem("invisibleNotification", false);
          handleClick(p2.user_who_follow + " " + p2.content);
        }
        if (
          p2.channel === userid &&
          (p2.content === "started following you." ||
            p2.content === "requested to following you.")
        ) {
          localStorage.setItem("invisibleNotification", false);
          handleClick(p2.user_who_follow + " " + p2.content);
        }
        if (
          p2.channel === userid &&
          (p2.content === "liked your photo." ||
            p2.content === "disliked your photo.")
        ) {
          var index = urls.length - 1;

          if (
            window.location.href.split("/")[
              window.location.href.split("/").length - 1
            ] !== p2.post_id
          ) {
            localStorage.setItem("invisibleNotification", false);
            handleClick(p2.user_who_follow + " " + p2.content);
          }
        }
        if (p2.content === "commented your post:" && p2.channel === userid) {
          localStorage.setItem("invisibleNotification", false);
          handleClick(p2.user_who_follow + " " + p2.content + " " + p2.comment);
        }

        if (p2.content === "sent you a message." && p2.channel === userid) {
          localStorage.setItem("invisibleNotification", false);
          handleClick(p2.user_who_follow + " " + p2.content);
        }
      };
    }
  }, []);

  return <div></div>;
};

export default Snackbar;
