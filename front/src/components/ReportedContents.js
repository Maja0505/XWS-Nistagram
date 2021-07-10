import { Paper, Grid } from "@material-ui/core";

import { useEffect, useState } from "react";

import axios from "axios";

import ReportedContentOne from "./ReportedContentOne";

const ReportedContents = () => {
  const [reportedContents, setReportedContents] = useState([]);

  useEffect(() => {
    axios.get("/api/post/get-all-reported-contents").then((res) => {
      if (res.data !== null) {
        setReportedContents(res.data);
      }
    }).catch((error) => {
      //console.log(error);
    });
  }, []);

  return (
    <div>
      {reportedContents.map((c, index) => (
        <Paper style={{ marginTop: "3%" }} key={index}>
          <Grid container style={{ marginBottom: "2%" }}></Grid>
          <ReportedContentOne
            content={c}
            setReportedContents={setReportedContents}
            reportedContents={reportedContents}
          />
          <Grid container style={{ marginTop: "2%" }}></Grid>
        </Paper>
      ))}
      {reportedContents !== null && reportedContents !== undefined && (
        <>
          {reportedContents.length === 0 && (
            <Paper style={{ marginTop: "3%" }}>
              <Grid container style={{ marginBottom: "2%" }}></Grid>
              <Grid container>
                <p style={{ margin: "auto" }}>NO REPORTED CONTENT</p>
              </Grid>
              <Grid container style={{ marginTop: "2%" }}></Grid>
            </Paper>
          )}
        </>
      )}
    </div>
  );
};

export default ReportedContents;
