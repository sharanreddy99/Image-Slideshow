import React, { useState } from "react";
import axios from "axios";

import "./Signup.css";
import ProfileImage from "../../assets/images/profile.png";
import TemplateModal from "../Modals/TemplateModal";

const Signup = () => {
  //Variables and Constants
  const initialState = {
    firstname: "",
    lastname: "",
    email: "",
    mobile: "",
    password: "",
    confirmpassword: "",
  };

  const initialState2 = {
    firstname: null,
    lastname: null,
    email: null,
    mobile: null,
    password: null,
    confirmpassword: null,
  };

  //States
  const [user, setUser] = useState(initialState);
  const [validUser, setValidUser] = useState(initialState2);
  const [modal, setModal] = useState({
    isShown: false,
    ModalTitle: "",
    ModalBody: "",
  });

  //Handlers
  const setUserHandler = async (e) => {
    const name = e.target.name;
    const value = e.target.value;

    setUser({ ...user, [name]: value });
  };

  const signupHandler = async () => {
    var pattern = /^[A-Za-z]+[A-Za-z ]{4,50}$/i;
    var firstname = false;
    var lastname = false;
    var email = false;
    var mobile = false;
    var password = false;
    var confirmpassword = false;

    firstname = pattern.test(user.firstname);
    lastname = pattern.test(user.lastname);

    pattern = /^[A-Za-z][A-Za-z0-9.]+@[A-Za-z]{1,20}\.[A-Za-z]{1,10}$/i;
    email = pattern.test(user.email);

    pattern = /^\d{10}$/i;
    mobile = pattern.test(user.mobile);

    pattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*.,])[A-Za-z\d!@#$%^&*.,]{8,50}$/s;
    password = pattern.test(user.password);
    confirmpassword = pattern.test(user.confirmpassword);

    if (password && confirmpassword) {
      password = user.password === user.confirmpassword;
      confirmpassword = password;
    } else {
      password = false;
      confirmpassword = false;
    }

    setValidUser({
      firstname,
      lastname,
      email,
      mobile,
      password,
      confirmpassword,
    });

    if (
      [firstname, lastname, email, mobile, password, confirmpassword].every(
        (value) => value === true
      )
    ) {
      var formData = new FormData();
      formData.append("firstname", user.firstname);
      formData.append("lastname", user.lastname);
      formData.append("email", user.email);
      formData.append("mobile", user.mobile);
      formData.append("password", user.password);

      const response = await axios({
        url: "/signup",
        method: "post",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        data: formData,
      });

      setModal({
        isShown: true,
        ModalTitle: response.data.ModalTitle,
        ModalBody: response.data.ModalBody,
      });

      if (response.status === 201) {
        setUser(initialState);
        setValidUser(initialState2);
      }
    }
  };

  return (
    <div className="Signup">
      <img
        src={ProfileImage}
        alt="Profile"
        style={{ width: "30%", margin: "2% 30%" }}
      />

      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">First Name</span>
        </div>
        <input
          type="text"
          className={`form-control ${
            validUser.firstname === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter First Name"
          name="firstname"
          value={user.firstname}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Last Name</span>
        </div>
        <input
          type="text"
          className={`form-control ${
            validUser.lastname === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter Last Name"
          name="lastname"
          value={user.lastname}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Email</span>
        </div>
        <input
          type="email"
          className={`form-control ${
            validUser.email === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter Email"
          name="email"
          value={user.email}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block ">
          <span className="input-group-text customtext">Mobile Number</span>
        </div>
        <input
          type="text"
          className={`form-control ${
            validUser.mobile === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter Mobile Number"
          name="mobile"
          value={user.mobile}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Password</span>
        </div>
        <input
          type="password"
          className={`form-control ${
            validUser.password === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter Password"
          name="password"
          value={user.password}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Confirm Password</span>
        </div>
        <input
          type="password"
          className={`form-control ${
            validUser.confirmpassword === false ? "invalidform" : "customform"
          }`}
          placeholder="Enter Password Again"
          name="confirmpassword"
          value={user.confirmpassword}
          onChange={setUserHandler}
        />
      </div>
      <button
        type="button"
        onClick={signupHandler}
        className="btn w-100 custombutton"
      >
        SIGN UP
      </button>
      <TemplateModal
        isShown={modal.isShown}
        setIsShown={setModal}
        ModalTitle={modal.ModalTitle}
        ModalBody={modal.ModalBody}
      />
    </div>
  );
};

export default Signup;
