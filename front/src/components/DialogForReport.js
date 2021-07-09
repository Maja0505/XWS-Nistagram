import { useState } from "react";
import { withStyles } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import MuiDialogActions from "@material-ui/core/DialogActions";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import { Divider } from "@material-ui/core";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import { makeStyles } from "@material-ui/core/styles";
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

export default function DialogForReport({ loggedUserId, post, open, setOpen }) {
  const classes = useStyles();
  const [inappropriate, setInappropriate] = useState(false);
  const handleClose = () => {
    setOpen(false);
  };
  const handleInappropriateButton = () => {
    setInappropriate(true);
  };

  const handleClickReport = (description) => {
    var reportedContentDto = {
      Description: description,
      ContentId: post,
      UserId: loggedUserId,
      AdminId: "60cb4d91f5c97c3aa5894ab3", //ispraviti
    };
    axios.post("/api/post/report-content", reportedContentDto).then((res) => {
      console.log("uspelo");
      console.log(reportedContentDto);
      setInappropriate(false);
      setOpen(false);
    });
  };

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
          Report
        </DialogTitle>
        <DialogContent dividers>
          <h3>Why are you reporting this post?</h3>
          <Divider />

          <List
            component="nav"
            className={classes.root}
            aria-label="mailbox folders"
          >
            {!inappropriate ? (
              <>
                <ListItem button>
                  <ListItemText
                    primary="It's spam"
                    onClick={() => handleClickReport("It's spam")}
                  />
                </ListItem>
                <Divider />
                <ListItem button divider onClick={handleInappropriateButton}>
                  <ListItemText primary="It's inappropriate" />
                </ListItem>
              </>
            ) : (
              <>
                <ListItem
                  button
                  onClick={() => handleClickReport("Nudity or sexual activity")}
                >
                  <ListItemText primary="Nudity or sexual activity" />
                </ListItem>
                <Divider />
                <ListItem
                  button
                  divider
                  onClick={() => handleClickReport("Hate speech or symbol")}
                >
                  <ListItemText primary="Hate speech or symbols" />
                </ListItem>
                <ListItem
                  button
                  onClick={() =>
                    handleClickReport("Violence or dangerous organizations")
                  }
                >
                  <ListItemText primary="Violence or dangerous organizations" />
                </ListItem>
                <Divider />
                <ListItem
                  button
                  divider
                  onClick={() =>
                    handleClickReport("Sale of illegal or regulated goods")
                  }
                >
                  <ListItemText primary="Sale of illegal or regulated goods" />
                </ListItem>
                <ListItem
                  button
                  onClick={() => handleClickReport("Bullying or harassment")}
                >
                  <ListItemText primary="Bullying or harassment" />
                </ListItem>
                <Divider />
                <ListItem
                  button
                  divider
                  onClick={() =>
                    handleClickReport("Intellectual property violation")
                  }
                >
                  <ListItemText primary="Intellectual property violation" />
                </ListItem>
                <ListItem
                  button
                  onClick={() => handleClickReport("Suicide or self-injury")}
                >
                  <ListItemText primary="Suicide or self-injury" />
                </ListItem>
                <Divider />
                <ListItem
                  button
                  divider
                  onClick={() => handleClickReport("Eating disorders")}
                >
                  <ListItemText primary="Eating disorders" />
                </ListItem>
                <ListItem
                  button
                  divider
                  onClick={() => handleClickReport("Scam or fraud")}
                >
                  <ListItemText primary="Scam or fraud" />
                </ListItem>
                <ListItem
                  button
                  divider
                  onClick={() => handleClickReport("False inforamtion")}
                >
                  <ListItemText primary="False inforamtion" />
                </ListItem>
                <ListItem
                  button
                  divider
                  onClick={() => handleClickReport("I just don't like it")}
                >
                  <ListItemText primary="I just don't like it" />
                </ListItem>
              </>
            )}
          </List>
        </DialogContent>
      </Dialog>
    </div>
  );
}
