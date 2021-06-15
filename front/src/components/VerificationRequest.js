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

const VerificationRequest = ({setOpen,setMessage}) => {
  const username = localStorage.getItem("username");
  const loggedUserId = localStorage.getItem("id");

  const classes = useStyles();
  const [selectedFile, setSelectedFile] = useState();
  const [fullName,setFullName] = useState('');
  const [knownAs,setKnownAs] = useState('');
  const [category,setCategory] = useState();
  const [image,setImage] = useState();
  const [updateMode,setUpdateMode] = useState(false)
  const [userHasRequest, setUserHasRequest] = useState(false)

  useEffect(() => {
    axios.get("/api/user/verification-request/" + loggedUserId).then((res) => {
      console.log(res.data)
      if(res.data !== null){
        setFullName(res.data.FullName)
        setKnownAs(res.data.KnowAs)
        setCategory(res.data.Category)
        setImage(res.data.Image)
        setUpdateMode(true)
        setUserHasRequest(true)
       
      }
    }).catch((err)=> {
      console.log('null')
    });

  }, []);

  const HandleUploadClick = (event) => {
    

    var formData = new FormData();
    console.log(event.target.files[0])
    var file = event.target.files[0];
    formData.append('myFile', file)
    const reader = new FileReader();
    var url = reader.readAsDataURL(file);
    reader.onloadend = function (e) {
      setSelectedFile(reader.result);
    }.bind(this);

    setImage(formData)

  };



  const HandleClickOnUpdate = () => {
    setUpdateMode(false)
  }


  const HandleClickOnSend = () => {
    

    var verification_request ={
      User: loggedUserId,
      Username: username,
      FullName: fullName,
      KnowAs: knownAs,
      Image: "" + loggedUserId + ".jpg",
      Category: category,

    }
    if(userHasRequest){
      axios.put('/api/user/verification-request/update/' + loggedUserId, verification_request)
      .then((res) => {
        axios.post('/api/user/verification-request/upload-verification-doc/' + loggedUserId,image,{
          headers: {'Content-Type': 'multipart/form-data'},
        }).then((res)=> {
          setOpen(true)
          setMessage('Successfully update verification request')
          setUpdateMode(true)
          
        })
      
      })  
    }else{
      axios.post('/api/user/verification-request/create', verification_request)
      .then((res) => {
        axios.post('/api/user/verification-request/upload-verification-doc/' + loggedUserId,image,{
          headers: {'Content-Type': 'multipart/form-data'},
        }).then((res)=> {
          setOpen(true)
          setMessage('Successfully sent verification request')
          setUpdateMode(true)
          
        })
      
      })  
    }

    
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
              <TextField disabled={updateMode} fullWidth variant="outlined" size="small" value={fullName} onChange={(event) => setFullName(event.target.value)} />
            </Grid>
            <Grid item xs={12} style={{ height: "12%", textAlign: "right" }} >
              <TextField disabled={updateMode} fullWidth variant="outlined" size="small" value={knownAs} onChange={(event) => setKnownAs(event.target.value)}/>
            </Grid>
            <Grid item xs={12} style={{ height: "20%", textAlign: "right" }}>
              <Select disabled={updateMode} native fullWidth variant="outlined" value={category} onChange={(event) => setCategory(event.target.value)}>
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
            <Grid  container item xs={12} style={{ height: "auto" }}>
              <Grid item xs={6}>
                {!userHasRequest && !selectedFile && <p style={{textAlign: "left" }}>Please attach a photo of your ID</p>}
                {!selectedFile && userHasRequest ?
                <img width="100%" src={"http://localhost:8080/api/user/verification-request/get-image/" + loggedUserId + ".jpg"} />
                : <img width="100%" src={selectedFile} />

                }
              </Grid>
              <Grid item xs={6} style={{ textAlign: "right"}}>
                    <Button disabled={updateMode} variant="contained" component="label">
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
              {updateMode ?
             <Button color="primary" variant="contained" onClick={HandleClickOnUpdate}>
              Update request
             </Button>
              :
              <Button disabled={fullName === '' || knownAs === '' || category === undefined || image === undefined} color="primary" variant="contained" onClick={HandleClickOnSend}>
                Send
              </Button>
              }
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default VerificationRequest;
