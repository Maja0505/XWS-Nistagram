import React, { Component } from 'react';
import Stories from 'react-insta-stories';
import {
	AppBar,
	Toolbar,
	Grid,
	Button,
	TextField,
	Typography,
	Avatar,
  } from "@material-ui/core";
import { useState, useEffect } from "react";
import avatar from "../images/nistagramAvatar.jpg";
import { WithHeader } from 'react-insta-stories';




const StoryBar = () => {
const [showStory , setShowStory]=useState(false)
const [currentStoryIndex,setCurrentStoryIndex]=useState(1)

const story = {
	id:"1",
	url: 'https://picsum.photos/432/768',
	duration: 5000,
	header: {
		heading: 'Marko Markovic',
		subheading: 'Posted 30m ago',
		profileImage: 'https://picsum.photos/100/100',
	},
	opened:"false",
}
	
const [stories, setStories ] = useState([
	{
		id:"1",
		url: 'https://picsum.photos/432/768',
		duration: 5000,
		header: {
			heading: 'Marko Markovic',
			subheading: 'Posted 30m ago',
			profileImage: 'https://picsum.photos/100/100',
		},
		opened:"false",
	},
	{
		id:"2",
		url: 'https://picsum.photos/432/768',
		duration: 5000,
		header: {
			heading: 'Zoran Zoric',
			subheading: 'Posted 30m ago',
			profileImage: 'https://picsum.photos/100/100',
		},
		opened:"true",
	},
	{
		id:"3",
		url: 'https://picsum.photos/432/768',
		duration: 5000,
		header: {
			heading: 'Mohit Karekar',
			subheading: 'Posted 30m ago',
			profileImage: 'https://picsum.photos/100/100',
		},
		opened:"true",
	},
	{
		id:"4",
		url: 'https://picsum.photos/432/768',
		duration: 5000,
		header: {
			heading: 'Pera Petricevic',
			subheading: 'Posted 30m ago',
			profileImage: 'https://picsum.photos/100/100',
		},
		opened:"false",
	},
	
]);

const storyClicked = (story) => {
	setShowStory(true)
	setCurrentStoryIndex(story.index)
	const updatedStories = stories.map((item) => {
		if (item.id === story.id) {
		  const updatedItem = {
			...item,
			opened: false,
		  };
   
		  return updatedItem;
		}
   
		return item;
	  });
   
	  setStories(updatedStories);
  };


  const storiesComponent = (
	<Stories
		stories={stories}
		defaultInterval={1500}
		width="100%"
		height="100%"
		currentIndex = "3"
		onAllStoriesEnd={() => closeStories()}
	/>)

	const closeStories= () => {
		setShowStory(false)
	}

	const CustomStoryContent = ( <WithHeader story={story} globalHeader="tralalla">
			<div>
				<h1>Hello!</h1>
				<p>This story would have the configured header!</p>
			</div>
		</WithHeader>)

  const storyToolbar = (<Grid style={{ backgroundColor: "white",width:"80%",height: "5%",marginTop: "1%",overflowX: "auto",border:"1px solid",borderColor:"slategray"  }}>
  <div style={{display:"flex"}}>
 {stories.map((story, index) => (
		 <img
			   onClick={ () => storyClicked(story)}
			   key={story.id}
			 src={avatar}
			 alt="Not founded"
			 style={{
			 borderRadius: "50%",
			 border: "5px solid",
			 borderColor: story.opened==="true" ? "orangered" : "whitesmoke",
			 height:"60%",
			 width:"10%",
			 marginLeft:"1%",
			 }}
		 />
   
 ))}
</div>
</Grid>)
    
  return  (<div>{showStory === true ? storiesComponent : storyToolbar}</div>); };
  
  export default StoryBar;