import React from "react";
import NavBar from "../components/NavBar/NavBar";

const HomePage: React.FC = () => {
  return (
    <div>
      <h1>Welcome to WineBaby</h1>
      <p>Your one-stop destination for wine enthusiasts.</p>
      <NavBar />
    </div>
  );
};

export default HomePage;
