import React from "react";
import "./App.css";
import RouterSetup from "./RouterSetup";
import axios from "axios";

const BACKEND_HOST = process.env.REACT_APP_IMAGE_VIEWER_BACKEND_HOST;

axios.interceptors.request.use(
  function (config) {
    switch (BACKEND_HOST) {
      case "go": {
        config.url = "/goapi" + config.url;
        break;
      }
      case "php": {
        config.url = "/phpapi" + config.url;
        break;
      }
      default: {
        break;
      }
    }

    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

const App = () => {
  return (
    <div className="App">
      <RouterSetup />
    </div>
  );
};

export default App;
