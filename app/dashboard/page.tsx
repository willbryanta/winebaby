"use client";

import React, { useEffect, useState } from "react";
import NavBar from "../src/components/NavBar/NavBar";

type Review = {
  ID: number;
  WineID: number;
  Comment: string;
  ReviewDate: string;
  ReviewDateTime: string;
  ReviewDateTimeUTC: string;
  Title: string;
  Description: string;
  Rating: number;
};

// TODO: Replace with actual data fetching logic
const reviews: Review[] = [
  {
    ID: 1,
    WineID: 101,
    Comment: "Great wine, loved the taste!",
    ReviewDate: "2023-10-01",
    ReviewDateTime: "2023-10-01T12:00:00Z",
    ReviewDateTimeUTC: "2023-10-01T12:00:00Z",
    Title: "Amazing!",
    Description: "This wine is fantastic. Highly recommend it.",
    Rating: 5,
  },
  {
    ID: 2,
    WineID: 102,
    Comment: "Not my favorite, but decent.",
    ReviewDate: "2023-10-02",
    ReviewDateTime: "2023-10-02T14:00:00Z",
    ReviewDateTimeUTC: "2023-10-02T14:00:00Z",
    Title: "Okay",
    Description: "It was okay, but I've had better.",
    Rating: 3,
  },
];

const Dashboard: React.FC = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const verifyToken = async () => {
      try {
        const res = await fetch("http://localhost:8080/verify", {
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
        <h1 className="text-2xl font-bold mb-4">Wine Reviews Dashboard</h1>
        <div className="w-full max-w-4xl bg-white shadow-md rounded-lg p-6 mx-4">
          {reviews.map((review) => (
            <div key={review.ID} className="mb-4 p-4 border-b last:border-b-0">
              <h2 className="text-xl font-semibold">{review.Title}</h2>
              <p className="text-gray-600">{review.Comment}</p>
              <p className="text-sm text-gray-500">{review.ReviewDate}</p>
              <p className="text-sm text-gray-500">Rating: {review.Rating}/5</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
