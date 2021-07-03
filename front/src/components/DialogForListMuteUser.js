import React,{useState} from "react";
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import MuiDialogActions from "@material-ui/core/DialogActions";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import { Grid, Divider } from "@material-ui/core";
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { makeStyles } from '@material-ui/core/styles';
import {useEffect } from "react";
import axios from "axios";

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
    width: '100%',
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

const DialogActions = withStyles((theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(1),
  },
}))(MuiDialogActions);

export default function DialogForListMuteUser({loggedUserId,open, setOpen }) {

  const classes = useStyles();
  const [allMuted,setAllMuted] = useState([])

  useEffect(() => {
   axios.get("/api/user-follow/allMuteFriends/" + loggedUserId)
    .then((res) => {
        if(res.data){
            setAllMuted(res.data)
        }
    })
  }, []);

  const handleClose = () => {
    setOpen(false);
  };
  const handleClickCancel = () => {
    setOpen(false);
  }

  const handleClickUnmute = (muted) => {
    var muteDto = {
        User: loggedUserId,
        Friend: muted.IdString,
        Mute : false
    }
    axios.put('/api/user-follow/setMuteFriend',muteDto)
    .then((res)=> {
      console.log('uspelo')
      var array = [...allMuted]; // make a separate copy of the array
      var index = array.indexOf(muted)
      if (index !== -1) {
        array.splice(index, 1);
        setAllMuted(array)
      }
    })
  }



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
          style={{ textAlign: "center" }}
        >
        <h3>Muted users</h3>
        </DialogTitle>
        <DialogContent dividers>
        {allMuted !== null && allMuted.map((muted) => (

            <Grid container>
               <Grid item xs={8}>
                    {muted.Username}
                </Grid>
                <Grid item xs={4}>
                    <Button onClick={() => handleClickUnmute(muted)}>Unmute</Button>
                </Grid>
            </Grid>
       ))}

          
        </DialogContent>
      </Dialog>
    </div>
  );
}
