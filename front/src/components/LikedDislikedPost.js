import { useParams } from "react-router-dom";

const LikedDislikedPost = () => {
  const { username } = useParams();

  const loggedUsername = localStorage.getItem("username");

  return <div style={{ marginTop: "10%" }}>CAO CAOO {username}</div>;
};

export default LikedDislikedPost;
