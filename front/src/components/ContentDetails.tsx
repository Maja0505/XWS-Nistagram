import "./ContentDetails.css";
import avatar from "../images/nistagramAvatar.jpg";

import  { useState, useEffect,useRef } from "react";
import Story from "./Story";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";

import {
	Grid,
  Paper,
  Grow,
  Popper,
  MenuItem,
  MenuList,
  } from "@material-ui/core";
  import AddStory from "./AddStoryDialog"
  import AddStoryCampaign from "./AddStoryCampaign"

import axios from "axios";




export default function ContentDetails() {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };
  const [username,setUsername] =  useState("")
  const [users,setUsers] = useState([])
  const loggedUserId = localStorage.getItem("id");
  const [haveProfileImage,setHaveProfileImage] = useState(false)
  const [haveStory,setHaveStory] = useState(false)
  const isAgent = localStorage.getItem("isAgent");
  const [openDialogForCampaign,setOpenDialogForCampaign] = useState(false)


  useEffect(() => {
    axios.get("/api/post/story/all-follows-with-stories/" + loggedUserId,authorization)
    .then((res) => {
      if(res.data){
        setUsers(res.data)
      }
    }).catch((error) =>{

    })
   axios.get("/api/post/story/all-for-close-friends/" + loggedUserId,authorization)
    .then((res) => {
      if(res.data){
        console.log(res.data)
        setStories(res.data)
        setHaveStory(true)
      }
     
    }).catch((error) =>{
      
    })

    axios.get("/api/media/get-profile-picture/" + loggedUserId + ".jpg").then((res) => {
      setHaveProfileImage(true)
    }).catch(error =>{
     
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
      axios.get("/api/user-follow/checkClosed/"+loggedUserId+"/" + username,authorization)
        .then((res) => {
          if(res.data){
            axios.get("/story/all-for-close-friends/" + username,authorization)
              .then((res) => {
                setStories(res.data)
                setStoriesOpen(true)
              }).catch((error) =>{
      
              })
          }else{
            axios.get("/api/post/story/all-not-expired/" + username,authorization)
            .then((res) => {
              setStories(res.data)
              setStoriesOpen(true)

            }).catch((error) =>{
      
            })
          }
        }).catch((error) => {
          //console.log(error);
        });
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

  const handleClickAddStoryCampaign = () => {
    setOpenDialogForCampaign(true)
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClickOpenMyStories = () => {
    axios.get("/api/post/story/all-for-close-friends/" + loggedUserId,authorization)
    .then((res) => {
      console.log(res.data)
      setStories(res.data)
      setStoriesOpen(true);
      setOpen((prevOpen) => !prevOpen);
    }).catch((error) =>{
      
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
                {haveStory &&
                <MenuItem onClick={handleClickOpenMyStories}>
                  <Grid container>
                    <Grid item xs={3}></Grid>
                    <Grid item xs={9}>
                      <div style={{ width: "100%",color: "red"  }}>
                        View story
                      </div>
                    </Grid>
                  </Grid>
                </MenuItem>}
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
                {isAgent &&
                <MenuItem onClick={handleClickAddStoryCampaign}>
                <Grid container>
                  <Grid item xs={3}></Grid>
                  <Grid item xs={9}>
                    <div style={{ width: "100%",color: "red"  }}>
                      Create campaign
                    </div>
                  </Grid>
                </Grid>
               </MenuItem>    
                }
              </MenuList>
            </ClickAwayListener>
          </Paper>
        </Grow>
      )}
    </Popper>
  );

  const showStoryBar=(  <div> 
                          <Grid style={{ backgroundColor: "white",overflowX: "auto",display:"flex" ,margin:"auto", border:"0.5px solid",borderColor: "#b9b9b9",}} >
                          <div>
                              <div   style={{width:"20px"}}>
                              </div>
                              
                            </div>
                            <div style={{fontSize:"13px",marginTop:"1.2%",marginBottom:"1.2%"}}>
                              <div  ref={anchorRef} onClick={ () => openStories(loggedUserId)} className="cover-image-box">
                                {haveProfileImage && <img src= {"http://localhost:8080/api/media/get-profile-picture/" + loggedUserId + ".jpg"} style={{cursor:"pointer"}}/>}
                                {!haveProfileImage && <img src= {avatar} style={{cursor:"pointer"}}/>}
                                
                              </div>
                              {"My story"}
                            </div>
                              {dropDowMenuForPost}
                            {users.map((user:any, index) => (
                              <div key={index} style={{fontSize:"13px",marginTop:"1.2%",marginBottom:"1.2%"}}>
                              <div  onClick={ () => openStories(user.IdString)} className="cover-image-box">
                                <img src= {"http://localhost:8080/api/media/get-profile-picture/" + user.IdString + ".jpg"} style={{cursor:"pointer"}}/>
                               
                              </div>
                              {user.Username}
                              </div>
                            ))}
                          </Grid>
                          
                          <AddStory open={openDialog} setOpen={setOpenDialog} setHaveStory={setHaveStory}></AddStory>
                          <AddStoryCampaign open={openDialogForCampaign} setOpen={setOpenDialogForCampaign} setHaveStory={setHaveStory}></AddStoryCampaign>
                        </div>)
        
  const showStories=(<div>{stories !== undefined && stories !== null && stories.length !== 0 &&  <Story stories={stories} onClose={closeStory} user={username}></Story>}</div>)



  return (<div>{storiesOpen == true  ? showStories:showStoryBar }</div>);
}