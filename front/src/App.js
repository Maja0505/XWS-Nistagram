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

function App() {
  const logedUsername = localStorage.getItem("username");

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
          <ContentDetails></ContentDetails>
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
              path="/homePage/:username"
              render={(props) => <UserHomePage {...props} />}
            ></Route>
            <Route path="/admin" component={AdminHomePage}></Route>

            <Route
              path="/dialog/:post"
              render={(props) => <PostDialog {...props} />}
            ></Route>
            <Route
              path="/dialog/:post"
              render={(props) => <PostDialog {...props} />}
            >
              <PostDialog></PostDialog>
            </Route>
            <Route
              exact
              path="/:username/liked-disliked/"
              render={(props) => <LikedDislikedPost {...props} />}
            ></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
