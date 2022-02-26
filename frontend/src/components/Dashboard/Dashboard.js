import axios from "axios";
import React, { useState } from "react";
import { useHistory, useLocation } from "react-router-dom";
import "./Dashboard.css";
import TemplateModal from "../Modals/TemplateModal";
import AuthorizationModal from "../Modals/AuthorizationModal";

const Dashboard = () => {
  //Constants and Variables
  const history = useHistory();
  const location = useLocation();
  const token = location.state && location.state.token;

  //States
  const [image, setImage] = useState(null);
  const [sepImage, setSepImage] = useState("");
  const [allImages, setAllImages] = useState(
    (location.state && location.state.allImages) || []
  );
  const [filenames, setFilenames] = useState(
    (location.state && location.state.filenames) || []
  );
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
  const [counter, setCounter] = useState(0);

  //Handlers
  const addImageHandler = async () => {
    if (image === null) {
      setModal({
        isShown: true,
        ModalTitle: "Choose an Image...",
        ModalBody: "Please upload an image to proceed...",
      });
      return;
    }

    var formData = new FormData();
    formData.append("imagefile", image);
    formData.append("token", token);

    try {
      const response = await axios({
        url: "/imageupload",
        method: "post",
        headers: {
          "Content-Type": "multipart/form-data",
        },
        data: formData,
      });

      if (response.data.status === "failure") {
        setModal({
          isShown: true,
          ModalTitle: response.data.ModalTitle,
          ModalBody: response.data.ModalBody,
        });
      } else {
        const arraylength = allImages.length;
        const imagefile = await fetchImage(response.data.filename);

        var extension = response.data.filename.split(".")[1];
        setAllImages([
          ...allImages,
          "data:image/" + extension + ";base64," + imagefile.imagedata,
        ]);
        setFilenames([...filenames, response.data.filename]);
        setCounter(arraylength);
      }
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

    setSepImage("");
    setImage(null);
  };

  const fetchImage = async (filename) => {
    var formData = new FormData();
    formData.append("filename", filename);
    formData.append("token", token);

    const response = await axios({
      url: "/getsingleimage",
      method: "post",
      headers: {
        "Content-Type": "multipart/form-data",
      },
      data: formData,
    });

    return response.data;
  };

  const deleteImageHandler = async () => {
    try {
      if (filenames.length === 0) return;

      var formData = new FormData();
      formData.append("filename", filenames[counter]);
      formData.append("token", token);

      const response = await axios({
        url: "/deletesingleimage",
        method: "post",
        headers: {
          "Content-Type": "multipart/form-data",
        },
        data: formData,
      });

      setModal({
        isShown: true,
        ModalTitle: response.data.ModalTitle,
        ModalBody: response.data.ModalBody,
      });

      const newFileNames = filenames.filter((file, ind) => ind !== counter);
      const newAllImages = allImages.filter((image, ind) => ind !== counter);

      history.replace({
        pathname: "/dashboard",
        state: {
          ...location.state,
          filenames: newFileNames,
          allImages: newAllImages,
        },
      });

      setFilenames(newFileNames);
      setAllImages(newAllImages);

      if (counter - 1 >= 0) {
        setCounter(counter - 1);
      }
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

  return (
    <div className="Dashboard">
      <div className="Dashboard__left Dashboard__container">
        {allImages.length > 0 ? (
          <div>
            <img src={allImages[counter]} alt="" />
          </div>
        ) : (
          <div className="emptyimage">
            <h2>No Image to Display</h2>
          </div>
        )}

        <i className="fa fa-trash" onClick={deleteImageHandler}></i>
        <i
          className="fa fa-arrow-left"
          onClick={() => {
            if (allImages.length > 0) {
              setCounter((counter - 1 + allImages.length) % allImages.length);
            }
          }}
        ></i>
        <i
          className="fa fa-arrow-right"
          onClick={() => {
            if (allImages.length > 0) {
              setCounter((counter + 1) % allImages.length);
            }
          }}
        ></i>
      </div>
      <div className="Dashboard__right">
        <input
          type="file"
          className="form-control customform"
          name="image"
          value={sepImage}
          onChange={(e) => {
            setSepImage(e.target.value);
            setImage(e.target.files[0]);
          }}
        />
        <button
          type="button"
          className="btn w-100 mt-3 custombutton"
          onClick={addImageHandler}
        >
          Upload Image
        </button>
      </div>
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

export default Dashboard;
