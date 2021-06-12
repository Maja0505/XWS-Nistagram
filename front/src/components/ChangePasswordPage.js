import { makeStyles } from "@material-ui/core/styles";
import React, { useState, useEffect } from "react";
import { Grid, Button, TextField } from "@material-ui/core";
import Avatar from "@material-ui/core/Avatar";
import { deepOrange } from "@material-ui/core/colors";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import axios from "axios";


const useStyles = makeStyles((theme) => ({
    root: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: "center",
    },
    orange: {
      color: theme.palette.getContrastText(deepOrange[500]),
      backgroundColor: deepOrange[500],
      marginLeft: "auto",
    },
  }));

const ChangePasswordPage = ({setOpen,setMessage}) => {
    const classes = useStyles();
    const username = localStorage.getItem("username");
    const [oldPassword,setOldPassword] = useState('')
    const [newPassword,setNewPassword] = useState('')
    const [confirmedNewPassword,setConfirmedNewPassword] = useState('')
    const handleClickChangePassword = () => {
        var passwordDto = {
            oldPassword: oldPassword,
            newPassword: newPassword,
            confirmNewPassword: confirmedNewPassword
        }
        axios.put('/api/user/change-password/' + username, passwordDto)
            .then((res)=>{
                    setOpen(true)
                    setMessage('Successful changed password')
                  
                
            }).catch((err)=> {
                if( err.response && err.response.data){
                    setOpen(true)
                    setMessage(err.response.data)
                }
            
            })
    }

    return (
     <Grid container item xs={9} style={{ height: 600 }}>
      <Grid container item xs={12}>
        <Grid item xs={2}>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            <Avatar className={classes.orange}>N</Avatar>
          </Grid>

          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Old password
          </Grid>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            New password
          </Grid>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Confirm new password
          </Grid>
        </Grid>
        <Grid container item xs={10}>
          <Grid item xs={1}></Grid>
          <Grid item xs={11}>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <p style={{ textAlign: "left", margin: 0, fontSize: 20 }}>
                {username}
              </p>

            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField
                fullWidth
                variant="outlined"
                size="small"
                onChange={(e) =>
                    setOldPassword(e.target.value)
                  }

              />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField
                fullWidth
                variant="outlined"
                size="small"
                onChange={(e) =>
                    setNewPassword(e.target.value)
                  }
              />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField
                fullWidth
                variant="outlined"
                size="small"
                onChange={(e) =>
                    setConfirmedNewPassword(e.target.value)
                  }
              />
            </Grid>
 
            <Grid item style={{ height: "12%", textAlign: "left" }}>
              <Button
             
                color="primary"
                variant="contained"
                onClick = {handleClickChangePassword}
                disabled = {newPassword === '' || oldPassword === '' || confirmedNewPassword === ''}
              >
                Change password
              </Button>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
    )
}

export default ChangePasswordPage
