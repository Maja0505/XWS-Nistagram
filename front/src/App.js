import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js";
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";
import PostDialog from "./components/PostDialog";
import StoryBar from "./components/StoryBar";
import ContentDetails from "./components/ContentDetails";
import AdminHomePage from "./components/AdminHomePage";
import LikedDislikedPost from "./components/LikedDislikedPost";
import HashTagPost from "./components/HashTagPost";
import PostsForCollection from "./components/PostsForCollection";
import { useEffect } from "react";

function App() {
  const logedUsername = localStorage.getItem("username");


  useEffect(() => {
    let userid = localStorage.getItem("id")
    let connected = localStorage.getItem("connected")
    if(userid){
      if(!connected){
        let socket = new WebSocket("ws://localhost:8080/api/notification/chat/" + userid)
        socket.onopen = () => {
          console.log("Successfully Connected");
          socket.send('{"command": 0, "channel": ' + '"' + userid + '"' + '}')
          localStorage.setItem('connected',true)
        };
  
        socket.onclose = event => {
          console.log("Socket Closed Connection: ", event);
          localStorage.setItem('connected',false)
      }

      
      socket.onmessage = event => {
       alert(`${event.data}`);
        };
      }


  
  };



  }, [])


  const users = [
    {
      Username: "Perica",
      FirstName: "Perica",
      LastName: "Peric",
      DateOfBirth: "krdlkjf",
      Email: "Peric.peric@gmail.com",
      PhoneNumber: "0490843",
      Gender: "Female",
      Biography: "Jedna vrlo uspesan gospodin",
      WebSite: "Pericaperic.com",
    },

    {
      Username: "marko",
      FirstName: "Marko",
      LastName: "Markovic",
      DateOfBirth: "krdlkjf",
      Email: "marko.markovic@gmail.com",
      PhoneNumber: "0490843",
      Gender: "Male",
      Biography: "Jedna vrlo uspesan gospodin",
      WebSite: "Pericaperic.com",
    },
  ];

  return (
    <div>
      <Router>
        <div className="App">
          <NavBar></NavBar>
          <Switch>
            <Route exact path="/" component={StartPage}></Route>
            <Route
              exact
              path="/registration"
              component={RegistrationPage}
            ></Route>
            <Route exact path="/accounts/*" component={Settings}></Route>
            <Route path="/login" component={LoginPage}></Route>
            <Route path="/registration" component={RegistrationPage}></Route>
            <Route
              exact
              path="/homePage/:username"
              render={(props) => <UserHomePage {...props} />}
            ></Route>

            <Route
              path="/homePage/:username/collection/:collection"
              render={(props) => <PostsForCollection {...props} />}
            ></Route>

            <Route path="/admin" component={AdminHomePage}></Route>

            <Route
              path="/dialog/:username/:post"
              render={(props) => <PostDialog {...props} />}
            ></Route>

            <Route
              exact
              path="/:username/liked-disliked/"
              render={(props) => <LikedDislikedPost {...props} />}
            ></Route>
            <Route
              exact
              path="/explore/tags/:tag/"
              render={(props) => <HashTagPost {...props} />}
            ></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
