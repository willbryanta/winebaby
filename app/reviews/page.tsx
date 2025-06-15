"use client";

import React, { useEffect, useState } from "react";
import NavBar from "../api/components/NavBar";
import { wines } from "../api/auth/data/mockWineData";
import ReviewCard from "../api/components/ReviewCard";

export default function ReviewsPage() {
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
    <div className="flex flex-col w-screen h-screen bg-gray-100 overflow-hidden">
      <NavBar />
      <div className="flex-1 overflow-y-auto py-6">
        {wines.map((wine) => (
          <ReviewCard
            key={wine.id}
            reviews={wine.reviews}
            wineName={wine.name}
            wineType={wine.type ?? ""}
          />
        ))}
      </div>
    </div>
  );
}
