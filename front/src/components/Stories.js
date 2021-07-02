import {useState } from 'react';
import Stories from 'react-insta-stories';
import { Grid} from "@material-ui/core";
import ChevronLeft from "@material-ui/icons/ChevronLeft";
import ChevronRight from "@material-ui/icons/ChevronRight";


const StoriesComponent = () => {
    
    const [currentIndex,setCurrentIndex] = useState(0)

    
    const stories = [		{
		id:"1",
		url: 'https://example.com/vid.mp4',
		duration: 5000,
		header: {
			heading: 'Marko Markovic',
			subheading: 'Posted 30m ago',
			profileImage: 'https://picsum.photos/100/100',
		},
        type:"video",
		opened:"false",
        seeMore: ({ close }) => {
			return <div onClick={close}>Hello, click to close this.</div>;
		},
	}]

    const setStoryToOpen = () => {
        setCurrentIndex(currentIndex+1)
        console.log(currentIndex)
       
    }
        return (
            <div>
        <Grid container>
            <Grid item xs={4}></Grid>
            <Grid item xs={4}>
                <Stories
                    stories={stories}
                    defaultInterval={1500}
                    width={432}
			        height={768}
                    currentIndex={currentIndex}
                    onStoryEnd = {setStoryToOpen}
                    isPaused = "false"
                >
                              {currentIndex !== 0 && <ChevronLeft onClick={(e) => setCurrentIndex(currentIndex - 1)} className="previous hoverable" />}
                              {currentIndex !== stories.length - 1 && <ChevronRight onClick={(e) => setCurrentIndex(currentIndex + 1)} className="next hoverable" />}

                </Stories>
                
            </Grid>
            <Grid item xs={4}></Grid>
        </Grid>
            </div>
    )
}

export default StoriesComponent
