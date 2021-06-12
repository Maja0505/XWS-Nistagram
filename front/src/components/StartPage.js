import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Settings from "./Settings.js"
import UserHomePage from "./UserHomePage.js";
import PostDialog from './PostDialog'

const StartPage = () => {
  return <div>
      <Router>
          <Switch>
            <Route  path="/accounts" component={Settings}></Route>
            <Route  path="/dialog" component={PostDialog}></Route>
            <Route  path="/homePage" component={UserHomePage}></Route>

          </Switch>
      </Router>

  </div>;
};

export default StartPage;
