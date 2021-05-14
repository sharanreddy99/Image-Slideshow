import React, { useState } from "react";
import { useHistory, useLocation } from "react-router-dom";
import "./EditUser.css";
import axios from "axios";

import ProfileImage from "../../assets/images/profile.png";
import TemplateModal from "../Modals/TemplateModal";
import AuthorizationModal from "../Modals/AuthorizationModal";

const EditUser = () => {
  //Constants and Variables
  const location = useLocation();
  const history = useHistory();
  const oldemail = location.state && location.state.email;

  const initialState = {
    firstname: location.state && location.state.firstname,
    lastname: location.state && location.state.lastname,
    email: location.state && location.state.email,
    mobile: location.state && location.state.mobile,
    password: location.state && location.state.password,
    confirmpassword: location.state && location.state.password,
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
  const [authmodal, setAuthModal] = useState({
    isShown: false,
    ModalTitle: "",
    ModalBody: "",
  });

  //Handlers
  const isChanged = () => {
    return !(
      user.firstname === location.state.firstname &&
      user.lastname === location.state.lastname &&
      user.email === location.state.email &&
      user.mobile === location.state.mobile &&
      user.password === location.state.password &&
      user.confirmpassword === location.state.password
    );
  };

  const setUserHandler = async (e) => {
    const name = e.target.name;
    const value = e.target.value;

    setUser({ ...user, [name]: value });
  };

  const updateUserHandler = async () => {
    if (!isChanged()) {
      setModal({
        isShown: true,
        ModalTitle: "Alert...",
        ModalBody: "No changes made...",
      });
      return;
    }

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
      formData.append("oldemail", oldemail);
      formData.append("token", location.state && location.state.token);

      try {
        const response = await axios({
          url: "/updateuser",
          method: "post",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          data: formData,
        });

        setAuthModal({
          isShown: true,
          ModalTitle: response.data.ModalTitle,
          ModalBody: response.data.ModalBody,
        });
      } catch (error) {
        if (error.response && error.response.status === 401) {
          setAuthModal({
            isShown: true,
            ModalTitle: error.response.data.ModalTitle,
            ModalBody: error.response.data.ModalBody,
          });
        } else {
          setModal({
            isShown: true,
            ModalTitle: error.response.data.ModalTitle,
            ModalBody: error.response.data.ModalBody,
          });
        }
      }
    }
  };
  return (
    <div className="EditUser">
      <img
        src={ProfileImage}
        alt="Profile"
        style={{ width: "15%", margin: "0% 40%" }}
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
        onClick={updateUserHandler}
        className="btn w-100 custombutton"
      >
        UPDATE USER
      </button>
      <TemplateModal
        isShown={modal.isShown}
        setIsShown={setModal}
        ModalTitle={modal.ModalTitle}
        ModalBody={modal.ModalBody}
      />
      <AuthorizationModal
        isShown={authmodal.isShown}
        setIsShown={setAuthModal}
        ModalTitle={authmodal.ModalTitle}
        ModalBody={authmodal.ModalBody}
      />
    </div>
  );
};

export default EditUser;
