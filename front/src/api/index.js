// api/index.js
var userid = localStorage.getItem("id");
var socket = new WebSocket("ws://localhost:8080/ws/msg/connect-ws/" + userid);


let connect = cb => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
    sendMsg('{"command": 0, "channel": '+ '"' + userid + '"' + '}');
  };

  socket.onmessage = msg => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
