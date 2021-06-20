import "./StoryBar.css";
import Button from ".././Button";
import avatar from "../../images/nistagramAvatar.jpg";
import React, { useState, useEffect } from "react";
import Story from "./Story";
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
  const [storyIndex,setStoryIndex] = useState(0)
  const [stories,setStories] = useState([
    {
      "opened":true,
      "profile_name": "Jeremy",
      "video_url": "assets/penguins-1.mp4",
      "duration": 7,
      "type":"video",
    },
    {
      "opened":false,
      "profile_name": "Aron",
      "video_url": 'https://picsum.photos/432/768',
      "duration": 7,
      "type":"image"
    },
    {
      "opened":false,
      "profile_name": "Aron",
      "video_url": "assets/penguins-2.mp4",
      "duration": 7,
      "type":"video",
    },
    {
      "opened":false,
      "profile_name": "Chupachup",
      "video_url": "assets/penguins-3.mp4",
      "duration": 7,
      "type":"video"
    },
    { "opened":true,
      "profile_name": "Aron",
      "video_url": "assets/penguins-4.mp4",
      "duration": 7,
      "type":"video"
    }
  ]);

  const storyClicked = (index:any) => {
    setStoryIndex(index)
    const updatedStories = stories.map((item) => {
      if (item === stories[index]) {
        const updatedItem = {
        ...item,
        opened: true,
        };
     
        return updatedItem;
      }
     
      return item;
      });
     
      setStories(updatedStories);
  }

  const [storiesOpen, setStoriesOpen] = useState(false);
  
  function openStories() {
    setStoriesOpen(true);
  };

  function closeStory() {
    setStoriesOpen(false);
  }

  const showStoryBar=(  <div>
                          <Grid style={{ backgroundColor: "transparent",width:"80%",height: "5%",marginTop: "1%",overflowX: "auto",display:"flex"  }} >
                            {stories.map((story, index) => (
                              <div  onClick={ () => openStories()} className={story.opened == false ? "cover-image-box-unopened":"cover-image-box-opened"}>
                                <img src={avatar} onClick={ () => storyClicked(index)} />
                              </div>
                            ))}
                          </Grid>
                        </div>)
        
  const showStories=(<div><Story currentStoryIndex={storyIndex} stories={stories} onClose={closeStory}/></div>)
  
  return (<div>{storiesOpen == true  ? showStories:showStoryBar }</div>);
}