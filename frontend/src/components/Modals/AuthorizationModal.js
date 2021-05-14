import React from "react";
import { useHistory } from "react-router-dom";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

const AuthorizationModal = ({ isShown, setIsShown, ModalTitle, ModalBody }) => {
  const history = useHistory();

  const handleClose = () => {
    setIsShown({ isShown: false, ModalTitle, ModalBody });
    history.push("/login");
  };

  return (
    <Modal show={isShown} onHide={handleClose} backdrop="static">
      <Modal.Header
        closeButton
        style={{
          backgroundColor: "#5d001e",
          color: "white",
          textShadow: "2px 2px black",
        }}
      >
        <Modal.Title>{ModalTitle}</Modal.Title>
      </Modal.Header>
      <Modal.Body style={{ fontWeight: "bold" }}>{ModalBody}</Modal.Body>
      <Modal.Footer>
        <Button
          style={{
            backgroundColor: "#5d001e",
            color: "white",
            fontWeight: "bold",
          }}
          onClick={handleClose}
        >
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
export default AuthorizationModal;
