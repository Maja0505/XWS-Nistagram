import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js";
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";
import PostDialog from "./components/PostDialog";
import AdminHomePage from "./components/AdminHomePage";
import LikedDislikedPost from "./components/LikedDislikedPost";
import HashTagPost from "./components/HashTagPost";
import PostsForCollection from "./components/PostsForCollection";
import LocationTagPost from "./components/LocationTagPost.js";
import FollowSuggestions from "./components/FollowSuggestions";

function App() {
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
              path="/dialog/:post"
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
            <Route
              exact
              path="/explore/locations/:location/"
              render={(props) => <LocationTagPost {...props} />}
            ></Route>

            <Route
              path="/follow-suggestions/"
              component={FollowSuggestions}
            ></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
