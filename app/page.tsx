"use client";
import React from "react";
import NavBar from "./src/components/NavBar/NavBar";

export default function HomePage() {
  return (
    <div className="p-4 bg-gray-100">
      <h1 className="text-3xl font-bold text-wine">Welcome to WineBaby</h1>
      <p className="text-lg text-gray-700">
        Your one-stop destination for wine enthusiasts.
      </p>
      <NavBar />
    </div>
  );
}
