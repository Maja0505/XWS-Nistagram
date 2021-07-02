import React,{useState} from "react";
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
import { makeStyles } from '@material-ui/core/styles';
import uuid from 'react-uuid'
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
      height:800,
      backgroundColor: theme.palette.background.paper,
    },
  }));

  
const AddStoryDialog = ({open,setOpen}) => {

    const classes = useStyles();
    const [inappropriate,setInappropriate] = useState(false)
    const [selectedFile, setSelectedFile] = useState();
    const [image, setImage] = useState();
    const [close,setClose] = useState(false);
    const [highlights,setHighLights] = useState(false)
    const loggedUserId = localStorage.getItem("id");

    
    const HandleOnChangeCloseFriends = () => {
        if(close){
            setClose(false)
        }else{
            setClose(true)
        }
    }

    const HandleOnChangeHighlights = () => {
        if(highlights){
            setHighLights(false)
        }else{
            setHighLights(true)
        }
    }

    const handleClose = () => {
        setSelectedFile()
        setImage()
      setOpen(false);
    };

    const HandleClickOnSend = () => {
        var story = {
            UserID: loggedUserId,
            Image: "" + loggedUserId + "-" + uuid(),
            Highlights: highlights,
            ForCloseFriends: close,
        };

        axios.post("/api/post/story/image-upload/" + story.Image)
            .then((res) => 
            {
                axios.post("/api/post/story/create", story)
                    .then((res)=> 
                    {
                        console.log("uspesno")
                    })
            })
 
      };
    

    const HandleUploadClick = (event) => {
        var formData = new FormData();
        console.log(event.target.files[0]);
        var file = event.target.files[0];
        formData.append("myFile", file);
        const reader = new FileReader();
        var url = reader.readAsDataURL(file);
        reader.onloadend = function (e) {
          setSelectedFile(reader.result);
        }.bind(this);
    
        setImage(formData);
      };
    
    

    const DialogContent = withStyles((theme) => ({
        root: {
            width: 500,
            height:500,
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
               <Grid container style={{height:"10%"}}>
                   <Grid item xs={8}></Grid>
                   <Grid item xs={4}>
                   <Button
                          variant="contained"
                          component="label"
                        >
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
               <Grid container style={{height:"2%"}}></Grid>

               <Grid container style={{margin:'auto'}}>
               <Grid item xs={1}></Grid>

                   <Grid item xs={10}> 
                   <img width="100%"  src={selectedFile} />
                    </Grid>
                   <Grid item xs={1}></Grid>

               </Grid>
               <Divider/>
               <Grid container style={{height:"10%"}} >
                   <Grid item xs={5}>
                   <FormGroup>
                        <FormControlLabel
                            control={<Checkbox  checked={close === true} onChange={HandleOnChangeCloseFriends} />}
                            label="Close friends"
                            style={{ fontSize: 15, fontWeight: "bold" }}
                        />
                        </FormGroup>
                   </Grid>
                   <Grid item xs={5}>
                        <FormGroup >
                        <FormControlLabel
                            control={<Checkbox checked={highlights === true} onChange={HandleOnChangeHighlights} />}
                            label="Highlights"
                            style={{ fontSize: 15, fontWeight: "bold" }}
                        />
                        </FormGroup>

                   </Grid>
                   <Grid item xs={2}><Button style={{alignItems:"end"}} variant="contained" onClick={HandleClickOnSend}>Add</Button></Grid>

               </Grid>
          </DialogContent>
        </Dialog>
      </div>
    )
}

export default AddStoryDialog
