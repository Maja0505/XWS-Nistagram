import "./ContentDetails.css";
import Button from "../components/Button";
import avatar from "../images/nistagramAvatar.jpg";
import React, { useState, useEffect } from "react";
import Story from "../components/Story";
import { Divider } from "@material-ui/core";
import {
	AppBar,
	Toolbar,
	Grid,
	TextField,
	Typography,
	Avatar,
  } from "@material-ui/core";





export default function ContentDetails() {
  const stories = [
    {
      "profile_name": "Jeremy",
      "video_url": "assets/penguins-1.mp4",
      "duration": 7,
      "type":"video",
    },
    {
      "profile_name": "Aron",
      "video_url": 'https://picsum.photos/432/768',
      "duration": 7,
      "type":"image"
    },
    {
      "profile_name": "Meko222",
      "video_url": "assets/penguins-2.mp4",
      "duration": 7,
      "type":"video",
    },
    {
      "profile_name": "Chupachup",
      "video_url": "assets/penguins-3.mp4",
      "duration": 7,
      "type":"video"
    },
    {
      "profile_name": "Aron",
      "video_url": "assets/penguins-4.mp4",
      "duration": 7,
      "type":"video"
    }
  ]

  const storyClicked = (story:any) => {
    console.log(story)
  }

  const [storiesOpen, setStoriesOpen] = useState(false);
  
  function openStories() {
    setStoriesOpen(true);
  };

  function closeStory() {
    setStoriesOpen(false);
  }

  const showStoryBar=(  <div>
                          <Grid style={{ backgroundColor: "white",width:"80%",height: "5%",marginTop: "1%",overflowX: "auto",display:"flex"  }} >
                            {stories.map((story, index) => (
                              <div  onClick={ () => openStories()} className="cover-image-box">
                                <img src={avatar} onClick={ () => storyClicked(story)} />
                              </div>
                            ))}
                          </Grid>
                        </div>)
        
  const showStories=(<div><Story stories={stories} onClose={closeStory}/></div>)
  
  return (<div>{storiesOpen == true  ? showStories:showStoryBar }</div>);
}