import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import "./App.css";
import RegistrationPage from "./components/RegistartionPage.js";
import LoginPage from "./components/LoginPage.js";

function App() {
  return (
    <div>
      <Router>
        <div className="App">
          <Switch>
            <Route exact path="/" component={LoginPage}></Route>
          </Switch>
          <Switch>
            <Route path="/registration" component={RegistrationPage}></Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
