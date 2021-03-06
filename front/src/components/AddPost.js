import { Button, Grid, TextField } from "@material-ui/core";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";
import Slider from "react-slick";
import { makeStyles } from "@material-ui/core/styles";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

import TagLocationAndUser from "./TagLocationAndUser.js";

const useStyles = makeStyles((theme) => ({
  settings: {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  },
  rtl: {
    rtl: true,
  },
}));

const AddPost = ({ setTabValue }) => {
  const classes = useStyles();

  const [selectedFile, setSelectedFile] = useState([]);
  const [image, setImage] = useState([]);
  const [description, setDescription] = useState("");
  const loggedUserId = localStorage.getItem("id");
  const loggedUsername = localStorage.getItem("username");

  const [isVideo, setIsVideo] = useState([]);
  const [imagesIdsForSave, setImagesIdsForSave] = useState([]);
  const [puklaSlika, setPuklaSlika] = useState(false);
  const [listOfTaggedUserid,setListOfTaggedUserid] = useState([])

  const [location, setLocation] = useState("");
  const [taggedUsers, setTaggedUsers] = useState("");

  const createPost = () => {
    console.log(image);

    for (let index = 0; index < image.length; index++) {
      if (!isVideo[index]) {
        uploadImage(image[index], index);
      } else {
        uploadVideo(image[index], index);
      }
    }
    savePost();
  };

  const addTags = (postDTO) => {
    var listOfTags = description.split("#");
    if (listOfTags.length > 0) {
      for (var i = 1; i < listOfTags.length; i++) {
        let tag = listOfTags[i].split(" ")[0];
        axios
          .post("/api/post/add-tag", {
            Tag: "#" + tag,
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
    for (var j = 0; j < taggedUsers.length; j++) {
      let userTag = taggedUsers[j];
      let userid = listOfTaggedUserid[j]
      axios
        .post("/api/post/add-tag", {
          Tag: userTag,
          PostID: postDTO.ID,
        })
        .then((res) => {
          let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + loggedUserId)
          socket.onopen = () => {
            console.log("Successfully Connected");
            console.log("aa")
            socket.send('{"user_who_follow":' + '"' + loggedUsername + '"' + ',"command": 2, "channel": ' + '"' + userid + '"' + ', "content": "tagged you in a post."' + ', "media": "' + postDTO.Media[0] + '"' + ', "post_id": "' + postDTO.ID + '"}')
          };
          console.log("Upisan user tag  " + userTag);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  };

  const uploadVideo = (imageForUpload, index) => {
    var imageId = uuidv4().toString() + "A" + loggedUserId.toString() + ".mp4";
    var array = imagesIdsForSave;
    array.push(imageId);
    setImagesIdsForSave(array);

    axios
      .post(
        "/api/media/upload-video/" +
          imageId.substring(0, imageId.length - 4) +
          "/" +
          "image" +
          index,
        imageForUpload,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {})
      .catch((error) => {
        alert(error);
        setPuklaSlika(true);
      });
  };

  const uploadImage = (imageForUpload, index) => {
    var imageId = uuidv4().toString() + "A" + loggedUserId.toString() + ".jpg";
    var array = imagesIdsForSave;
    array.push(imageId);
    setImagesIdsForSave(array);
    axios
      .post(
        "/api/media/upload-media-image/" +
          imageId.substring(0, imageId.length - 4) +
          "/" +
          "image" +
          index,
        imageForUpload,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      )
      .then((res) => {})
      .catch((error) => {
        //alert(error);
        setPuklaSlika(true);
      });
  };

  const savePost = () => {
    console.log(puklaSlika);
    console.log(imagesIdsForSave);
    if (!puklaSlika) {
      var postDTO = {
        Description: description,
        Media: imagesIdsForSave,
        MediaCount: imagesIdsForSave.length,
        Album: imagesIdsForSave.length === 0 ? false : true,
        UserID: loggedUserId,
        Location: location,
      };
      console.log("Uspesno upload-ovao sliku");
      axios
        .post("/api/post/create", postDTO)
        .then((res1) => {
          console.log("Uspesno kreirao post");
          var postDTONew = { ...postDTO, ID: res1.data };
          addTags(postDTONew);
          setTabValue(0);
          setPuklaSlika(false);
        })
        .catch((error) => {
          console.log(error);
          setPuklaSlika(false);
        });
    }
  };

  const HandleUploadMedia = (event) => {
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
      <Grid container style={{ marginTop: "2%" }}>
        <Grid item xs={2} />
        <Grid item xs={4}>
          <Button
            variant="contained"
            component="label"
            style={{ margin: "auto" }}
          >
            {selectedFile.length === 0 ? `Choose media` : `Change media`}
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
          {selectedFile.length !== 0 && (
            <Button
              variant="contained"
              color="primary"
              onClick={createPost}
              style={{ margin: "auto" }}
              disabled={image.length === 0}
            >
              Create post
            </Button>
          )}
        </Grid>
        <Grid item xs={2} />
      </Grid>

      {selectedFile.length !== 0 && (
        <>
          <TagLocationAndUser
            setLocation={setLocation}
            setTaggedUsers={setTaggedUsers}
            taggedUsers={taggedUsers}
            setListOfTaggedUserid = {setListOfTaggedUserid}
            listOfTaggedUserid = {listOfTaggedUserid}
          />

          <Grid container style={{ marginTop: "1%" }}>
            <Grid item xs={3} />
            <Grid item xs={6}>
              <TextField
                label="Description"
                fullWidth
                size="small"
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
              <div>
                <Slider {...settings}>
                  {selectedFile.map((media, index) => (
                    <div>
                      {!isVideo[index] && (
                        <img
                          width="100%"
                          height="500"
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
              </div>
            </Grid>
            <Grid item xs={2} />
          </Grid>
        </>
      )}
    </div>
  );
};

export default AddPost;
