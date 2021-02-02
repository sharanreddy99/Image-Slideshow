import React from "react";
import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from "react-router-dom";
import Login from "./components/Login/Login";
import Dashboard from "./components/Dashboard/Dashboard";
import Error from "./components/Error/Error";
import PageTemplate from "./components/PageTemplate/PageTemplate";
import EditUser from "./components/EditUser/EditUser";

const RouterSetup = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Redirect to="/login" />
        </Route>
        <Route exact path="/login">
          <Login />
        </Route>
        <Route exact path="/dashboard">
          <PageTemplate>
            <Dashboard />
          </PageTemplate>
        </Route>
        <Route exact path="/edituser">
          <PageTemplate>
            <EditUser />
          </PageTemplate>
        </Route>
        <Route path="*">
          <Error />
        </Route>
      </Switch>
    </Router>
  );
};

export default RouterSetup;
