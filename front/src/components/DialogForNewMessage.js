import { useState } from "react";
import { withStyles } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import { Divider } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { connect, sendMsg } from "../api/index";

import {

    Grid,
    TextField,
    Avatar,
    Button

  } from "@material-ui/core";
  import { Autocomplete } from "@material-ui/lab";
  import axios from "axios";
  import avatar from "../images/nistagramAvatar.jpg";
  import {
 
    RoomRounded, ThreeDRotation,
  } from "@material-ui/icons";
const styles = (theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(2),
  },
  closeButton: {
    position: "absolute",
    right: theme.spacing(1),
    top: theme.spacing(1),
    color: theme.palette.grey[500],
  },
});

const useStyles = makeStyles((theme) => ({
  root: {
    width: "100%",
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  },
}));

const DialogTitle = withStyles(styles)((props) => {
  const { children, classes, onClose, ...other } = props;
  return (
    <MuiDialogTitle disableTypography className={classes.root} {...other}>
      <Typography variant="h6">{children}</Typography>
      {onClose ? (
        <IconButton
          aria-label="close"
          className={classes.closeButton}
          onClick={onClose}
        >
          <CloseIcon />
        </IconButton>
      ) : null}
    </MuiDialogTitle>
  );
});

const DialogContent = withStyles((theme) => ({
  root: {
    padding: theme.spacing(2),
  },
}))(MuiDialogContent);

export default function DialogForNewMessage({ open, setOpen }) {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };
  const classes = useStyles();
  const [inappropriate, setInappropriate] = useState(false);
  const username = localStorage.getItem("username");
  const loggedUserId = localStorage.getItem("id");
  const [userForNewMessage,setUserForNewMessage] = useState()

  const [searchedContent, setSearchedContent] = useState([]);
  const [text,setText] = useState("")

  const handleClose = () => {
    setOpen(false);
  };

  const handleChangeInput = (text) => {
    if (text.length !== 0) {
        axios
          .get("/api/user/search/" + username + "/" + text)
          .then((res) => {
            
            setSearchedContent(res.data);
          })
          .catch((error) => {
            setSearchedContent([]);
          });
    } else {
      setSearchedContent([]);
    }
  };

  const sendMessage = () => {
    var user = {}
    axios.get("/api/user/" + userForNewMessage,authorization)
        .then((res) => {
          user = res.data
          sendMsg('{"id":true' + ',"command": 2, "channel":"' + loggedUserId +   '-' +  user.ID  + '", "content": "","opened":false,"type":0,"text":"' +  text + '","user_from":"' + username + '","user_to":"' + user.Username + '"}')
          
                
            
            
        
        
        }).catch((error) => {
          //console.log(error);
        });

    

  
    


    
  }


  const goToSearchContent = (content) => {
    if (content !== null) {
        setUserForNewMessage(content.Username)
    }
  };



  const searchBar = (
    <Grid item xs={6} style={{ textAlign: "center" }}>
      <Autocomplete
        freeSolo
        renderOption={(option) => (
          <Grid container>
            <Grid item xs={2}>
              { option.Username !== undefined && (
                <Avatar
                  alt="N"
                  src={avatar}
                  style={{ border: "1px solid" }}
                ></Avatar>
              )}
              {option.Username === undefined && (
                <Avatar
                  alt="N"
                  style={{
                    backgroundColor: "#ECECEC",
                    border: "1px solid black",
                    color: "black",
                  }}
                >
                  <RoomRounded />
                </Avatar>
              )}
            </Grid>
            <Grid item xs={10} style={{ marginTop: "3%" }}>
              {option.Username !== undefined ? option.Username : option}
            </Grid>
          </Grid>
        )}
        options={
          searchedContent !== null && searchedContent.length !== 0
            ? searchedContent.map((o) => o)
            : []
        }
        getOptionLabel={(option) =>
          option.Username !== undefined ? option.Username : option
        }
        onChange={(event, value) => goToSearchContent(value)}

        renderInput={(params) => (
          <>
            <TextField
              {...params}
              variant="outlined"
              size="small"
              style={{ width: "70%" }}
              onChange={(e) => handleChangeInput(e.target.value)}
            ></TextField>
          </>
        )}
      />
    </Grid>
  );


  return (
    <div>
      <Dialog
        onClose={handleClose}
        aria-labelledby="customized-dialog-title"
        open={open}
      >
        <DialogTitle
          id="customized-dialog-title"
          onClose={handleClose}
          style={{ textAlign: "center",width:400 }}
        >
          Report
        </DialogTitle>
        <DialogContent dividers>
          <h3>New Message</h3>
          <Divider />
            Search user:{searchBar}
            <TextField multiline style={{height:100}} value={text} onChange={(e) => setText(e.target.value)}>

            </TextField>
            <Button onClick={sendMessage}>Send</Button>
        </DialogContent>
      </Dialog>
    </div>
  );
}
