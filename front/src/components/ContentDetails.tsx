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






export default function ContentDetails() {
  const users = [
    {
      username:"Jeremy"
    },
    {
      username:"Riko"
    }
  ]
  const [open, setOpen] = useState(false);
  const [openDialog,setOpenDialog] = useState(false)

  const anchorRef = useRef(null);
  const [myStories,setMyStories] = useState([])
  const [stories,setStories] = useState([
    {
      "profile_name": "Jeremy",
      "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
      "duration": 7,
      "type":"video",
      "subheading": 'Posted 30m ago',
			"profile_image": 'https://picsum.photos/100/100',
    },
    {
      "profile_name": "Jeremy",
      "video_url": 'https://picsum.photos/432/768',
      "duration": 7,
      "type":"image",
      "subheading": 'Posted 30m ago',
			"profile_image": 'https://picsum.photos/100/100',
    },
    {
      "profile_name": "Jeremy",
      "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
      "duration": 7,
      "type":"video",
      "subheading": 'Posted 30m ago',
			"profile_image": 'https://picsum.photos/100/100',
    },
    {
      "profile_name": "Jeremy",
      "video_url": "https://picsum.photos/432/768",
      "duration": 7,
      "type":"image",
      "subheading": 'Posted 30m ago',
			"profile_image": 'https://picsum.photos/100/100',
    },
    {
      "profile_name": "Jeremy",
      "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
      "duration": 7,
      "type":"video",
      "subheading": 'Posted 30m ago',
			"profile_image": 'https://picsum.photos/100/100',
    }
  ])

  const storyClicked = (story:any) => {
    console.log(story)
  }

  const [storiesOpen, setStoriesOpen] = useState(false);
  



  const openStories = (username:any) => {
    if(username === "Jeremy"){
      setStories(
        [
          {
            "profile_name": "Jeremy",
            "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
            "duration": 7,
            "type":"video",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          }
        ]
      )
      setStoriesOpen(true);
    }

    if(username === "Riko"){
      setStories(
        [
          {
            "profile_name": "Riko",
            "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
            "duration": 7,
            "type":"video",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          },
          {
            "profile_name": "Riko",
            "video_url": 'https://picsum.photos/432/768',
            "duration": 7,
            "type":"image",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          },
          {
            "profile_name": "Riko",
            "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
            "duration": 7,
            "type":"video",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          },
          {
            "profile_name": "Riko",
            "video_url": "https://picsum.photos/432/768",
            "duration": 7,
            "type":"image",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          },
          {
            "profile_name": "Riko",
            "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
            "duration": 7,
            "type":"video",
            "subheading": 'Posted 30m ago',
            "profile_image": 'https://picsum.photos/100/100',
          }
        ]
      )
      setStoriesOpen(true);

    }
    if(username === "MyStory"){
      setOpen((prevOpen) => !prevOpen);
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
    setStories(
      [
        {
          "profile_name": "My story",
          "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
          "duration": 7,
          "type":"video",
          "subheading": 'Posted 30m ago',
          "profile_image": 'https://picsum.photos/100/100',
        },
        {
          "profile_name": "My story",
          "video_url": 'https://picsum.photos/432/768',
          "duration": 7,
          "type":"image",
          "subheading": 'Posted 30m ago',
          "profile_image": 'https://picsum.photos/100/100',
        },
        {
          "profile_name": "My story",
          "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
          "duration": 7,
          "type":"video",
          "subheading": 'Posted 30m ago',
          "profile_image": 'https://picsum.photos/100/100',
        },
        {
          "profile_name": "My story",
          "video_url": "https://picsum.photos/432/768",
          "duration": 7,
          "type":"image",
          "subheading": 'Posted 30m ago',
          "profile_image": 'https://picsum.photos/100/100',
        },
        {
          "profile_name": "My story",
          "video_url": "http://techslides.com/demos/sample-videos/small.ogv",
          "duration": 7,
          "type":"video",
          "subheading": 'Posted 30m ago',
          "profile_image": 'https://picsum.photos/100/100',
        }
      ]
    )
    setStoriesOpen(true);
    setOpen((prevOpen) => !prevOpen);
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
                              <div  ref={anchorRef} onClick={ () => openStories("MyStory")} className="cover-image-box">
                                <img src={myPic} />
                              </div>
                              {dropDowMenuForPost}
                            {users.map((user, index) => (
                              <div  onClick={ () => openStories(user.username)} className="cover-image-box">
                                <img src={avatar} onClick={ () => storyClicked(user)} />
                              </div>
                            ))}
                          </Grid>
                          <AddStory open={openDialog} setOpen={setOpenDialog}></AddStory>
                        </div>)
        
  const showStories=(<div><Story stories={stories} onClose={closeStory}></Story></div>)



  return (<div>{storiesOpen == true  ? showStories:showStoryBar }</div>);
}