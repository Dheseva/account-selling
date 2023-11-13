import React from "react";
import ReactDOM from "react-dom/client";

export default function Application() {
  return <div>Application</div>;
}

const root = ReactDOM.createRoot(document.querySelector("#application")!);
root.render(<Application />);
