import { makeStyles } from "@material-ui/core/styles";
import React, { useState, useEffect } from "react";
import { Grid, Button, TextField } from "@material-ui/core";
import Select from "@material-ui/core/Select";
import axios from "axios";

const useStyles = makeStyles((theme) => ({
  input: {
    display: "none",
  },
}));

const VerificationRequest = () => {
  const username = localStorage.getItem("username");
  const classes = useStyles();
  const [selectedFile, setSelectedFile] = useState();
  const [fullName,setFullName] = useState('');
  const [knownAs,setKnownAs] = useState('');
  const [category,setCategory] = useState();
  const [fileName,setFileName] = useState();

  const HandleUploadClick = (event) => {
    var formData = new FormData();
    var file = event.target.files[0];
    formData.append('myFile', file)
    const reader = new FileReader();
    var url = reader.readAsDataURL(file);

    reader.onloadend = function (e) {
      setSelectedFile(reader.result);
    }.bind(this);

    //console.log(file);
    setFileName(file.name)

    const header = ""
    axios.post('/api/user/verification-request/upload-verification-doc',formData,{
      headers: {'Content-Type': 'multipart/form-data'},
    }).then((res)=> {
      console.log('uspesno')
    })

    /*reader.onloadend = function(e) {
      this.setState({
        selectedFile: [reader.result]
      });
    }.bind(this);
    console.log(url); // Would see a path?*/

    /*this.setState({
      mainState: "uploaded",
      selectedFile: event.target.files[0],
      imageUploaded: 1
    });*/
  };

  const HandleClickOnSend = () => {
    console.log(fullName)
    console.log(knownAs)
    console.log(category)
  }

  return (
    <Grid container item xs={9} style={{ height: 600 }}>
      <Grid container item xs={12} style={{ height: "30%" }}>
          <Grid item xs={1}></Grid>
          <Grid item xs={10}>
              <Grid style={{height:"5%", textAlign: "left"}}>
                <h3>Apply for Instagram Verification</h3>
              </Grid>
              <Grid style={{height:"95%", textAlign: "left"}}>
                <p style={{color:"#868f8b"}}>
                    A verified badge is check that appears next to an Instagram account's name to indicatw that the account is the authentic presence of
                    a notable public figure,celebrity,global brand or entity it represents.

                    Submitting a request for verification does not guarantee that your account will be verified.
                </p>
                <p></p>
              </Grid>
         
          </Grid>
          <Grid item xs={1}></Grid>
      </Grid>
      <Grid container item xs={12}>
        <Grid item xs={2}>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Username
          </Grid>

          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Full name
          </Grid>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Known as
          </Grid>
          <Grid item style={{ height: "12%", textAlign: "right" }}>
            Category
          </Grid>
        </Grid>
        <Grid container item xs={10}>
          <Grid item xs={1}></Grid>
          <Grid item xs={11}>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField
                fullWidth
                variant="outlined"
                size="small"
                disabled
                value={username}
              />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }}>
              <TextField fullWidth variant="outlined" size="small" onChange={(event) => setFullName(event.target.value)} />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }} >
              <TextField fullWidth variant="outlined" size="small"  onChange={(event) => setKnownAs(event.target.value)}/>
            </Grid>
            <Grid item xs={12} style={{ height: "20%", textAlign: "right" }}>
              <Select native fullWidth variant="outlined" onChange={(event) => setCategory(event.target.value)}>
                <option aria-label="None" value="">
                  Select a category for your account
                </option>
                <option value="Blogger/Influencer">Blogger/Influencer</option>
                <option value="Sports">Sports</option>
                <option value="News/Media">News/Media</option>
                <option value="Business/Brand/Organization">
                  Business/Brand/Organization
                </option>
                <option value="Government/Politics">Government/Politics</option>
                <option value="Music">Music</option>
                <option value="Fashion">Fashion</option>
                <option value="Entertainment">Entertainment</option>
                <option value="Other">Other</option>
              </Select>
            </Grid>
            <Grid container item xs={12} style={{ height: "40%" }}>
              <Grid item xs={6}>
                {!selectedFile && <p style={{textAlign: "left" }}>Please attach a photo of your ID</p>}
                <img width="100%" src={selectedFile} />
              </Grid>
              <Grid item xs={6} style={{ textAlign: "right"}}>
                    <Button variant="contained" component="label">
                    Choose file
                    <input
                        hidden
                        accept="image/*"
                        className={classes.input}
                        multiple
                        type="file"
                        name="myFile"
                        onChange={(event) => HandleUploadClick(event)}
                    />
                    </Button>
                </Grid>
            </Grid>
            <Grid item xs={12} style={{ height: "30%", textAlign: "left" }}>
              <p style={{color:"#868f8b"}}>
                  We require a government-issued photo ID that shows your name and date of birth(e.g. driver's license, passport or national identification
                  card) or official business documents( tax filing, recent utility bill, article of incorporation) in order to review your request.
              </p>
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "left" }}>
              <Button color="primary" variant="contained" onClick={HandleClickOnSend}>
                Send
              </Button>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default VerificationRequest;
