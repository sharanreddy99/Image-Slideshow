import axios from "axios";
import React, { useState } from "react";
import { useHistory, useLocation } from "react-router-dom";
import "./Dashboard.css";
import TemplateModal from "../Modals/TemplateModal";

const Dashboard = () => {
  //Constants and Variables
  const history = useHistory();
  const location = useLocation();

  //States
  const [image, setImage] = useState(null);
  const [sepImage, setSepImage] = useState("");
  const [allImages, setAllImages] = useState(location.state.allImages);
  const [modal, setModal] = useState({
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
    formData.append("email", location.state.email);

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
      setCounter(arraylength);
    }

    setSepImage("");
    setImage(null);
  };

  const fetchImage = async (filename) => {
    var formData = new FormData();
    formData.append("filename", filename);

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

  const logoutHandler = () => {
    history.push("/login");
  };

  if (!(location.state && location.state.email)) {
    history.push("/login");
  }

  return (
    <div className="Dashboard">
      <div className="Dashboard__Gallery Dashboard__container">
        {allImages.length > 0 ? (
          <img src={allImages[counter]} alt="" />
        ) : (
          <div className="emptyimage">
            <h4>No Image to Display</h4>
          </div>
        )}
        <div className="row mt-3 p-5">
          <div className="col">
            <button
              type="button"
              name="previous"
              className="btn w-100 custombutton"
              onClick={() => {
                if (allImages.length > 0) {
                  setCounter(
                    (counter - 1 + allImages.length) % allImages.length
                  );
                }
              }}
            >
              Previous
            </button>
          </div>
          <div className="col">
            <button
              type="button"
              name="next"
              className="btn w-100 custombutton"
              onClick={() => {
                if (allImages.length > 0) {
                  setCounter((counter + 1) % allImages.length);
                }
              }}
            >
              Next
            </button>
          </div>
        </div>
      </div>
      <div className="Dashboard__buttonsdiv Dashboard__container">
        <h3>Upload Image</h3>
        <div className="Dashboard__buttons_container">
          <div className="form-group">
            <label style={{ fontWeight: "bold" }}>Upload Image:</label>
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
          </div>

          <button
            type="button"
            className="btn w-100 custombutton"
            onClick={addImageHandler}
          >
            Add to Gallery
          </button>

          <button
            type="button"
            className="btn w-100 mt-3 custombutton"
            onClick={logoutHandler}
          >
            Log Out
          </button>
        </div>
      </div>
      <TemplateModal
        isShown={modal.isShown}
        setIsShown={setModal}
        ModalTitle={modal.ModalTitle}
        ModalBody={modal.ModalBody}
      />
    </div>
  );
};

export default Dashboard;
