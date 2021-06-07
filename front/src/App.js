import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import UserHomePage from "./components/UserHomePage.js";
import StartPage from "./components/StartPage.js";
import NavBar from "./components/NavBar";

function App() {
  return (
    <div>
      <Router>
        <div className="App">
          <NavBar></NavBar>
          <Switch>
            <Route exact path="/" component={StartPage}></Route>
          </Switch>
          <Switch>
            <Route path="/login" component={LoginPage}></Route>
          </Switch>
          <Switch>
            <Route path="/registration" component={RegistrationPage}></Route>
          </Switch>
          <Switch>
            <Route
              path="/homePage/:username"
              render={(props) => <UserHomePage {...props} />}
            ></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
