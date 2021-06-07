import { AppBar, Toolbar, Grid, Button } from "@material-ui/core";
import { Link } from "react-router-dom";

const NavBar = () => {
  const user = localStorage.getItem("user");

  const clearLocalStorage = () => {
    localStorage.clear();
  };

  const NavBarForUnregisteredUser = (
    <Toolbar style={{ backgroundColor: "white" }}>
      <Grid container>
        <Grid item xs={6}></Grid>
        <Grid item xs={6} container style={{ textAlign: "right" }}>
          <Grid item xs={2} />
          <Grid item xs={2} />
          <Grid item xs={2}></Grid>
          <Grid item xs={2}>
            <Button variant="contained" color="primary">
              <Link
                to="/login"
                style={{ textDecoration: "none", color: "white" }}
              >
                Log in
              </Link>
            </Button>
          </Grid>
          <Grid item xs={2}>
            <Button variant="text" onClick={clearLocalStorage}>
              <Link to="/registration" style={{ textDecoration: "none" }}>
                Sing up
              </Link>
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Toolbar>
  );

  const NavBarForRegistredUser = (
    <Toolbar style={{ backgroundColor: "white" }}>
      <Grid container>
        <Grid item xs={6}></Grid>
        <Grid item xs={6} container style={{ textAlign: "right" }}>
          <Grid item xs={2} />
          <Grid item xs={2} />
          <Grid item xs={2}></Grid>
          <Grid item xs={2}></Grid>
          <Grid item xs={2}>
            <Button variant="text" onClick={clearLocalStorage}>
              <a href="/" style={{ textDecoration: "none" }}>
                Sing out
              </a>
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Toolbar>
  );

  return (
    <>
      <AppBar position="static">
        {(user === null || user === undefined) && NavBarForUnregisteredUser}
        {user !== null && user !== undefined && NavBarForRegistredUser}
      </AppBar>
    </>
  );
};

export default NavBar;
