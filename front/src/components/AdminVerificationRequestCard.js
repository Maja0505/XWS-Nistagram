import React, { useEffect, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import { Grid, Paper } from "@material-ui/core";
import axios from "axios";

const useStyles = makeStyles({
  root: {
    width: "100%",
    marginTop: "3%",
  },
});

const AdminVerificationRequestCard = () => {
  const classes = useStyles();
  const [allRequests, setAllRequests] = useState([]);

  const HandleOnClickApprove = (request) => {
    axios
      .put("/api/user/verification-request/approve/" + request.User, {})
      .then((res) => {
        console.log("uspesno");
        var array = [...allRequests]; // make a separate copy of the array
        var index = array.indexOf(request);
        if (index !== -1) {
          array.splice(index, 1);
          setAllRequests(array);
        }
      }).catch((error) => {
        //console.log(error);
      });;
  };

  const HandleOnClickDelete = (request) => {
    axios
      .put("/api/user/verification-request/delete/" + request.User, {})
      .then((res) => {
        console.log("uspesno");
        var array = [...allRequests]; // make a separate copy of the array
        var index = array.indexOf(request);
        if (index !== -1) {
          array.splice(index, 1);
          setAllRequests(array);
        }
      }).catch((error) => {
        //console.log(error);
      });;
  };

  useEffect(() => {
    axios.get("/api/user/verification-request/all").then((res) => {
      if (res.data) {
        console.log(res.data);
        setAllRequests(res.data);
      }
    }).catch((error) => {
      //console.log(error);
    });;
  }, []);

  return (
    <div>
      {allRequests.map((request) => (
        <Card className={classes.root}>
          <Grid container>
            <Grid item xs={6}>
              <CardMedia
                component="img"
                alt="Contemplative Reptile"
                height="200"
                image={
                  "http://localhost:8080/api/media/get-verification-doc/" +
                  request.User +
                  ".jpg"
                }
                title="Contemplative Reptile"
              />
            </Grid>
            <Grid item xs={6}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="h2">
                  USER ID: {request.User}
                </Typography>
                <Typography variant="body2" color="textSecondary" component="p">
                  <Grid container>
                    <Grid item xs={6}>
                      Username:
                    </Grid>
                    <Grid item xs={6} style={{ textAlign: "left" }}>
                      {request.Username}
                    </Grid>
                  </Grid>
                  <Grid container>
                    <Grid item xs={6}>
                      Full name:
                    </Grid>
                    <Grid item xs={6} style={{ textAlign: "left" }}>
                      {request.FullName}
                    </Grid>
                  </Grid>
                  <Grid container>
                    <Grid item xs={6}>
                      KnownAs:
                    </Grid>
                    <Grid item xs={6} style={{ textAlign: "left" }}>
                      {request.KnowAs}
                    </Grid>
                  </Grid>
                  <Grid container>
                    <Grid item xs={6}>
                      Category:
                    </Grid>
                    <Grid item xs={6} style={{ textAlign: "left" }}>
                      {request.Category}
                    </Grid>
                  </Grid>
                </Typography>
              </CardContent>
            </Grid>
          </Grid>
          <Grid container>
            <Grid item xs={6}></Grid>
            <Grid container item xs={6}>
              <Grid item xs={6}></Grid>
              <Grid item xs={6}>
                <CardActions>
                  <Button
                    size="small"
                    color="primary"
                    onClick={() => HandleOnClickApprove(request)}
                  >
                    Approve
                  </Button>
                  <Button
                    size="small"
                    color="primary"
                    onClick={() => HandleOnClickDelete(request)}
                  >
                    Delete
                  </Button>
                </CardActions>
              </Grid>
            </Grid>
          </Grid>
        </Card>
      ))}

      {allRequests !== null && allRequests !== undefined && (
        <>
          {allRequests.length === 0 && (
            <Paper style={{ marginTop: "3%" }}>
              <Grid container style={{ marginBottom: "2%" }}></Grid>
              <Grid container>
                <p style={{ margin: "auto" }}>NO VERIFICATION REQUESTS</p>
              </Grid>
              <Grid container style={{ marginTop: "2%" }}></Grid>
            </Paper>
          )}
        </>
      )}
    </div>
  );
};

export default AdminVerificationRequestCard;
