import React, { useState } from "react";
import axios from "axios";

import "./Signin.css";
import ProfileImage from "../../assets/images/profile.png";
import TemplateModal from "../Modals/TemplateModal";
import { useHistory } from "react-router-dom";

const Signin = () => {
  //Variables and Constants
  const initialState = {
    email: "",
    password: "",
  };

  const history = useHistory();

  //States
  const [user, setUser] = useState(initialState);
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

  const signInHandler = async () => {
    var formData = new FormData();
    formData.append("email", user.email);
    formData.append("password", user.password);

    const response = await axios({
      url: "http://localhost/dynamicimageslideshow/backend_php/signin.php",
      method: "post",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      data: formData,
    });

    if (response.data.status === "success") {
      history.push({
        pathname: "/dashboard",
        state: {
          email: response.data.email,
          password: response.data.password,
        },
      });
    } else {
      setUser(initialState);
      setModal({
        isShown: true,
        ModalTitle: response.data.ModalTitle,
        ModalBody: response.data.ModalBody,
      });
    }
  };

  return (
    <div className="Signin">
      <img
        src={ProfileImage}
        alt="Profile"
        style={{ width: "30%", margin: "2% 30%" }}
      />

      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Email</span>
        </div>
        <input
          type="email"
          className="form-control customform"
          placeholder="Enter Email"
          name="email"
          value={user.email}
          onChange={setUserHandler}
        />
      </div>
      <div className="input-group mb-3">
        <div className="input-group-prepend d-none d-sm-block">
          <span className="input-group-text customtext">Password</span>
        </div>
        <input
          type="password"
          className="form-control customform"
          placeholder="Enter Password"
          name="password"
          value={user.password}
          onChange={setUserHandler}
        />
      </div>
      <button
        type="button"
        onClick={signInHandler}
        className="btn w-100 custombutton"
      >
        SIGN IN
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

export default Signin;
