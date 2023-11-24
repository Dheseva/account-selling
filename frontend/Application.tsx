import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import "@fortawesome/fontawesome-svg-core/styles.css";
import Login from "./pages/Login";
// export default function Application() {
//   return <div>Application</div>;
// }

const root = ReactDOM.createRoot(document.querySelector("#application")!);
root.render(
  <BrowserRouter>
    <Routes>
      <Route index element={<Home />} />
      <Route path="/about" element={<About />} />
      <Route path="/login" element={<Login />} />
    </Routes>
  </BrowserRouter>
);
