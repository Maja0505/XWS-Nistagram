import * as React from "react";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import {
    Grid,
    Typography,
    Paper
  } from "@material-ui/core";

const ProfileDialog = ({openD}) => {
  const [open, setOpen] = React.useState(true);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
    
        <Paper   open={open}
        onClose={handleClose}>
            <Grid container style={{width:"100%",height:500}}>
                <Grid container>
                    <Grid item></Grid>
                </Grid>
            </Grid>
        </Paper>
    </div>
  );
};

export default ProfileDialog;
