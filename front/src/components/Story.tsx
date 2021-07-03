import React, { useState, useEffect, useRef } from "react";

import "./Story.css";

import MoreHoriz from "@material-ui/icons/MoreHoriz";
import ChevronRight from "@material-ui/icons/ChevronRight";
import Close from "@material-ui/icons/Close";
import ChevronLeft from "@material-ui/icons/ChevronLeft";
import { Story as StoryModel } from "../models/Story";
import { User as UserModel } from "../models/User";

import avatar from "../images/nistagramAvatar.jpg";
import axios from "axios";



interface Props {
  onClose: Function;
  stories: StoryModel[];
  user: string;
}
 
export default function Story({ onClose, stories, user }: Props ){
  const [storyPaused, setStoryPaused] = useState(false);
  const [storyIndex, setStoryIndex] = useState(0);
  const storyIndexRef = useRef(0);
  const [imageProfile,setImageProfile] = useState("")
  const [username,setUsername] = useState("")





  useEffect(() => {
    const video = document.getElementById("video") as HTMLVideoElement;

    if (video) {
      video.onended = (e) => {
        if (storyIndexRef.current === stories.length - 1) {
          onClose();
        } else {
          setStoryIndex((value) => value + 1);
        }
      };
    }
  }, []);

  useEffect(() => {
    storyIndexRef.current = storyIndex;
  }, [storyIndex]);

  useEffect(() => {
    axios.get("/api/user/userid/" + user )
    .then((res)=> {
      setImageProfile(res.data.ProfilePicture)
      setUsername(res.data.Username)
    })
   
    if(stories[storyIndex].Type === "video"){
          if (storyPaused) {
            (document.getElementById("video") as HTMLVideoElement).pause();
          } else {
            (document.getElementById("video") as HTMLVideoElement).play();
          }
    }
  }, [storyPaused]);

  function onClickStory(element: EventTarget) {
    if ((element as HTMLElement).className === "story-container") onClose();
  }



  function getProgressBarClassName(index: number) {
    if (index < storyIndex) {
      return "progress-bar progress-bar-finished";
    } else if (index === storyIndex) {
      return storyPaused ? "progress-bar progress-bar-active  progress-bar-paused" : "progress-bar progress-bar-active";
    } else {
      return "progress-bar";
    }
  }

  const storyVideo = (

<video onMouseDown={(e) => setStoryPaused(true)} onMouseUp={(e) => setStoryPaused(false)} id="video" src= {"http://localhost:8080/api/post/story/video-get/" + stories[storyIndex].Media} autoPlay  width="100%" style={{marginTop:"30%"}} ></video>  )

  const storyImage = (<img onMouseDown={(e) => setStoryPaused(true)} onMouseUp={(e) => setStoryPaused(false)} id="video" src= {"http://localhost:8080/api/post/story/image-get/" + stories[storyIndex].Media} style={{height:"90%",width:"100%"}} ></img>
  )

  return (
    <div onClick={(e) => onClickStory(e.target)} className="story-container">
      <div className="story">
        <div className="title">
          <img src= {"http://localhost:8080/api/user/get-image/" + imageProfile} />
          <div className="details">
            <span>{username}</span>
            <span>{stories[storyIndex].Subheading}</span>
          </div>
          {storyPaused && <span className="pause">PAUSED</span>}
          <MoreHoriz style={{marginLeft: storyPaused==true ? "10px" : "300px"}} />
          <Close style={{marginLeft:"20px"}} onClick={(e) => onClose()}/>
        </div>
        <div className="progress-bars">
          {stories.map((story, index) => (
            <div className="progress-bar-container" id="progress_bar">
              <div style={{ animationDuration: `${story.Duration}s` }} className={getProgressBarClassName(index)}></div>
            </div>
          ))}
        </div>
        <div className= {stories[storyIndex].Type === "image" ? "image" : "video" }>
          {stories[storyIndex].Type === "image" ? storyImage : storyVideo }
          {storyIndex !== 0 && <ChevronLeft onClick={(e) => setStoryIndex((value) => value - 1)} className="previous hoverable" />}
          {storyIndex !== stories.length - 1 && <ChevronRight onClick={(e) => setStoryIndex((value) => value + 1)} className="next hoverable" />}
        </div>
      </div>
    </div>
  );
}