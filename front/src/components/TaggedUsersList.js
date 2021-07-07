import {
  Dialog,
  DialogTitle,
  DialogContent,
  Grid,
  Button,
} from "@material-ui/core";
import { useEffect } from "react";

const TaggedUsersList = ({ label, setOpen, users, open, setTaggedUsers,setListOfTaggedUserid,listOfTaggedUserid }) => {
  useEffect(() => {}, []);

  const handleClose = () => {
    setOpen(false);
  };

  const removeTaggedUser = (user) => {
    var array = [...users];
    var index = array.indexOf(user);
    var array2 = [...listOfTaggedUserid];

    if (index !== -1) {
      array.splice(index, 1);
      array2.splice(index,1)
      setListOfTaggedUserid(array2)
      setTaggedUsers(array);
    }
    
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
          style={{ textAlign: "center", width: 400 }}
        >
          <h3>{label}</h3>
        </DialogTitle>
        <DialogContent style={{ width: 400, height: 400 }}>
          {users !== undefined &&
            users !== null &&
            users.map((user) => (
              <Grid container>
                <Grid item xs={8} style={{ margin: "auto", textAlign: "left" }}>
                  {user.substring(1)}
                </Grid>
                <Grid item xs={4} style={{ margin: "auto", textAlign: "left" }}>
                  <Button
                    color="secondary"
                    onClick={() => removeTaggedUser(user)}
                  >
                    Remove
                  </Button>
                </Grid>
              </Grid>
            ))}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default TaggedUsersList;
