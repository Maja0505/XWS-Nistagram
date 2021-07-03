import "./ContentDetails.css";
import Button from "../components/Button";
import avatar from "../images/nistagramAvatar.jpg";
import myPic from "../images/background.jpg";

import React, { useState, useEffect,useRef } from "react";
import Story from "../components/Story";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";

import {
	AppBar,
	Toolbar,
	Grid,
  Paper,

  Grow,
  Popper,
  MenuItem,
  MenuList,
  } from "@material-ui/core";
  import AddStory from "../components/AddStoryDialog"
import axios from "axios";
import { User as UserModel } from "../models/User";





export default function ContentDetails() {
  const [username,setUsername] =  useState("")
  const [users,setUsers] = useState([])
  const loggedUserId = localStorage.getItem("id");

  useEffect(() => {
    axios.get("/api/post/story/all-follows-with-stories/" + loggedUserId)
    .then((res) => {
      if(res.data){
        setUsers(res.data)
      }
    })
   axios.get("/api/post/story/all-for-close-friends/" + loggedUserId)
    .then((res) => {
      console.log(res.data)
      setStories(res.data)
    })
  }, [])
  const [open, setOpen] = useState(false);
  const [openDialog,setOpenDialog] = useState(false)

  const anchorRef = useRef(null);
  const [myStories,setMyStories] = useState([])
  const [stories,setStories] = useState([])

  const storyClicked = (story:any) => {
    console.log(story)
  }

  const [storiesOpen, setStoriesOpen] = useState(false);
  



  const openStories = (username:any) => {
   
    setUsername(username)
    
    if(username === loggedUserId){
      setOpen((prevOpen) => !prevOpen);
 
    }else{
      axios.get("/api/user-follow/checkClosed/"+loggedUserId+"/" + username)
        .then((res) => {
          if(res.data){
            axios.get("/story/all-for-close-friends/" + username)
              .then((res) => {
                setStories(res.data)
                setStoriesOpen(true)
              })
          }else{
            axios.get("/api/post/story/all-not-expired/" + username)
            .then((res) => {
              setStories(res.data)
              setStoriesOpen(true)

            })
          }
        })
    }
  };

  function closeStory() {
    setStoriesOpen(false);
  }

  const handleClose = (event:any) => {
    setOpen(false);
  };

  function handleListKeyDown(event:any) {
    if (event.key === "Tab") {
      event.preventDefault();
      setOpen(false);
    }
  }

  const handleClickAddStory = () => {
    setOpenDialog(true)
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClickOpenMyStories = () => {
    axios.get("/api/post/story/all-for-close-friends/" + loggedUserId)
    .then((res) => {
      console.log(res.data)
      setStories(res.data)
      setStoriesOpen(true);
      setOpen((prevOpen) => !prevOpen);
    })
   
    }

  const dropDowMenuForPost = (
    <Popper
      open={open}
      anchorEl={anchorRef.current}
      role={undefined}
      transition
      disablePortal
      style={{ width: "15%", zIndex:1 }}
    >
      {({ TransitionProps, placement }) => (
        <Grow
          {...TransitionProps}
          style={{
            transformOrigin:
              placement === "bottom" ? "center top" : "center bottom",
          }}
        >
          <Paper>
            <ClickAwayListener onClickAway={handleClose}>
              <MenuList
                autoFocusItem={open}
                id="menu-list-grow"
                onKeyDown={handleListKeyDown}
              >
                <MenuItem onClick={handleClickOpenMyStories}>
                  <Grid container>
                    <Grid item xs={3}></Grid>
                    <Grid item xs={9}>
                      <div style={{ width: "100%",color: "red"  }}>
                        View story
                      </div>
                    </Grid>
                  </Grid>
                </MenuItem>
                <MenuItem onClick={handleClickAddStory}>
                  <Grid container>
                    <Grid item xs={3}></Grid>
                    <Grid item xs={9}>
                      <div style={{ width: "100%",color: "red"  }}>
                        Add Story
                      </div>
                    </Grid>
                  </Grid>
                </MenuItem>
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

  const showStoryBar=(  <div>
                          <Grid style={{ backgroundColor: "white",width:"70%",height: "5%",overflowX: "auto",display:"flex" ,margin:"auto",marginTop:"2%" }} >
                            <div>
                              <div  ref={anchorRef} onClick={ () => openStories(loggedUserId)} className="cover-image-box">
                                <img src= {"http://localhost:8080/api/user/get-image/" + loggedUserId + ".jpg"} />
                              </div>
                              {"My story"}
                            </div>
                              {dropDowMenuForPost}
                            {users.map((user:any, index) => (
                              <div>
                              <div  onClick={ () => openStories(user.IdString)} className="cover-image-box">
                                <img src= {"http://localhost:8080/api/user/get-image/" + user.IdString + ".jpg"} />
                               
                              </div>
                              {user.Username}
                              </div>
                            ))}
                          </Grid>
                          <AddStory open={openDialog} setOpen={setOpenDialog}></AddStory>
                        </div>)
        
  const showStories=(<div>{stories !== undefined && stories !== null && stories.length !== 0 &&  <Story stories={stories} onClose={closeStory} user={username}></Story>}</div>)



  return (<div>{storiesOpen == true  ? showStories:showStoryBar }</div>);
}