import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Settings from "./Settings.js"
import UserHomePage from "./UserHomePage.js";
import ProfileDialog from './ProfileDialog'

const StartPage = () => {
  return <div>
      <Router>
          <Switch>
            <Route  path="/settings" component={Settings}></Route>
            <Route  path="/dialog" component={ProfileDialog}></Route>
            <Route  path="/homePage" component={UserHomePage}></Route>

          </Switch>
      </Router>

  </div>;
};

export default StartPage;
