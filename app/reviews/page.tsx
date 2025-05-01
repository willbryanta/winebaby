import React, { useState, useEffect } from "react";

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

const ReviewsPage: React.FC = () => {
  const [reviews, setReviews] = useState<Review[]>([]);
  const [expandedReviewId, setExpandedReviewId] = useState<number | null>(null);

  useEffect(() => {
    // Fetch reviews from the API
    const fetchReviews = async () => {
      try {
        const response = await fetch("/api/reviews");
        const data = await response.json();
        setReviews(data);
      } catch (error) {
        console.error("Error fetching reviews:", error);
      }
    };

    fetchReviews();
  }, []);

  const toggleExpand = (id: number) => {
    setExpandedReviewId(expandedReviewId === id ? null : id);
  };
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-2xl font-bold mb-4">Wine Reviews</h1>
      <div className="w-full max-w-2xl bg-white shadow-md rounded-lg p-6">
        {reviews.map((review) => (
          <div key={review.ID} className="mb-4 p-4 border-b">
            <h2
              onClick={() => toggleExpand(review.ID)}
              className="text-xl font-semibold cursor-pointer"
            >
              {review.Title}
            </h2>
            {expandedReviewId === review.ID && (
              <div>
                <p className="text-gray-600">{review.Comment}</p>
                <p className="text-sm text-gray-500">{review.ReviewDate}</p>
                <p className="text-sm text-gray-500">
                  Rating: {review.Rating}/5
                </p>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default ReviewsPage;
