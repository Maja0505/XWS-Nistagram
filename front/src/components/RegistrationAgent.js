import { useState } from "react";
import {
  TextField,
  Grid,
  Button,
  RadioGroup,
  FormLabel,
  FormControlLabel,
  Radio,
} from "@material-ui/core";
import validator from "validator";
import axios from "axios";

const RegistrationAgent = ({ setRedirection }) => {
  const [user, setUser] = useState({ gender: 1 });

  const handleSubmitClick = () => {
    if (!validator.isEmail(user.Email)) {
      alert("Invalid email format \n Example : pera@pera.com");
      setUser({ ...user, Email: "" });
      return;
    }

    console.log(user);
    let userForRegistration = {
      ...user,
      DateOfBirth: user.DateOfBirth + "T00:00:00+01:00",
      Approved: false,
    };
    axios
      .post("/api/user/agent-registration-request/create", userForRegistration)
      .then((res) => {
        setRedirection(true);
      })
      .catch((error, res) => {
        //alert(error);
        console.log(error.message);
      });
  };

  return (
    <div>
      <Grid container style={{ marginTop: "2%" }}>
        <Grid item xs={3}></Grid>
        <Grid item xs={6}>
          <form
            noValidate
            autoComplete="off"
            style={{ width: "80%", margin: "auto" }}
          >
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="First Name"
              fullWidth
              onChange={(e) => setUser({ ...user, FirstName: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Last Name"
              fullWidth
              onChange={(e) => setUser({ ...user, LastName: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Email"
              fullWidth
              onChange={(e) => setUser({ ...user, Email: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Phone Number"
              fullWidth
              onChange={(e) =>
                setUser({ ...user, PhoneNumber: e.target.value })
              }
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Username"
              fullWidth
              onChange={(e) => setUser({ ...user, Username: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              type="password"
              placeholder="Password"
              fullWidth
              onChange={(e) => setUser({ ...user, Password: e.target.value })}
            />
            <br></br>
            <br></br>
            <TextField
              color="primary"
              variant="outlined"
              size="small"
              placeholder="Web Site"
              fullWidth
              onChange={(e) => setUser({ ...user, WebSite: e.target.value })}
            />
            <br></br>
            <br></br>
            <Grid container>
              <RadioGroup
                row
                aria-label="gender"
                name="gender1"
                defaultValue="female"
                onChange={(e) =>
                  setUser({
                    ...user,
                    gender: e.target.value === "female" ? 0 : 1,
                  })
                }
              >
                <FormControlLabel
                  value="female"
                  control={<Radio color="primary" />}
                  label="Female"
                />
                <FormControlLabel
                  value="male"
                  control={<Radio color="primary" />}
                  label="Male"
                />
              </RadioGroup>
              <br></br>
              <br></br>
              <Grid container>
                <FormLabel style={{ marginTop: "3%" }}>Date of birth</FormLabel>
                <TextField
                  id="date"
                  type="date"
                  variant="outlined"
                  color="primary"
                  size="small"
                  onChange={(e) =>
                    setUser({ ...user, DateOfBirth: e.target.value })
                  }
                  InputLabelProps={{
                    shrink: true,
                  }}
                  style={{ marginTop: "0.5%", marginLeft: "3.8%" }}
                />
              </Grid>
            </Grid>
            <br></br>
            <br></br>
          </form>
        </Grid>
        <Grid item xs={3}></Grid>
      </Grid>
      <Button variant="contained" color="primary" onClick={handleSubmitClick}>
        Send Registration Request
      </Button>
    </div>
  );
};

export default RegistrationAgent;
