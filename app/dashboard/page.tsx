// /Users/will/Documents/Software_Engineering_Projects/winebaby/app/dashboard/page.tsx

import React from "react";

type Review = {
  id: number;
  user: string;
  wine: string;
  rating: number;
  comment: string;
};

const reviews: Review[] = [
  {
    id: 1,
    user: "Alice",
    wine: "Chardonnay",
    rating: 4,
    comment: "Crisp and refreshing!",
  },
  {
    id: 2,
    user: "Bob",
    wine: "Merlot",
    rating: 5,
    comment: "Rich and full-bodied.",
  },
  {
    id: 3,
    user: "Charlie",
    wine: "Pinot Noir",
    rating: 3,
    comment: "Light and fruity.",
  },
];

const Dashboard: React.FC = () => {
  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h1 className="text-2xl font-bold mb-4">User Feed</h1>
      <div className="space-y-4">
        {reviews.map((review) => (
          <div
            key={review.id}
            className="p-4 bg-white rounded-lg shadow-md border border-gray-200"
          >
            <h2 className="text-lg font-semibold">{review.wine}</h2>
            <p className="text-sm text-gray-600">Reviewed by: {review.user}</p>
            <p className="text-yellow-500">
              Rating: {"⭐".repeat(review.rating)}
            </p>
            <p className="mt-2">{review.comment}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Dashboard;
