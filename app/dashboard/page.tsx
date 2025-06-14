"use client";

import React, { useEffect, useState } from "react";
import NavBar from "@/app/api/components/NavBar";
import { wines } from "@/app/api/auth/data/mockWineData";

const Dashboard: React.FC = () => {
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
      <div className="flex-1 flex flex-col items-center overflow-y-auto py-6">
        <h1 className="text-2xl font-bold mb-4">Wine Dashboard</h1>
        <div className="w-full max-w-4xl bg-white shadow-md rounded-lg p-6 mx-4">
          {wines.map((wine) => (
            <div key={wine.id} className="mb-4 p-4 border-b last:border-b-0">
              <h2 className="text-xl font-semibold">{wine.name}</h2>
              <p className="text-gray-600">{wine.manufacturer}</p>
              <p className="text-sm text-gray-500">{wine.region}</p>
              <p className="text-sm text-gray-500">{wine.type}</p>
              <p className="text-sm text-gray-500">
                Alcohol content: {wine.alcoholContent}
              </p>
              <p className="text-sm text-gray-500">Price: {wine.price} AUD</p>
              <p className="text-sm text-gray-500">Rating: {wine.rating}/5</p>
              <p className="text-sm text-gray-500">
                Average Rating: {wine.averageRating}/5
              </p>
              <div className="mt-2">
                <p className="text-sm text-gray-500 font-semibold">Reviews:</p>
                <div className="pl-4 space-y-2">
                  {wine.reviews.map((review) => (
                    <div key={review.ID} className="text-sm text-gray-500">
                      <h3 className="font-medium">{review.Title}</h3>
                      <p>{review.Content}</p>
                      <p>Rating: {review.Rating}/5</p>
                      <p>Date: {review.ReviewDate}</p>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
