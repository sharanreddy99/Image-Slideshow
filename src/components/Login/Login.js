import React from "react";
import "./Login.css";
import Signin from "../Signin/Signin";
import Signup from "../Signup/Signup";

const Login = () => {
  return (
    <div className="Login">
      <div className="Login__subcontainer">
        <Signin />
      </div>
      <div className="Login__subcontainer">
        <Signup />
      </div>
    </div>
  );
};

export default Login;
