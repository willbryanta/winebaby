// /Users/will/Documents/Software_Engineering_Projects/winebaby/app/dashboard/page.tsx

import React from "react";
import NavBar from "../src/components/NavBar/NavBar";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const Dashboard: React.FC = () => {
  const { data: session, status } = useSession();
  const [data, setData] = useState(null);

  useEffect(() => {
    if (session?.user?.token) {
      fetch(`http://localhost:8080/${process.env.PROTECTED_ENDPOINT}`, {
        headers: {
          Authorization: `Bearer ${session.user.token}`,
        },
      })
        .then((res) => res.json())
        .then((data) => setData(data))
        .catch((err) => console.error(err));
    }
  }, [session]);

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
