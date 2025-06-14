"use client";
import React from "react";
import NavBar from "./api/components/NavBar";
import Image from "next/image";

export default function HomePage() {
  return (
    <div className="p-4 bg-gray-100">
      <h1 className="text-3xl font-bold text-wine">Welcome to WineBaby</h1>
      <p className="text-lg text-gray-700">
        Your one-stop destination for wine enthusiasts.
      </p>
      <NavBar />
      <Image
        width={1920}
        height={1080}
        src="/public/images/winebaby-graphic.avif"
        alt="WineBaby Logo"
        className="mt-4 w-48 h-auto"
      />
    </div>
  );
}
