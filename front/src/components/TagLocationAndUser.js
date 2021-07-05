import { Grid, TextField, Avatar } from "@material-ui/core";
import { Autocomplete } from "@material-ui/lab";
import { useState } from "react";

import axios from "axios";

import avatar from "../images/nistagramAvatar.jpg";

const TagLocationAndUser = ({ setLocation, setTaggedUsers }) => {
  const username = localStorage.getItem("username");
  const [searchedContent, setSearchedContent] = useState([]);

  const handleChangeInput = (text) => {
    if (text.length !== 0) {
      axios
        .get("/api/user/search/" + username + "/" + text)
        .then((res) => {
          setSearchedContent(res.data);
        })
        .catch((error) => {
          setSearchedContent([]);
        });
    } else {
      setSearchedContent([]);
    }
  };

  return (
    <div>
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
          <Autocomplete
            freeSolo
            renderOption={(option) => (
              <Grid container>
                <Grid item xs={2}>
                  <Avatar
                    alt="N"
                    src={avatar}
                    style={{ border: "1px solid" }}
                  ></Avatar>
                </Grid>
                <Grid item xs={10} style={{ marginTop: "3%" }}>
                  {option}
                </Grid>
              </Grid>
            )}
            options={
              searchedContent !== null && searchedContent.length !== 0
                ? searchedContent.map((o) => o.Username)
                : []
            }
            //onChange={(event, value) => goToSearchContent(value)}
            renderInput={(params) => (
              <>
                <TextField
                  {...params}
                  variant="outlined"
                  label="Add user"
                  size="small"
                  style={{ width: "100%" }}
                  onChange={(e) => handleChangeInput(e.target.value)}
                ></TextField>
              </>
            )}
          />
        </Grid>
        <Grid item xs={3} />
      </Grid>
    </div>
  );
};

export default TagLocationAndUser;
