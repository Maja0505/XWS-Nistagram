import { Grid, TextField, Avatar, Button } from "@material-ui/core";
import { Autocomplete } from "@material-ui/lab";
import { useState } from "react";

import axios from "axios";

import avatar from "../images/nistagramAvatar.jpg";
import TaggedUsersList from "./TaggedUsersList.js";

const TagLocationAndUser = ({ setLocation, setTaggedUsers, taggedUsers }) => {
  const username = localStorage.getItem("username");
  const [searchedContent, setSearchedContent] = useState([]);
  const [userForTag, setUserForTag] = useState();
  const [open, setOpen] = useState(false);

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

  const addUserInTaggedUsers = () => {
    var array = [...taggedUsers];
    var index = array.indexOf("@" + userForTag);
    if (index === -1) {
      setTaggedUsers((prevState) => [...prevState, "@" + userForTag]);
    }
  };

  const viewAllTaggedUsers = () => {
    console.log(taggedUsers);
    setOpen(true);
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
        <Grid item xs={4}>
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
            onChange={(event, value) => setUserForTag(value)}
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

        <Grid item xs={1}>
          <Button
            disabled={userForTag === null || userForTag === undefined}
            onClick={addUserInTaggedUsers}
          >
            Add
          </Button>
        </Grid>

        <Grid item xs={1}>
          <Button
            disabled={taggedUsers.length === 0}
            onClick={viewAllTaggedUsers}
          >
            View All
          </Button>
        </Grid>
        <Grid item xs={3} />
      </Grid>

      {open && (
        <TaggedUsersList
          label="Tagged users"
          users={taggedUsers}
          open={open}
          setOpen={setOpen}
          setTaggedUsers={setTaggedUsers}
        ></TaggedUsersList>
      )}
    </div>
  );
};

export default TagLocationAndUser;
