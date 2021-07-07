import { useEffect, useState } from "react";

import axios from "axios";

const ReportedContents = () => {
  const [reportedContents, setReportedContents] = useState([]);

  useEffect(() => {
    axios.get("/api/post/get-all-reported-contents").then((res) => {
      if (res.data !== null) {
        console.log(res.data);
        setReportedContents(res.data);
      }
    });
  }, []);

  return <div></div>;
};

export default ReportedContents;
