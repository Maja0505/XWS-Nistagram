import { Grid,Paper,MenuItem,MenuList,Button ,InputBase,Divider
} from '@material-ui/core'
import { AccountBalanceOutlined } from '@material-ui/icons'
import { Link } from "react-router-dom";
import axios from "axios";
import SentimentSatisfiedRoundedIcon from "@material-ui/icons/SentimentSatisfiedRounded";
import Picker from "emoji-picker-react";

import DialogForNewMessage from './DialogForNewMessage';


import React, { useEffect,useState } from 'react'

const Message = () => {

    const loggedUserId = localStorage.getItem("id");
    const username = localStorage.getItem("username");

    const [fromUsersMessage, setFromUsersMessage] = useState([])
    const [openDialogForNewMessage,setOpenDialogForNewMessage] = useState(false)
    const [openChat,setOpenChat] = useState(false)
    const [messages,setMessages] = useState(false)
    const [newComment, setNewComment] = useState("");
    const [channelName,setChannelName] = useState("")
    const [userFromChatUsername,setUserFromChatUsername] = useState("")

    const [openPicker, setOpenPicker] = useState(false);

    const handleClickPostComment = () => {
        let socket = new WebSocket("ws://localhost:8080/api/message/chat/" + loggedUserId)
        socket.onopen = () => {
        console.log("Successfully Connected");                
        socket.send('{"id":true' + ',"command": 2, "channel":"' + channelName + '", "content": "","opened":false,"type":0,"text":"' +  newComment + '","user_from":"' + username + '","user_to":"' + userFromChatUsername + '"}')
        setNewComment("")
    }

    }
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
    


    
    useEffect(() => {
        axios.get("http://localhost:8080/api/message/user/" + loggedUserId + "/chats")
            .then((res) => {
                setFromUsersMessage(res.data)
            })
       
        
    }, [])

    const handleClickOpenNewMessageDialog = () => {
        setOpenDialogForNewMessage(true)
    }

    const handleOpenChat = (user) => {
        axios.get("/api/message/channels/" + loggedUserId +  "/" + user + "/" + "messages")
            .then((res) => {
                console.log(res.data)
                setMessages(res.data)
                setChannelName(res.data[0].channel)
                if(res.data[0].user_from !== username){
                    setUserFromChatUsername(res.data[0].user_from)
                }else{
                    setUserFromChatUsername(res.data[0].user_to)
                }
                let socket = new WebSocket("ws://localhost:8080/api/message/chat/" + loggedUserId)
    
          
                socket.onclose = event => {
                  console.log("Socket Closed Connection: ", event);
                  let socket = new WebSocket("ws://localhost:8080/api/message/chat/" + loggedUserId)
    
                }
                    socket.onmessage = event => {
                    if (res.data[0].user_from === username){
                    var p1 = JSON.parse(event.data);
                    var p2 = JSON.parse(p1.content)
                    var array = res.data;
                    array.push(p2);
                    setMessages(array);
                    }else{
                        var p1 = JSON.parse(event.data);
                        var p2 = JSON.parse(p1.content)
                        var array = res.data;
                        array.push(p2);
                        setMessages(array);
                    }
                }

                setOpenChat(true)
            })
    }

    

    return (
        <div>
           <Button onClick={handleClickOpenNewMessageDialog}>New message</Button>

            <Grid container>
                <Grid item xs={3}>
            {fromUsersMessage !== null &&
                <Paper>
              <MenuList
                id="menu-list-grow"
              >
                  {fromUsersMessage.map((user) => (
                        <MenuItem onClick={() => handleOpenChat(user)}>
                        <Grid container>
                            <Grid item xs={3}>
                           
                            </Grid>
                            <Grid item xs={7}>
                            <Link>
                                    {user}
                            </Link>
                            </Grid>
                        </Grid>
                        </MenuItem>
                  ))}

              </MenuList>
          </Paper>
            }

            </Grid>
           

                <Grid item xs={9}>
                {openChat && 
                <Paper>
                    {messages.map((message) => (
                        <>
                        {message.type === 0 &&
                            <>
                            {message.user_from !== username &&
                            <Grid container>
                                <Grid item xs={3}>
                                <Paper style={{backgroundColor:"gray", width:"100%",textAlign:"left"}}>
                                     <p style={{textAlign:"left",color:"white"}}>{message.user_from}: {message.text}</p>
                                </Paper>
                                </Grid>
                                <Grid item xs={9}></Grid>
    
                            </Grid>
                              
                            }
                            {message.user_from === username &&
                              <Grid container>
                              <Grid item xs={9}></Grid>
                              <Grid item xs={3}>
                              <Paper style={{backgroundColor:"purple", width:"100%"}}>
    
                                     <p style={{textAlign:"right",color:"white",textAlign:"left"}}>{message.user_from}: {message.text}</p>
                                </Paper>
                              </Grid>
    
                          </Grid>
                               
                            }                 
                         </>
                        }
                     {message.type === 1 &&
                            <>
                            {message.user_from !== username &&
                            <Grid container>
                                <Grid item xs={3}>
                                <Paper style={{backgroundColor:"gray", width:"100%",textAlign:"left"}}>
                                     <p style={{textAlign:"left",color:"white"}}>{message.user_from}: link:<Link>{message.content_id}</Link><p>{message.text}</p></p>
                                </Paper>
                                </Grid>
                                <Grid item xs={9}></Grid>
    
                            </Grid>
                              
                            }
                            {message.user_from === username &&
                              <Grid container>
                              <Grid item xs={9}></Grid>
                              <Grid item xs={3}>
                              <Paper style={{backgroundColor:"purple", width:"100%"}}>
    
                              <p style={{textAlign:"left",color:"white"}}>{message.user_from}: post= <Link to={"/dialog/" + message.content_id}>{message.content_id}</Link><p>{message.text}</p></p>
                                </Paper>
                              </Grid>
    
                          </Grid>
                               
                            }                 
                         </>
                        }

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
                        style={{height:"90%",width:"90%"}}
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

                </Paper>
                }
                {Emojis}

                </Grid>

              </Grid>
            {openDialogForNewMessage && 
            <DialogForNewMessage open={openDialogForNewMessage} setOpen={setOpenDialogForNewMessage}></DialogForNewMessage>
            }
        </div>
    )
}

export default Message
