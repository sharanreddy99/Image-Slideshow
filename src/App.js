import React from "react";
import "./App.css";
import RouterSetup from "./RouterSetup";
import axios from "axios";
import {
  startingPartGO,
  endingPartGO,
  startingPartPHP,
  endingPartPHP,
} from "./baseURL";

axios.interceptors.request.use(
  function (config) {
    config.url = startingPartPHP + config.url + endingPartPHP;
    // config.url = startingPartGO + config.url + endingPartGO;
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
