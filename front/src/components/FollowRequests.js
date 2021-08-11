import React, { useState } from "react";
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import MuiDialogActions from "@material-ui/core/DialogActions";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import CircularProgress from "@material-ui/core/CircularProgress";
import Typography from "@material-ui/core/Typography";
import { Grid, Divider } from "@material-ui/core";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import { makeStyles } from "@material-ui/core/styles";
import { useEffect } from "react";
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

const DialogActions = withStyles((theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(1),
  },
}))(MuiDialogActions);

export default function FollowRequest({
  loggedUserId,
  open,
  setOpen,
  setAllFollowers,
}) {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const classes = useStyles();
  const [allRequests, setAllRequests] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios
      .get("/api/user-follow/allFollowRequests/" + loggedUserId, authorization)
      .then((res) => {
        if (res.data) {
          setAllRequests(res.data);
        }
        setLoading(false);
      })
      .catch((error) => {
        //console.log(error);
      });
  }, []);

  const handleClose = () => {
    setOpen(false);
  };

  const handleClickAccept = (request) => {
    var requestDto = {
      User: loggedUserId,
      UserWitchSendRequest: request.IdString,
    };
    axios
      .put("/api/user-follow/acceptFollowRequest", requestDto, authorization)
      .then((res) => {
        console.log("uspelo");
        var array = [...allRequests]; // make a separate copy of the array
        var index = array.indexOf(request);
        setAllFollowers((prevState) => [
          ...prevState,
          { IdString: request.IdString, Username: request.Username },
        ]);
        if (index !== -1) {
          array.splice(index, 1);
          setAllRequests(array);
        }
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleClickCancel = (request) => {
    var requestDto = {
      User: loggedUserId,
      UserWitchSendRequest: request.IdString,
    };
    axios
      .put("/api/user-follow/cancelFollowRequest", requestDto, authorization)
      .then((res) => {
        console.log("uspelo");
        var array = [...allRequests]; // make a separate copy of the array
        var index = array.indexOf(request);
        if (index !== -1) {
          array.splice(index, 1);
          setAllRequests(array);
        }
      })
      .catch((error) => {
        console.log(error);
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
          style={{ width: 400 }}
        >
          <h3>Follow Requests</h3>
        </DialogTitle>
        {loading && (
          <DialogContent style={{ width: 400, height: 400 }} dividers>
            <CircularProgress
              disableShrink
              style={{ marginLeft: "180px", marginTop: "180px" }}
            />
          </DialogContent>
        )}
        {loading === false && (
          <DialogContent style={{ width: 400, height: 400 }} dividers>
            {allRequests === undefined ||
              allRequests === null ||
              (allRequests !== null && allRequests.length === 0 && (
                <p>No follow requests</p>
              ))}
            {allRequests !== null &&
              allRequests.length !== 0 &&
              allRequests.map((request) => (
                <Grid container>
                  <Grid item xs={4}>
                    {request.Username}
                  </Grid>
                  <Grid item xs={4}>
                    <Button onClick={() => handleClickAccept(request)}>
                      Accept
                    </Button>
                  </Grid>
                  <Grid item xs={4}>
                    <Button onClick={() => handleClickCancel(request)}>
                      Cancel
                    </Button>
                  </Grid>
                </Grid>
              ))}
          </DialogContent>
        )}
      </Dialog>
    </div>
  );
}
