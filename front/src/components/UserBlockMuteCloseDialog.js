import {
  Dialog,
  DialogTitle,
  DialogContent,
  Grid,
  Button,
} from "@material-ui/core";

import { makeStyles } from "@material-ui/core/styles";

import { useEffect, useState } from "react";
import DialogForBlockUser from "./DialogForBlockUser";
import DialogForMuteUser from "./DialogForMuteUser";

import axios from "axios";

const useStyles = makeStyles((theme) => ({
  button: {
    height: "50px",
    width: "100%",
    fontSize: "17px",
    color: "red",
  },
}));

const UserBlockMuteCloseDialog = ({
  open,
  setOpen,
  user,
  relationShip,
  setRelationShip,
  allFollowers,
  setAllFollowers,
}) => {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  const classes = useStyles();
  const loggedUserId = localStorage.getItem("id");

  const [openDialogForBlock, setOpenDialogForBlock] = useState(false);
  const [openDialogForMute, setOpenDialogForMute] = useState(false);

  const handleClose = () => {
    setOpen(false);
  };

  useEffect(() => {}, []);

  const handleOpenDialogForBlock = () => {
    setOpenDialogForBlock(true);
  };

  const handleOpenDialogForMute = () => {
    setOpenDialogForMute(true);
  };

  const unmute = () => {
    var muteDto = {
      User: loggedUserId,
      Friend: user.ID,
      Mute: false,
    };
    axios
      .put("/api/user-follow/setMuteFriend", muteDto, authorization)
      .then((res) => {
        console.log("uspelo");
        setRelationShip({ ...relationShip, IsMuted: false });
        handleClose();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const unblock = () => {
    var blockDto = {
      User: loggedUserId,
      BlockedUser: user.IdString,
    };
    axios
      .put("/api/user-follow/unblockUser", blockDto, authorization)
      .then((res) => {
        console.log("uspelo");
        setRelationShip({ ...relationShip, IsBlocked: false });
        handleClose();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const setToClose = () => {
    var closeDto = {
      User: loggedUserId,
      Friend: user.IdString,
      Close: true,
    };
    axios
      .put("/api/user-follow/setCloseFriend", closeDto, authorization)
      .then((res) => {
        console.log("uspesno");
        setRelationShip({ ...relationShip, IsClosed: true });
        handleClose();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const removeFromClose = () => {
    var closeDto = {
      User: loggedUserId,
      Friend: user.IdString,
      Close: false,
    };
    axios
      .put("/api/user-follow/setCloseFriend", closeDto, authorization)
      .then((res) => {
        console.log("uspesno");
        setRelationShip({ ...relationShip, IsClosed: false });
        handleClose();
      })
      .catch((error) => {
        //console.log(error);
      });
  };

  const buttonForBlock = (
    <Button className={classes.button} onClick={handleOpenDialogForBlock}>
      Block this user
    </Button>
  );

  const buttonForUnblock = (
    <Button className={classes.button} onClick={unblock}>
      Unblock this user
    </Button>
  );

  const buttonForClose = (
    <Button className={classes.button} onClick={setToClose}>
      Set for close friend
    </Button>
  );

  const buttonForUnclose = (
    <Button className={classes.button} onClick={removeFromClose}>
      Delete from close friends
    </Button>
  );

  const buttonForMute = (
    <Button className={classes.button} onClick={handleOpenDialogForMute}>
      Mute this user
    </Button>
  );

  const buttonForUnmute = (
    <Button className={classes.button} onClick={unmute}>
      Unmute this user
    </Button>
  );

  return (
    <div>
      {user !== null && relationShip !== null && (
        <Dialog
          onClose={handleClose}
          aria-labelledby="customized-dialog-title"
          open={open}
        >
          <DialogContent style={{ width: 300, height: 250 }}>
            <div style={{ textAlign: "center", marginTop: "50px" }}>
              {!relationShip.IsBlocked && buttonForBlock}
              {relationShip.IsBlocked && buttonForUnblock}
            </div>
            <div style={{ textAlign: "center" }}>
              {relationShip.IsFollowing &&
                !relationShip.IsMuted &&
                buttonForMute}
              {relationShip.IsFollowing &&
                relationShip.IsMuted &&
                buttonForUnmute}
            </div>
            <div style={{ textAlign: "center" }}>
              {relationShip.IsFollowing &&
                !relationShip.IsClosed &&
                buttonForClose}
              {relationShip.IsFollowing &&
                relationShip.IsClosed &&
                buttonForUnclose}
            </div>
          </DialogContent>
        </Dialog>
      )}

      {user !== null && (
        <DialogForBlockUser
          loggedUserId={loggedUserId}
          blockedUser={user}
          open={openDialogForBlock}
          setOpen={setOpenDialogForBlock}
          setRelationShip={setRelationShip}
          setOpenFirstDialog={setOpen}
          relationShip={relationShip}
          allFollowers={allFollowers}
          setAllFollowers={setAllFollowers}
        ></DialogForBlockUser>
      )}

      {user !== null && (
        <DialogForMuteUser
          loggedUserId={loggedUserId}
          muteUser={user}
          open={openDialogForMute}
          setOpen={setOpenDialogForMute}
          setRelationShip={setRelationShip}
          setOpenFirstDialog={setOpen}
          relationShip={relationShip}
        ></DialogForMuteUser>
      )}
    </div>
  );
};

export default UserBlockMuteCloseDialog;
