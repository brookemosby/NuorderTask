import React from "react";
import { render } from "react-dom";

import Autocomplete from "./Autocomplete";

require("./styles.css");

function App() {
  return (
    <div style={{display:'flex', justifyContent:'center', alignItems:'center', height: '100vh'}} >
      <div style={{display:'block'}} >
        <h1 > Search React Issues</h1>
        <Autocomplete />
      </div>
    </div>
  );
}

const container = document.createElement("div");
document.body.appendChild(container);
render(<App />, container);
