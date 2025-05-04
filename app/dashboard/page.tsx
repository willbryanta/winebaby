"use client";

import React from "react";
import NavBar from "../src/components/NavBar/NavBar";
import { useSession } from "next-auth/react";

declare module "next-auth" {
  interface Session {
    user: {
      name?: string | null;
      email?: string | null;
      image?: string | null;
      token?: string;
    };
  }
}

type Review = {
  ID: number;
  WineID: number;
  Comment: string;
  ReviewDate: string;
  ReviewDateTime: string;
  RewviewDateTimeUTC: string;
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
    RewviewDateTimeUTC: "2023-10-01T12:00:00Z",
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
    RewviewDateTimeUTC: "2023-10-02T14:00:00Z",
    Title: "Okay",
    Description: "It was okay, but I've had better.",
    Rating: 3,
  },
];

const Dashboard: React.FC = () => {
  const { data: session, status } = useSession();

  if (status === "loading") {
    return <div>Loading...</div>;
  }
  if (!session) {
    return <p>Access Denied</p>;
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <NavBar />
      <h1 className="text-2xl font-bold mb-4">Wine Reviews Dashboard</h1>
      <div className="w-full max-w-2xl bg-white shadow-md rounded-lg p-6">
        {reviews.map((review) => (
          <div key={review.ID} className="mb-4 p-4 border-b">
            <h2 className="text-xl font-semibold">{review.Title}</h2>
            <p className="text-gray-600">{review.Comment}</p>
            <p className="text-sm text-gray-500">{review.ReviewDate}</p>
            <p className="text-sm text-gray-500">Rating: {review.Rating}/5</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Dashboard;
