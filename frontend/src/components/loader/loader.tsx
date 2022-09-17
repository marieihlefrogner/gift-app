import React from "react";
import { Vortex } from "react-loader-spinner";
import "./loader.scss";

const Loader = () => (
  <div className="loader">
    <Vortex
      visible={true}
      height="80"
      width="80"
      ariaLabel="vortex-loading"
      wrapperStyle={{}}
      wrapperClass="vortex-wrapper"
      colors={['#fd8800', '#fd008f', '#9700fd', '#0054fd', '#05c7e6', '#4bd58d']}
    />
  </div>
);

export default Loader;