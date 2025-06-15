"use client";

import { ReviewCardProps } from "../auth/types/page";

export default function ReviewCard({
  reviews,
  wineName,
  wineType,
}: ReviewCardProps) {
  return (
    <div className="w-full max-w-4xl mx-auto bg-white shadow-md rounded-lg p-6 mb-4">
      <div className="mb-4">
        <h2 className="text-xl font-semibold text-gray-800">{wineName}</h2>
        <p className="text-sm text-gray-500">{wineType}</p>
      </div>
      <div>
        <h3 className="text-lg font-medium text-gray-700 mb-2">Reviews</h3>
        {reviews && reviews.length > 0 ? (
          <div className="space-y-4">
            {reviews.map((review) => (
              <div
                key={review.ID}
                className="p-4 border border-gray-200 rounded-md"
              >
                <h4 className="text-sm font-medium text-gray-800">
                  {review.Title}
                </h4>
                <p className="text-sm text-gray-600">{review.Content}</p>
                <p className="text-sm text-gray-500">
                  Rating: {review.Rating}/5
                </p>
                <p className="text-sm text-gray-500">
                  Date: {review.ReviewDate}
                </p>
              </div>
            ))}
          </div>
        ) : (
          <p className="text-sm text-gray-500">No reviews available.</p>
        )}
      </div>
    </div>
  );
}
