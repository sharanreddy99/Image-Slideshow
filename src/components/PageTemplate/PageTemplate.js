import React, { useState, useEffect } from "react";
import { useLocation, useHistory } from "react-router-dom";
import axios from "axios";

import "./PageTemplate.css";
import AuthorizationModal from "../Modals/AuthorizationModal";
import TemplateModal from "../Modals/TemplateModal";

const PageTemplate = (props) => {
  const location = useLocation();
  const history = useHistory();

  const token = location.state && location.state.token;
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

  const isAuthorized = async () => {
    try {
      var formData = new FormData();
      formData.append("token", token);

      const response = await axios({
        url: "/isvaliduser",
        method: "post",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        data: formData,
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
  };

  const editUserHandler = async () => {
    try {
      var formData = new FormData();
      formData.append("token", token);

      const response = await axios({
        url: "/fetchuser",
        method: "post",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        data: formData,
      });

      history.push({
        pathname: "/edituser",
        state: {
          ...location.state,
          firstname: response.data.firstname,
          lastname: response.data.lastname,
          mobile: response.data.mobile,
          password: response.data.password,
          email: response.data.email,
        },
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
  };

  const logoutHandler = async () => {
    var formData = new FormData();
    formData.append("token", token);
    await axios({
      url: "/logout",
      method: "post",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      data: formData,
    });
    history.push("/login");
  };

  //Effects
  useEffect(() => {
    isAuthorized();
  }, []);

  return (
    <div className="PageTemplate">
      <div className="PageTemplate__navbar">
        <h5
          onClick={() => {
            history.push({
              pathname: "/dashboard",
              state: {
                email: location.state.email,
                token: location.state.token,
                allImages: location.state.allImages,
                filenames: location.state.filenames,
              },
            });
          }}
        >
          Slideshow
        </h5>

        <div className="PageTemplate__navbar_icons">
          <i className="fa fa-cog" onClick={editUserHandler}></i>
          <i className="fa fa-sign-out" onClick={logoutHandler}></i>
        </div>
      </div>
      <div className="PageTemplate__body">{props.children}</div>

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

export default PageTemplate;
