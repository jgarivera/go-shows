import React from "react";
import Navbar from "./components/Navbar";
import { Route, Routes } from "react-router-dom";
import Shows from "./pages/Shows";
import About from "./pages/About";

export default function App(): JSX.Element {
  return (
    <React.Fragment>
      <Navbar />
      <Routes>
        <Route path="/" element={<Shows />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </React.Fragment>
  );
}
