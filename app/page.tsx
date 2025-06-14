"use client";
import React from "react";
import NavBar from "./api/components/NavBar";

export default function Dashboard() {
  return (
    <div className="w-screen h-screen flex flex-col bg-gray-100">
      <NavBar />
      <div className="flex-1 flex flex-col items-center justify-center">
        <h1 className="text-[4.8rem] font-bold text-wine">
          Welcome to WineBaby
        </h1>
        <p className="text-[2.2rem] text-gray-700">
          Your one-stop destination for wine enthusiasts.
        </p>
      </div>
    </div>
  );
}
