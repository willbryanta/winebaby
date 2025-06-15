"use client";

import React, { useEffect, useState } from "react";
import NavBar from "../api/components/NavBar";

export default function ProfilePage() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const verifyToken = async () => {
      try {
        const res = await fetch("http://localhost:8080/verify-token", {
          method: "GET",
          credentials: "include",
        });

        if (res.ok) {
          setIsAuthenticated(true);
        } else {
          setIsAuthenticated(false);
        }
      } catch (error) {
        console.error("Error verifying token:", error);
        setIsAuthenticated(false);
      } finally {
        setIsLoading(false);
      }
    };
    verifyToken();
  }, []);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-screen">
        Loading...
      </div>
    );
  }
  if (!isAuthenticated) {
    return (
      <p className="flex items-center justify-center h-screen">Access Denied</p>
    );
  }

  return (
    <div>
      <NavBar />
      <div className="flex items-center justify-center h-screen">
        <h1 className="text-4xl font-bold">Currently still building...</h1>
      </div>
    </div>
  );
}
