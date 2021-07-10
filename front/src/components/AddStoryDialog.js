import React, { useState } from "react";
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
import Checkbox from "@material-ui/core/Checkbox";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import { makeStyles } from "@material-ui/core/styles";
import uuid from "react-uuid";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import Slider from "react-slick";

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
    width: 800,
    height: 800,
    backgroundColor: theme.palette.background.paper,
  },
}));

const AddStoryDialog = ({ open, setOpen,setHaveStory }) => {
  const classes = useStyles();
  const [inappropriate, setInappropriate] = useState(false);
  const [selectedFile, setSelectedFile] = useState([]);
  const [image, setImage] = useState([]);
  const [close, setClose] = useState(false);
  const [highlights, setHighLights] = useState(false);
  const loggedUserId = localStorage.getItem("id");
  const [isVideo, setIsVideo] = useState([]);

  const HandleOnChangeCloseFriends = () => {
    if (close) {
      setClose(false);
    } else {
      setClose(true);
    }
  };

  const HandleOnChangeHighlights = () => {
    if (highlights) {
      setHighLights(false);
    } else {
      setHighLights(true);
    }
  };

  const handleClose = () => {
    setSelectedFile();
    setImage();
    setOpen(false);
  };

  const createPost = () => {
    console.log(image);

    for (let index = 0; index < image.length; index++) {
      
      HandleClickOnSend(image[index], index);
      
    }

    setOpen(false);
  };

  const HandleClickOnSend = (imageForUpload, index) => {
    if (!isVideo[index]) {
      var imageString = "" + loggedUserId + "-" + uuid();
     

      axios
        .post("/api/media/upload-media-image/" + imageString + "/" + "image" + index, imageForUpload, {
          headers: { "Content-Type": "multipart/form-data" },
        })

        .then((res) => {
          var story = {
            UserID: loggedUserId,
            Image: imageString + ".jpg",
            Highlights: highlights,
            ForCloseFriends: close,
            Link:''
          };
          axios.post("/api/post/story/create", story).then((res) => {
            console.log("uspesno");
            //setOpen(false);
            setHaveStory(true)
          });
        }).catch((error) => {
          alert(error)
        });;
    } else {
      var imageString = "" + loggedUserId + "-" + uuid();
     

      axios
        .post(
          "/api/media/upload-video/"  + imageString + "/" + "image" + index,
          imageForUpload,

          {
            headers: { "Content-Type": "multipart/form-data" },
          }
        )
        .then((res) => {
          var story = {
            UserID: loggedUserId,
            Image: imageString + ".mp4",
            Highlights: highlights,
            ForCloseFriends: close,
            Link:''
          };
          axios.post("/api/post/story/create", story).then((res) => {
            console.log("uspesno");
            //setOpen(false);
            setHaveStory(true)
          });
        }).catch((error) => {
          alert(error)
        });
    }
  };

  const HandleUploadClick = (event) => {
    setSelectedFile([]);
    setIsVideo([]);
    setImage([]);

    var formData = new FormData();
    for (let index = 0; index < event.target.files.length; index++) {
      if (event.target.files[index].type === "video/mp4") {
        var array = isVideo;
        array.push(true);
        setIsVideo(array);
      } else {
        var array = isVideo;
        array.push(false);
        setIsVideo(array);
      }

      var file = event.target.files[index];
      formData.append("image" + index, file);
      const reader = new FileReader();
      var url = reader.readAsDataURL(file);
      reader.onloadend = function (e) {
        setSelectedFile((prevState) => [...prevState, reader.result]);
      }.bind(this);
      setImage((prevState) => [...prevState, formData]);
      console.log(formData);
    }
    console.log(isVideo);
  };

  const DialogContent = withStyles((theme) => ({
    root: {
      width: 500,
      height: 500,
    },
  }))(MuiDialogContent);

  const DialogActions = withStyles((theme) => ({
    root: {
      margin: 0,
      padding: theme.spacing(1),
    },
  }))(MuiDialogActions);

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


  function SampleNextArrow(props) {
    const { className, style, onClick } = props;
    return (
      <div
        className={className}
        style={{ ...style, zIndex: 1, right: 0, width: 30, height: 30 }}
        onClick={onClick}
      />
    );
  }

  function SamplePrevArrow(props) {
    const { className, style, onClick } = props;
    return (
      <div
        className={className}
        style={{ ...style, zIndex: 1, left: 0 }}
        onClick={onClick}
      />
    );
  }

  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    nextArrow: <SampleNextArrow />,
    prevArrow: <SamplePrevArrow />,
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
          Add Story Form
        </DialogTitle>
        <DialogContent dividers>
          <Grid container style={{ height: "10%" }}>
            <Grid item xs={8}></Grid>
            <Grid item xs={4}>
              <Button variant="contained" component="label">
                Choose file
                <input
                  hidden
                  accept="image/*"
                  className={classes.input}
                  multiple
                  type="file"
                  name="myFile"
                  onChange={(event) => HandleUploadClick(event)}
                />
              </Button>
            </Grid>
          </Grid>
          <Grid container style={{ height: "2%" }}></Grid>

          <Grid container style={{ margin: "auto" }}>
            <Grid item xs={1}></Grid>

            <Grid item xs={10}>
            {selectedFile &&
            <div>
                <Slider {...settings}>
                  {selectedFile.map((media, index) => (
                    <div>
                      {!isVideo[index] && (
                        <img
                          width="100%"
                          height="500px"
                          src={selectedFile[index]}
                        />
                      )}
                      {isVideo[index] && (
                        <video width="100%" controls>
                          <source src={selectedFile[index]} type="video/mp4" />
                        </video>
                      )}
                    </div>
                  ))}
                </Slider>
              </div>}
            </Grid>
            <Grid item xs={1}></Grid>
          </Grid>
          <Divider />
          {selectedFile !== undefined && selectedFile !== null && selectedFile.length !== 0 &&
          <Grid container style={{ height: "10%" }}>
            <Grid item xs={5}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Checkbox
                      checked={close === true}
                      onChange={HandleOnChangeCloseFriends}
                    />
                  }
                  label="Close friends"
                  style={{ fontSize: 15, fontWeight: "bold" }}
                />
              </FormGroup>
            </Grid>
            <Grid item xs={5}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Checkbox
                      checked={highlights === true}
                      onChange={HandleOnChangeHighlights}
                    />
                  }
                  label="Highlights"
                  style={{ fontSize: 15, fontWeight: "bold" }}
                />
              </FormGroup>
            </Grid>
            <Grid item xs={2}>
              <Button
                style={{ alignItems: "end" }}
                variant="contained"
                onClick={createPost}
                
              >
                Add
              </Button>
            </Grid>
          </Grid>}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default AddStoryDialog;
