import CommentOneForFeed from "./CommentOneForFeed";

const CommentsForFeeds = ({ comments }) => {
  return (
    <div style={{ overflow: "auto", height: "100px" }}>
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

export default CommentsForFeeds;
