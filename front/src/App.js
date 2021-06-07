import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js"
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";
import { Redirect } from "react-router-dom";

function App() {
  const user = localStorage.getItem("user");

  return (
    <div>
      <Router>
        <div className="App">
          <NavBar></NavBar>
          {user !== null && user !== undefined && <Redirect to="/homePage" />}
          <Switch>
            <Route exact path="/" component={StartPage}></Route>
          </Switch>
          <Switch>
            <Route exact path="/" component={LoginPage}></Route>
            <Route exact path="/registration" component={RegistrationPage}></Route>
            <Route exact path="/profile" component={Settings}></Route>
            <Route path="/login" component={LoginPage}></Route>
          </Switch>
          <Switch>
            <Route path="/registration" component={RegistrationPage}></Route>
          </Switch>
          <Switch>
            <Route path="/homePage" component={UserHomePage}></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
