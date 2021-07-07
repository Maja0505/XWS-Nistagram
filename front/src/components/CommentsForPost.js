import CommentOneForFeed from "./CommentOneForFeed";
const CommentsForPost = ({ comments }) => {
  return (
    <div style={{ overflow: "auto", width: "100%" }}>
      {comments !== null &&
        comments !== undefined &&
        comments.map((c, index) => (
          <div key={index} style={{ textAlign: "left", marginLeft: "7%" }}>
            <CommentOneForFeed comment={c} />
          </div>
        ))}
    </div>
  );
};

export default CommentsForPost;
