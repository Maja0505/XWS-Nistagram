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

export default function DialogForBlockUser({loggedUserId, blockedUserId ,open, setOpen }) {
  const classes = useStyles();
  const [inappropriate,setInappropriate] = useState(false)
  const handleClose = () => {
    setOpen(false);
  };
  const handleClickCancel = () => {
    setOpen(false);
  }

  const handleClickBlockUser = () => {
    var blockDto = {
        User: loggedUserId,
        BlockedUser: blockedUserId
    }
    axios.post('/api/user-follow/blockUser',blockDto)
    .then((res)=> {
      console.log('uspelo')
      setOpen(false)
    }).catch((error) => {
      //console.log(error);
    });
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
          style={{ textAlign: "center",width:400 }}
        >
        <h3> Block wajwai?</h3>
        <p style={{textAlign:"left"}}>They won't be able to find your profile, posts or story on Instagram. Instagram won't let them know you blocked them.</p>
        </DialogTitle>
        <DialogContent
        dividers>
          

          <List
            component="nav"
            className={classes.root}
            aria-label="mailbox folders"
          >

           <ListItem button>
              <ListItemText primary="Block" onClick={handleClickBlockUser}/>
            </ListItem>
            <Divider />
            <ListItem button divider onClick={handleClickCancel}>
              <ListItemText primary="Cancel" />
            </ListItem>
        

          </List>
        </DialogContent>
      </Dialog>
    </div>
  );
}
