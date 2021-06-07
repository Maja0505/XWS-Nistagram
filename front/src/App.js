import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js";
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";
import ProfileDialog from "./components/ProfileDialog";

function App() {
  const logedUsername = localStorage.getItem("username");

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
            <Route exact path="/settings" component={Settings}></Route>
            <Route path="/login" component={LoginPage}></Route>
            <Route path="/registration" component={RegistrationPage}></Route>
            <Route
              path="/homePage/:username"
              render={(props) => <UserHomePage {...props} />}
            ></Route>
            <Route path="/dialog">
              <ProfileDialog openD={true}></ProfileDialog>
            </Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
