import {
  Grid,
  Paper,
  MenuItem,
  MenuList,
  Button,
  InputBase,
  Divider,
} from "@material-ui/core";
import { AccountBalanceOutlined } from "@material-ui/icons";
import { Link } from "react-router-dom";
import axios from "axios";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import Picker from "emoji-picker-react";
import { v4 as uuidv4 } from "uuid";

import DialogForNewMessage from "./DialogForNewMessage";
import OneTimeMessage from "./OneTimeMessage";

import React, { useEffect, useState } from "react";

const Message = () => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const loggedUserId = localStorage.getItem("id");
  const username = localStorage.getItem("username");

  const [selectedFile, setSelectedFile] = useState();
  const [image, setImage] = useState();
  const [isVideo, setIsVideo] = useState();
  const [puklaSlika, setPuklaSlika] = useState(false);

  const [fromUsersMessage, setFromUsersMessage] = useState([]);
  const [openDialogForNewMessage, setOpenDialogForNewMessage] = useState(false);
  const [openChat, setOpenChat] = useState(false);
  const [messages, setMessages] = useState([]);
  const [newComment, setNewComment] = useState("");
  const [channelName, setChannelName] = useState("");
  const [userFromChatUsername, setUserFromChatUsername] = useState("");
  const [openOneTimeMessage, setOpenOneTimeMessage] = useState(false);
  const [imageForSend, setImageForSend] = useState();
  const [userIDFromMessage, setUserIDFromMessage] = useState();

  const [openPicker, setOpenPicker] = useState(false);

  const handleClickPostComment = () => {
    let socket = new WebSocket(
      "ws://localhost:8080/api/message/chat/" + loggedUserId
    );
    socket.onopen = () => {
      console.log("Successfully Connected");
      socket.send(
        '{"id":true' +
          ',"command": 2, "channel":"' +
          channelName +
          '", "content": "","opened":false,"type":0,"text":"' +
          newComment +
          '","user_from":"' +
          username +
          '","user_to":"' +
          userFromChatUsername +
          '"}'
      );
      setNewComment("");
      socket.send(
        '{"user_who_follow":' +
          '"' +
          username +
          '"' +
          ',"command": 2, "channel": ' +
          '"' +
          userIDFromMessage +
          '"' +
          ', "content": "sent you a message."' +
          ', "media": "' +
          '"' +
          ', "post_id": "' +
          '"}'
      );
    };
  };
  const addToComment = (event, emojiObject) => {
    setNewComment(newComment + emojiObject.emoji);
  };

  const Emojis = (
    <>
      {openPicker && (
        <Grid container>
          <Picker onEmojiClick={addToComment} />
        </Grid>
      )}
    </>
  );

  const handleClickOpenPicker = () => {
    if (openPicker) {
      setOpenPicker(false);
    } else {
      setOpenPicker(true);
    }
  };

  const handleOpenOneTimeMessage = (message) => {
    setOpenOneTimeMessage(true);
    setImageForSend(message);
  };

  useEffect(() => {
    axios
      .get(
        "http://localhost:8080/api/message/user/" + loggedUserId + "/chats",
        authorization
      )
      .then((res) => {
        setFromUsersMessage(res.data);
      })
      .catch((error) => {
        //console.log(error);
      });
    let socket = new WebSocket(
      "ws://localhost:8080/api/message/chat/" + loggedUserId
    );

    socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      let socket = new WebSocket(
        "ws://localhost:8080/api/message/chat/" + loggedUserId
      );
    };
  }, []);

  const handleClickOpenNewMessageDialog = () => {
    setOpenDialogForNewMessage(true);
  };

  const uploadImage = (imageForUpload, index) => {
    var imageId = uuidv4().toString() + "A" + loggedUserId.toString() + ".jpg";
    axios
      .post(
        "/api/media/upload-media-image/" +
          imageId.substring(0, imageId.length - 4) +
          "/" +
          "image" +
          index,
        imageForUpload,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {
        let socket = new WebSocket(
          "ws://localhost:8080/api/message/chat/" + loggedUserId
        );
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send(
            '{"id":true' +
              ',"command": 2, "channel":"' +
              channelName +
              '", "content": "","opened":false,"type":3,"text":"' +
              newComment +
              '","user_from":"' +
              username +
              '","user_to":"' +
              userFromChatUsername +
              '"' +
              ',"user_for_content_id":' +
              '"' +
              username +
              '"' +
              ',"content_id":' +
              '"' +
              imageId +
              '"}'
          );
          socket.send(
            '{"user_who_follow":' +
              '"' +
              username +
              '"' +
              ',"command": 2, "channel": ' +
              '"' +
              userIDFromMessage +
              '"' +
              ', "content": "sent you a message."' +
              ', "media": "' +
              '"' +
              ', "post_id": "' +
              '"}'
          );
        };
        setNewComment("");
      })
      .catch((error) => {
        //alert(error);
        setPuklaSlika(true);
      });
  };

  const uploadVideo = (imageForUpload, index) => {
    var imageId = uuidv4().toString() + "A" + loggedUserId.toString() + ".mp4";

    axios
      .post(
        "/api/media/upload-video/" +
          imageId.substring(0, imageId.length - 4) +
          "/" +
          "image" +
          index,
        imageForUpload,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {
        let socket = new WebSocket(
          "ws://localhost:8080/api/message/chat/" + loggedUserId
        );

        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send(
            '{"id":true' +
              ',"command": 2, "channel":"' +
              channelName +
              '", "content": "","opened":false,"type":3,"text":"' +
              newComment +
              '","user_from":"' +
              username +
              '","user_to":"' +
              userFromChatUsername +
              '"' +
              ',"user_for_content_id":' +
              '"' +
              username +
              '"' +
              ',"content_id":' +
              '"' +
              imageId +
              '"}'
          );
          socket.send(
            '{"user_who_follow":' +
              '"' +
              username +
              '"' +
              ',"command": 2, "channel": ' +
              '"' +
              userIDFromMessage +
              '"' +
              ', "content": "sent you a message."' +
              ', "media": "' +
              '"' +
              ', "post_id": "' +
              '"}'
          );
        };
        setNewComment("");
      })
      .catch((error) => {
        //alert(error);
        setPuklaSlika(true);
      });
  };

  const handleOpenChat = (user) => {
    axios
      .get(
        "/api/message/channels/" + loggedUserId + "/" + user + "/" + "messages",
        authorization
      )
      .then((res) => {
        console.log(res.data);
        setMessages(res.data);
        setChannelName(res.data[0].channel);
        setUserIDFromMessage(user);
        if (res.data[0].user_from !== username) {
          setUserFromChatUsername(res.data[0].user_from);
        } else {
          setUserFromChatUsername(res.data[0].user_to);
        }
        let socket = new WebSocket(
          "ws://localhost:8080/api/message/chat/" + loggedUserId
        );

        socket.onclose = (event) => {
          console.log("Socket Closed Connection: ", event);
          let socket = new WebSocket(
            "ws://localhost:8080/api/message/chat/" + loggedUserId
          );
        };

        setOpenChat(true);
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const HandleUploadMedia = (event) => {
    var formData = new FormData();
    if (event.target.files[0].type === "video/mp4") {
      setIsVideo(false);
    } else {
      setIsVideo(false);
    }

    var file = event.target.files[0];
    formData.append("image0", file);
    const reader = new FileReader();
    var url = reader.readAsDataURL(file);
    reader.onloadend = function (e) {
      setSelectedFile(reader.result);
    }.bind(this);
    setImage(formData);
    console.log(formData);
    setTimeout(500);
    if (isVideo) {
      uploadVideo(image, 0);
    } else {
      uploadImage(image, 0);
    }

    console.log(isVideo);
  };

  return (
    <div>
      <Button onClick={handleClickOpenNewMessageDialog}>New message</Button>

      <Grid container>
        <Grid item xs={3}>
          {fromUsersMessage !== null && (
            <Paper>
              <MenuList id="menu-list-grow">
                {fromUsersMessage.map((user) => (
                  <MenuItem onClick={() => handleOpenChat(user)}>
                    <Grid container>
                      <Grid item xs={3}></Grid>
                      <Grid item xs={7}>
                        <Link>{user}</Link>
                      </Grid>
                    </Grid>
                  </MenuItem>
                ))}
              </MenuList>
            </Paper>
          )}
        </Grid>

        <Grid item xs={9}>
          {openChat && (
            <Paper>
              {messages.map((message) => (
                <>
                  {message.type === 0 && (
                    <>
                      {message.user_from !== username && (
                        <Grid container>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "gray",
                                width: "100%",
                                textAlign: "left",
                              }}
                            >
                              <p style={{ textAlign: "left", color: "white" }}>
                                {message.user_from}: {message.text}
                              </p>
                            </Paper>
                          </Grid>
                          <Grid item xs={9}></Grid>
                        </Grid>
                      )}
                      {message.user_from === username && (
                        <Grid container>
                          <Grid item xs={9}></Grid>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "purple",
                                width: "100%",
                              }}
                            >
                              <p
                                style={{
                                  textAlign: "right",
                                  color: "white",
                                  textAlign: "left",
                                }}
                              >
                                {message.user_from}: {message.text}
                              </p>
                            </Paper>
                          </Grid>
                        </Grid>
                      )}
                    </>
                  )}
                  {message.type === 1 && (
                    <>
                      {message.user_from !== username && (
                        <Grid container>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "gray",
                                width: "100%",
                                textAlign: "left",
                              }}
                            >
                              <p style={{ textAlign: "left", color: "white" }}>
                                {message.user_from}: link:
                                <Link>{message.content_id}</Link>
                                <p>{message.text}</p>
                              </p>
                            </Paper>
                          </Grid>
                          <Grid item xs={9}></Grid>
                        </Grid>
                      )}
                      {message.user_from === username && (
                        <Grid container>
                          <Grid item xs={9}></Grid>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "purple",
                                width: "100%",
                              }}
                            >
                              <p style={{ textAlign: "left", color: "white" }}>
                                {message.user_from}: post={" "}
                                <Link to={"/dialog/" + message.content_id}>
                                  {message.content_id}
                                </Link>
                                <p>{message.text}</p>
                              </p>
                            </Paper>
                          </Grid>
                        </Grid>
                      )}
                    </>
                  )}
                  {message.type === 3 && (
                    <>
                      {message.user_from !== username && (
                        <Grid container>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "gray",
                                width: "100%",
                                textAlign: "left",
                              }}
                            >
                              <p style={{ textAlign: "left", color: "white" }}>
                                {message.user_from}: one-time message:
                                <Button
                                  disabled={message.opened}
                                  onClick={() =>
                                    handleOpenOneTimeMessage(message)
                                  }
                                >
                                  Open
                                </Button>
                                <p>{message.text}</p>
                              </p>
                            </Paper>
                          </Grid>
                          <Grid item xs={9}></Grid>
                        </Grid>
                      )}
                      {message.user_from === username && (
                        <Grid container>
                          <Grid item xs={9}></Grid>
                          <Grid item xs={3}>
                            <Paper
                              style={{
                                backgroundColor: "purple",
                                width: "100%",
                              }}
                            >
                              <p style={{ textAlign: "left", color: "white" }}>
                                {message.user_from}: one-time message:
                                <Button
                                  disabled={message.opened}
                                  onClick={() =>
                                    handleOpenOneTimeMessage(message)
                                  }
                                >
                                  Open
                                </Button>
                                <p>{message.text}</p>
                              </p>
                            </Paper>
                          </Grid>
                        </Grid>
                      )}
                    </>
                  )}
                </>
              ))}

              <Grid container style={{ height: "30%" }}>
                <Grid item xs={3}>
                  <Divider />
                  <SentimentSatisfiedRoundedIcon
                    style={{
                      margin: "auto",
                      marginTop: "5%",
                      cursor: "pointer",
                    }}
                    fontSize="large"
                    onClick={handleClickOpenPicker}
                  ></SentimentSatisfiedRoundedIcon>
                </Grid>
                <Grid item xs={7}>
                  <Divider />
                  <InputBase
                    placeholder="Add a comment..."
                    inputProps={{ "aria-label": "naked" }}
                    value={newComment}
                    style={{ height: "90%", width: "90%" }}
                    onChange={(e) => setNewComment(e.target.value)}
                  />
                </Grid>
                <Grid item xs={2}>
                  <Divider />
                  <Button
                    disabled={newComment === ""}
                    style={{ margin: "auto", marginTop: "10%" }}
                    color="primary"
                    onClick={handleClickPostComment}
                  >
                    Post
                  </Button>
                </Grid>
              </Grid>
              <Button
                variant="contained"
                component="label"
                style={{ margin: "auto" }}
              >
                {selectedFile !== undefined && selectedFile.length === 0
                  ? `Choose media`
                  : `Change media`}
                <input
                  hidden
                  accept="image/*,video/mp4,video/x-m4v,video/*"
                  multiple
                  type="file"
                  name="myFile"
                  onChange={(event) => HandleUploadMedia(event)}
                />
              </Button>
            </Paper>
          )}
          {Emojis}
        </Grid>
      </Grid>
      {openDialogForNewMessage && (
        <DialogForNewMessage
          open={openDialogForNewMessage}
          setOpen={setOpenDialogForNewMessage}
        ></DialogForNewMessage>
      )}
      {openOneTimeMessage && (
        <OneTimeMessage
          open={openOneTimeMessage}
          setOpen={setOpenOneTimeMessage}
          message={imageForSend}
          user={loggedUserId}
        ></OneTimeMessage>
      )}
    </div>
  );
};

export default Message;
