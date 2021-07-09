import { Button, FormLabel, Grid, TextField } from "@material-ui/core";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";
import Slider from "react-slick";
import { makeStyles } from "@material-ui/core/styles";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { DateTimePicker } from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import * as moment from "moment";
import { MuiPickersUtilsProvider } from "@material-ui/pickers";

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

const AddCampaign = ({ setTabValue }) => {
  const classes = useStyles();

  const [selectedFile, setSelectedFile] = useState([]);
  const [image, setImage] = useState([]);
  const [description, setDescription] = useState("");
  const loggedUserId = localStorage.getItem("id");
  const loggedUsername = localStorage.getItem("username");

  const [isVideo, setIsVideo] = useState([]);
  const [imagesIdsForSave, setImagesIdsForSave] = useState([]);
  const [puklaSlika, setPuklaSlika] = useState(false);
  const [listOfTaggedUserid, setListOfTaggedUserid] = useState([]);
  const [links, setLinks] = useState([]);
  const [repeat, setRepeat] = useState("");
  const [valueOneStart, onChangeOneStart] = useState(null);
  const [valueMultieStart, onChangeMultieStart] = useState(null);
  const [valueMultieEnd, onChangeMultieEnd] = useState(null);
  const [repeatFactor, setRepeatFactor] = useState(0);

  const [location, setLocation] = useState("");
  const [taggedUsers, setTaggedUsers] = useState("");
  const [clearedDate, handleClearedDateChange] = useState(null);
  const [tags, setTags] = useState([]);

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

  const addTags = () => {
    var listOfTags = description.split("#");
    if (listOfTags.length > 0) {
      for (var i = 1; i < listOfTags.length; i++) {
        let tag = listOfTags[i].split(" ")[0];
        var array = tags;
        array.push("#" + tag);
        setTags(array);
      }
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
        alert(error);
        setPuklaSlika(true);
      });
  };

  const savePost = () => {
    console.log(puklaSlika);
    console.log(imagesIdsForSave);
    var end = valueOneStart;
    var start = valueOneStart;
    var d1 = new Date(valueOneStart);
    var d2 = new Date(valueMultieEnd);
    var d3 = new Date(valueMultieStart);

    const valueOneStart2 = moment(d1).format("YYYY-MM-DD HH:mm:ss");
    const valueMultieEnd2 = moment(d2).format("YYYY-MM-DD HH:mm:ss");
    const valueMultieStart2 = moment(d3).format("YYYY-MM-DD HH:mm:ss");
    console.log(valueMultieEnd2);
    console.log(valueMultieStart2);

    if (repeat === "multiple-time") {
      var endArray = valueMultieEnd2.split(" ");
      end = endArray[0] + "T" + endArray[1] + ".141Z";
      var startMArray = valueMultieStart2.split(" ");
      start = startMArray[0] + "T" + startMArray[1] + ".141Z";
      console.log(endArray);
      console.log(startMArray);
    } else {
      var startOArray = valueOneStart2.split(" ");
      start = startOArray[0] + "T" + startOArray[1] + ".141Z";
      end = startOArray[0] + "T" + startOArray[1] + ".141Z";
      console.log(startOArray);
    }

    addTags();

    if (!puklaSlika) {
      var postDTO = {
        description: description,
        media: imagesIdsForSave,
        end: end,
        start: start,
        links: links,
        tags: tags,
        ispost: true,
        userid: loggedUserId,
        location: location,
        repeat: repeat === "multiple-time" ? true : false,
        repeatfactor: Number(repeatFactor),
        //Influencers:[]
      };
      console.log("Uspesno upload-ovao sliku");
      axios
        .post("/api/agent/create-campaign", postDTO)
        .then((res1) => {
          console.log("Uspesno kreirao post");
          //var postDTONew = { ...postDTO, ID: res1.data };
          //addTags(postDTONew);
          setTabValue(0);
          setPuklaSlika(false);
        })
        .catch((error) => {
          console.log(error);
          setPuklaSlika(false);
        });
    }
  };

  const HandleChangeLink = (index, value) => {
    let newArr = [...links]; // copying the old datas array
    newArr[index] = value; // replace e.target.value with whatever you want to change it to
    setLinks(newArr);
    console.log(links);
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
    for (let index = 0; index < event.target.files.length; index++) {
      var array = links;
      array.push("");
      setLinks(array);
    }
    console.log(links);
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
          <Grid container style={{ marginTop: "1%" }}>
            <Grid item xs={3} />
            <Grid item xs={6}>
              <TextField
                label="Add location"
                fullWidth
                variant="outlined"
                size="small"
                onChange={(e) => setLocation(e.target.value)}
              ></TextField>
            </Grid>
            <Grid item xs={3} />
          </Grid>

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

          <Grid container style={{ marginTop: "1%" }}>
            <Grid item xs={3} />

            <Grid item xs={3}>
              <Button
                variant="contained"
                color="primary"
                marginLeft={10}
                onClick={() => setRepeat("one-time")}
                style={{ margin: "auto", width: "100%" }}
              >
                one-time
              </Button>
            </Grid>
            <Grid item xs={3}>
              <Button
                variant="contained"
                color="primary"
                style={{ margin: "auto", width: "100%" }}
                onClick={() => setRepeat("multiple-time")}
              >
                multiple-time
              </Button>
            </Grid>
            <Grid item xs={3} />
          </Grid>
          {repeat === "multiple-time" && (
            <>
              <Grid container style={{ marginTop: "1%" }}>
                <Grid item xs={3} />

                <Grid item xs={3}>
                  <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <DateTimePicker
                      autoOk
                      ampm={false}
                      value={valueMultieStart}
                      onChange={onChangeMultieStart}
                      format="yyyy-MM-dd hh:mm"
                      label="Start"
                    />
                  </MuiPickersUtilsProvider>
                </Grid>
                <Grid item xs={3}>
                  <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <DateTimePicker
                      autoOk
                      ampm={false}
                      value={valueMultieEnd}
                      onChange={onChangeMultieEnd}
                      format="yyyy-MM-dd hh:mm"
                      label=" End"
                    />
                  </MuiPickersUtilsProvider>
                </Grid>
                <Grid item xs={3} />
              </Grid>
              <Grid container style={{ marginTop: "1%" }}>
                <Grid item xs={3}></Grid>
                <Grid item xs={6}>
                  <Grid item xs={6}>
                    <TextField
                      id="outlined-number"
                      label="Repeat in day"
                      type="number"
                      textAlign="left"
                      style={{ textAlign: "left" }}
                      value={repeatFactor}
                      onChange={(event) => setRepeatFactor(event.target.value)}
                      InputLabelProps={{
                        shrink: true,
                      }}
                      variant="outlined"
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                </Grid>
                <Grid item xs={3}></Grid>
              </Grid>
            </>
          )}

          {repeat === "one-time" && (
            <Grid container style={{ marginTop: "1%" }}>
              <Grid item xs={3} />
              <Grid item xs={3}>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                  <DateTimePicker
                    autoOk
                    ampm={false}
                    value={valueOneStart}
                    onChange={onChangeOneStart}
                    format="yyyy-MM-dd hh:mm"
                    label="Start"
                  />
                </MuiPickersUtilsProvider>
              </Grid>
              <Grid item xs={3} />
            </Grid>
          )}
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
                      <TextField
                        label="Link to web site of this product:"
                        fullWidth
                        size="small"
                        variant="outlined"
                        rowsMax={1}
                        onChange={(e) =>
                          HandleChangeLink(index, e.target.value)
                        }
                      ></TextField>
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

export default AddCampaign;
