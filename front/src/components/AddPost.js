import { Button, Grid, TextField } from "@material-ui/core";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";

const AddPost = ({ setTabValue }) => {
  const [selectedFile, setSelectedFile] = useState();
  const [image, setImage] = useState();
  const [description, setDescription] = useState("");
  const loggedUserId = localStorage.getItem("id");
  const [isVideo, setIsVideo] = useState(false);

  const createPost = async () => {
    if (!isVideo) {
      uploadImage();
    } else {
      uploadVideo();
    }
  };

  const addTags = (postDTO) => {
    var listOfTags = description.split("#");
    if (listOfTags.length > 0) {
      for (var i = 1; i < listOfTags.length; i++) {
        let tag = listOfTags[i].split(" ")[0];
        axios
          .post("/api/post/add-tag", {
            Tag: tag,
            PostID: postDTO.ID,
          })
          .then((res) => {
            console.log("Upisan tag  " + tag);
          })
          .catch((error) => {
            console.log(error);
          });
      }
    }
  };

  const uploadVideo = () => {
    var postDTO = {
      ID: uuidv4(),
      Description: description,
      Image: uuidv4().toString() + "A" + loggedUserId.toString() + ".mp4",
      UserID: loggedUserId,
    };

    axios
      .post(
        "/api/post/video-upload/" +
          postDTO.Image.substring(0, postDTO.Image.length - 4),
        image,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {
        console.log("Uspesno upload-ovao video");
        axios
          .post("/api/post/create", postDTO)
          .then((res1) => {
            console.log("Uspesno kreirao post");
            console.log(postDTO);
            addTags(postDTO);
            setTabValue(0);
          })
          .catch((error) => {
            console.log(error);
          });
      })
      .catch((error) => {
        alert(error);
      });
  };

  const uploadImage = () => {
    var postDTO = {
      ID: uuidv4(),
      Description: description,
      Image: uuidv4().toString() + "A" + loggedUserId.toString() + ".jpg",
      UserID: loggedUserId,
    };

    axios
      .post(
        "/api/post/upload-image/" +
          postDTO.Image.substring(0, postDTO.Image.length - 4),
        image,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {
        console.log("Uspesno upload-ovao sliku");
        axios
          .post("/api/post/create", postDTO)
          .then((res1) => {
            console.log("Uspesno kreirao post");
            console.log(postDTO);
            addTags(postDTO);
            setTabValue(0);
          })
          .catch((error) => {
            console.log(error);
          });
      })
      .catch((error) => {
        alert(error);
      });
  };

  const HandleUploadMedia = (event) => {
    setSelectedFile(null);

    var formData = new FormData();
    console.log(event.target.files[0]);
    var file = event.target.files[0];
    if (event.target.files[0].type === "video/mp4") {
      setIsVideo(true);
    } else {
      setIsVideo(false);
    }
    formData.append("myFile", file);
    const reader = new FileReader();
    var url = reader.readAsDataURL(file);
    reader.onloadend = function (e) {
      setSelectedFile(reader.result);
    }.bind(this);

    setImage(formData);
  };

  return (
    <div>
      <Grid container style={{ marginTop: "2%" }}>
        <Grid item xs={2} />
        <Grid item xs={4}>
          <Button
            variant="contained"
            component="label"
            style={{ margin: "auto" }}
          >
            {!selectedFile ? `Choose media` : `Change media`}
            <input
              hidden
              accept="image/*,video/mp4,video/x-m4v,video/*"
              multiple
              type="file"
              name="myFile"
              onChange={(event) => HandleUploadMedia(event)}
            />
          </Button>
        </Grid>
        <Grid item xs={4}>
          {selectedFile && (
            <Button
              variant="contained"
              color="primary"
              onClick={createPost}
              style={{ margin: "auto" }}
            >
              Create post
            </Button>
          )}
        </Grid>
        <Grid item xs={2} />
      </Grid>
      {selectedFile && (
        <>
          <Grid container style={{ marginTop: "2%" }}>
            <Grid item xs={3} />
            <Grid item xs={6}>
              <TextField
                label="Description"
                fullWidth
                variant="outlined"
                multiline
                rowsMax={5}
                onChange={(e) => setDescription(e.target.value)}
              ></TextField>
            </Grid>
            <Grid item xs={3} />
          </Grid>

          <Grid container style={{ marginTop: "3%" }}>
            <Grid item xs={2} />
            <Grid item xs={8}>
              {!isVideo && <img width="100%" src={selectedFile} />}
              {isVideo && (
                <video width="100%" controls>
                  <source src={selectedFile} type="video/mp4" />
                </video>
              )}
            </Grid>
            <Grid item xs={2} />
          </Grid>
        </>
      )}
    </div>
  );
};

export default AddPost;
