import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js"
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";
import { Redirect } from "react-router-dom";
import ProfileDialog from './components/ProfileDialog'

function App() {
  const user = localStorage.getItem("user");

  return (
    <div>
      <Router>
        <div className="App">
          <NavBar></NavBar>
          {user !== null && user !== undefined && <Redirect to="/homePage"/>}
          <Switch>
            <Route exact path="/" component={StartPage}></Route>
            <Route  path="/registration" component={RegistrationPage}></Route>
            <Route  path="/settings" component={Settings}></Route>
            <Route  path="/login" component={LoginPage}></Route>
            <Route  path="/registration" component={RegistrationPage}></Route>
            <Route  path="/dialog" component={ProfileDialog}></Route>
            <Route  path="/homePage" component={UserHomePage}></Route>

          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
