import { Grid, Button, TextField, Divider } from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import React, { useState, useEffect } from "react";
import axios from "axios";


const ProfilePrivacy = ({profileSettings,setProfileSettings,load}) => {
    const username = localStorage.getItem("username");

    useEffect(() => {
     
    }, [profileSettings.Public,profileSettings.MessageRequest,profileSettings.AllowTags])

      const HandleOnChangeAccountPrivacy = () => {
          if(profileSettings.Public){
            console.log(profileSettings.Public)
              axios.put("/api/user/" + username + "/public-profile/false" )
                .then((res)=> {
                    setProfileSettings({...profileSettings, Public:false})
                })
          }else{
            console.log(profileSettings.Public)

            axios.put("/api/user/" + username + "/public-profile/true" )
            .then((res)=> {
              setProfileSettings({...profileSettings, Public:true})
            })
          }

          
      }

      const HandleOnChangeMessageRequest = () => {
        if(profileSettings.MessageRequest){
            axios.put("/api/user/" + username + "/message-request/false" )
              .then((res)=> {
                setProfileSettings({...profileSettings, MessageRequest:false})
              })
        }else{
          axios.put("/api/user/" + username + "/message-request/true" )
          .then((res)=> {
            setProfileSettings({...profileSettings, MessageRequest:true})
          })
        }
    }

    const HandleOnChangeAllowTags = () => {
        if(profileSettings.AllowTags){
            axios.put("/api/user/" + username + "/allow-tags/false" )
              .then((res)=> {
                setProfileSettings({...profileSettings, AllowTags:false})
              })
        }else{
          axios.put("/api/user/" + username + "/allow-tags/true" )
          .then((res)=> {
            setProfileSettings({...profileSettings, AllowTags:true})
          })
        }
    }
  return (
    <Grid container item xs={9} style={{ height: 600 }}>
      <Grid item xs={1}></Grid>
     {load && profileSettings !== undefined && <Grid container item xs={10}>
        <Grid style={{ height: "30%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Account Privacy</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <FormGroup>
              <FormControlLabel
                control={<Checkbox name="checkedF" checked={profileSettings.Public === true} onChange={HandleOnChangeAccountPrivacy} />}
                label="Private Account"
                style={{ fontSize: 15, fontWeight: "bold" }}
              />
            </FormGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
              When your account is private, only people you approve can see your
              photos and videos on Instagram. Your existing followers won't be
              affected.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "30%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Message Request</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <FormGroup>
              <FormControlLabel
                control={<Checkbox name="checkedF" checked={profileSettings.MessageRequest === true} onChange={HandleOnChangeMessageRequest} />}
                label="Allow message request"
                style={{ fontSize: 15, fontWeight: "bold" }}
              />
            </FormGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
                Let people send you message request.
            </p>
          </Grid>
          <Divider />
        </Grid>
        <Grid style={{ height: "30%", width: "100%" }}>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 25, textAlign: "left" }}>Tag settings</p>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <FormGroup>
              <FormControlLabel
                control={<Checkbox name="checkedF" checked={profileSettings.AllowTags === true} onChange={HandleOnChangeAllowTags}/>}
                label="Allow tags"
                style={{ fontSize: 15, fontWeight: "bold" }}
              />
            </FormGroup>
          </Grid>
          <Grid style={{ height: "30%", width: "100%" }}>
            <p style={{ fontSize: 13, textAlign: "left" }}>
                Let people tag you on posts,stories and albums.
            </p>
          </Grid>
          <Divider />
        </Grid>
      </Grid>}
      <Grid item xs={1}></Grid>
    </Grid>
  );
};

export default ProfilePrivacy;
