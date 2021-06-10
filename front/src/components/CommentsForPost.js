import { ListItem } from '@material-ui/core';
import { Avatar, List, makeStyles } from '@material-ui/core'
import { Grid } from '@material-ui/core'
import { deepOrange } from '@material-ui/core/colors';
import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom";




const useStyles = makeStyles((theme) => ({
    orange: {
      color: theme.palette.getContrastText(deepOrange[500]),
      backgroundColor: deepOrange[500],
      marginLeft: "auto",
    },
    margin: {
      margin: theme.spacing(1),
    },
  }));
  
const CommentsForPost = ({commentsForPost,setCommentsForPost}) => {
    const classes = useStyles();
    const { post } = useParams();



    useEffect(() => {
        axios.get('/api/post/get-comments-for-post/' + post)
            .then((res) => {
                setCommentsForPost(res.data)

            })
        
      }, [])

    return (
        <List style={{width:'100%',maxHeight:300}}>
             {commentsForPost !== null && commentsForPost.map((comment) => (
            <ListItem>
                            <Grid item xs={3}>
                    <Avatar
                        className={classes.orange}
                        style={{ margin: "auto" }}
                    >
                        N
                    </Avatar>
                </Grid>
                <Grid item xs={7} style={{textAlign:'left'}}>
            
                    {comment.Content}

                </Grid>
                <Grid item xs={2}></Grid>
            </ListItem>
            
            ))}
        </List>
    )
}

export default CommentsForPost
