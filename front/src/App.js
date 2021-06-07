import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";
import Settings from "./components/Settings.js"

function App() {
  return (
    <div>
      <Router>
        <div className="App">
          <Switch>
            <Route exact path="/" component={LoginPage}></Route>
            <Route exact path="/registration" component={RegistrationPage}></Route>
            <Route exact path="/profile" component={Settings}></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
