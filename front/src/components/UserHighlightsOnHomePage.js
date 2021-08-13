import { useState, useEffect } from "react";
import { Grid } from "@material-ui/core";

import Story from "./Story";

import highlights from "../images/highlights.jpg";

const UserHighlightsOnHomePage = ({ highlightStories, userId }) => {
  const [openHighlightsDialog, setOpenHighlightsDialog] = useState(false);

  function closeStory() {
    setOpenHighlightsDialog(false);
  }

  const handleClickOpen = () => {
    setOpenHighlightsDialog(true);
  };

  return (
    <div>
      {highlightStories !== null && (
        <>
          <Grid container style={{ marginTop: "1%" }}></Grid>
          <Grid container style={{ margin: "auto" }}>
            <Grid item xs={2}></Grid>
            <Grid container item xs={8}>
              <Grid item xs={4}>
                <div onClick={handleClickOpen}>
                  <div>
                    <img
                      src={highlights}
                      style={{
                        borderRadius: "50%",
                        border: "1px solid",
                        width: "50px",
                        height: "50px",
                        cursor: "pointer",
                      }}
                    />
                  </div>
                  <div>{"Highlights"}</div>
                </div>
              </Grid>
              <Grid item xs={7}></Grid>
            </Grid>
            <Grid item xs={2}></Grid>
          </Grid>
        </>
      )}

      {openHighlightsDialog && highlightStories !== null && (
        <Story
          stories={highlightStories}
          onClose={closeStory}
          user={userId}
        ></Story>
      )}
    </div>
  );
};

export default UserHighlightsOnHomePage;
