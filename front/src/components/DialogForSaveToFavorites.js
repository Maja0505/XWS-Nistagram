import { useState, useEffect } from "react";
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
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import { makeStyles } from "@material-ui/core/styles";
import { TextField } from "@material-ui/core";
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

export default function DialogForSaveToFavorites({
  loggedUserId,
  post,
  open,
  setOpen,
  saved,
  setSaved,
}) {
  const authorization = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };
  const classes = useStyles();
  const [saveToFavorites, setSaveToFavorites] = useState(true);
  const [saveToCollection, setSaveToCollection] = useState(false);
  const [createNewCollection, setCreateNewCollection] = useState(false);
  const [saveToAnExistingCollection, setSaveToAnExistingCollection] =
    useState(false);
  const [allUserCollections, setAllUserCollections] = useState([]);
  const [collectionName, setCollectionName] = useState("");
  const [collectionsForPost, setCollectionsForPost] = useState([]);
  const [removingFromCollections, setRemovngFromCollections] = useState(false);
  const handleClose = () => {
    setSaveToCollection(false);
    setSaveToAnExistingCollection(false);
    setCreateNewCollection(false);
    setSaveToFavorites(true);
    setOpen(false);
  };

  useEffect(() => {
    axios
      .get(
        "/api/post/get-all-collections-for-post-by-user/" +
          loggedUserId +
          "/" +
          post,authorization
      )
      .then((res) => {
        if (res.data) {
          setCollectionsForPost(res.data);
        }
      }).catch((error) => {
        //console.log(error);
      });;
  }, []);

  const handleClickSavePost = (description) => {
    console.log(post);
    var favoritesDto = {
      PostID: post,
      UserID: loggedUserId,
      Collection: collectionName,
    };
    axios.post("/api/post/add-to-favourites", favoritesDto,authorization).then((res) => {
      if (createNewCollection || saveToAnExistingCollection) {
        axios.post("/api/post/add-to-collection", favoritesDto,authorization).then((res) => {
          console.log("uspelo");
          setOpen(false);
          setSaved(true);
        });
      } else {
        console.log("uspelo");
        setOpen(false);
        setSaved(true);
      }
    }).catch((error) => {
      //console.log(error);
    });
  };

  const handleClickSaveToCollection = () => {
    axios
      .get("/api/post/get-collections-for-user/" + loggedUserId,authorization)
      .then((res) => {
        if (res.data) {
          setAllUserCollections(res.data);
        }
        setSaveToFavorites(false);
        setSaveToCollection(true);
      }).catch((error) => {
        //console.log(error);
      });
  };

  const handleClickCreateNewCollection = () => {
    setSaveToCollection(false);
    setCreateNewCollection(true);
  };

  const handleClickeSaveToAnExistngCollection = () => {
    setSaveToCollection(false);
    setSaveToAnExistingCollection(true);
  };

  const handleClickeRemovingFromAnExistngCollection = () => {
    setSaveToFavorites(false);
    setRemovngFromCollections(true);
  };

  const handleClickRemoveFromFavourites = () => {
    var favoritesDto = {
      PostID: post,
      UserID: loggedUserId,
      Collection: "",
    };
    axios
      .post("/api/post/remove-post-from-favourites", favoritesDto,authorization)
      .then((res) => {
        console.log("uspenso");
        setSaved(false);
        setOpen(false);
      }).catch((error) => {
        //console.log(error);
      });
  };

  const handleClickRemovePostFromCollection = () => {
    var favoritesDto = {
      PostID: post,
      UserID: loggedUserId,
      Collection: collectionName,
    };
    axios
      .post("/api/post/remove-post-from-collection", favoritesDto,authorization)
      .then((res) => {
        console.log("uspenso");

        setOpen(false);
      }).catch((error) => {
        //console.log(error);
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
          Save post
        </DialogTitle>
        <DialogContent dividers>
          <h4>Only you can see what you've saved</h4>
          <Divider />

          <List
            component="nav"
            className={classes.root}
            aria-label="mailbox folders"
          >
            {saveToFavorites && (
              <>
                {!saved && (
                  <ListItem button>
                    <ListItemText
                      primary="Save to favourites"
                      onClick={() => handleClickSavePost("Save to favorites")}
                    />
                  </ListItem>
                )}
                {saved && (
                  <ListItem button>
                    <ListItemText
                      primary="Remove from favourites"
                      styles={{ color: "red" }}
                      onClick={handleClickRemoveFromFavourites}
                    />
                  </ListItem>
                )}
                {saved &&
                  collectionsForPost &&
                  collectionsForPost.length !== 0 && (
                    <ListItem button>
                      <ListItemText
                        primary="Remove from collection"
                        styles={{ color: "red" }}
                        onClick={handleClickeRemovingFromAnExistngCollection}
                      />
                    </ListItem>
                  )}
                <Divider />
                <ListItem button divider onClick={handleClickSaveToCollection}>
                  <ListItemText primary="Save to collection" />
                </ListItem>
              </>
            )}
            {saveToCollection && (
              <>
                <ListItem button onClick={handleClickCreateNewCollection}>
                  <ListItemText primary="Create new collection" />
                </ListItem>
                <Divider />
                {collectionsForPost && allUserCollections.length !== 0 && (
                  <ListItem
                    button
                    divider
                    onClick={handleClickeSaveToAnExistngCollection}
                  >
                    <ListItemText primary="Save to an existing collection" />
                  </ListItem>
                )}
              </>
            )}
            {createNewCollection && (
              <Grid>
                <TextField
                  fullWidth
                  variant="outlined"
                  size="small"
                  onChange={(event) => setCollectionName(event.target.value)}
                />
                <Button
                  disabled={collectionName === ""}
                  onClick={() => handleClickSavePost("")}
                >
                  Save
                </Button>
              </Grid>
            )}
            {saveToAnExistingCollection && (
              <>
                <Button
                  disabled={collectionName === ""}
                  onClick={() => handleClickSavePost("")}
                >
                  Add
                </Button>
                {allUserCollections !== null &&
                  allUserCollections.map((collection) => (
                    <>
                      <ListItem
                        button
                        onClick={() => setCollectionName(collection)}
                      >
                        <ListItemText primary={collection} />
                      </ListItem>
                      <Divider />
                    </>
                  ))}
              </>
            )}

            {removingFromCollections && (
              <>
                <Button
                  disabled={collectionName === ""}
                  onClick={() => handleClickRemovePostFromCollection("")}
                >
                  Remove
                </Button>
                {collectionsForPost !== null &&
                  collectionsForPost.map((collection) => (
                    <>
                      <ListItem
                        button
                        onClick={() => setCollectionName(collection)}
                      >
                        <ListItemText primary={collection} />
                      </ListItem>
                      <Divider />
                    </>
                  ))}
              </>
            )}
          </List>
        </DialogContent>
      </Dialog>
    </div>
  );
}
