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

export default function OneTimeMessage({open, setOpen,message,user}) {
  const classes = useStyles();
  const handleClose = () => {
      axios.put("/api/message/user/" + user + "/channels/" + message.channel + "/update",message)
        .then((res) => {
            setOpen(false);
        }).catch((error) => {
          //console.log(error);
        });
  };
  const handleClickCancel = () => {
    setOpen(false);
  }

  return (
    <div>
      <Dialog
        onClose={handleClose}
        aria-labelledby="customized-dialog-title"
        open={open}
      >

        <DialogContent
        dividers>
                         {message.content_id.substring(message.content_id.length - 3, message.content_id.length) ===
                          "jpg" && (
                          <img
                            src={
                              "http://localhost:8080/api/media/get-media-image/" +
                              message.content_id
                            }
                            style={{ width: "100%", height: "600px" }}
                          />
                        )}
                        {message.content_id.substring(message.content_id.length - 3, message.content_id.length) !==
                          "jpg" && (
                          <video
                            width="100%"
                            height="100%"
                            style={{ marginTop: "25%" }}
                            controls
                          >
                            <source
                              src={
                                "http://localhost:8080/api/media/get-video/" +
                                message.content_id
                              }
                              style={{ width: "100%", height: "100%" }}
                              type="video/mp4"
                            />
                          </video>
                        )}


        </DialogContent>
      </Dialog>
    </div>
  );
}
