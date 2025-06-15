"use client";

import { useState, useEffect } from "react";
import { Wine, Review, WineCardProps } from "../auth/types/page";

export default function WineCard({ wines }: WineCardProps) {
  const [openFormWineId, setOpenFormWineId] = useState<number | null>(null);
  const [formData, setFormData] = useState({
    title: "",
    content: "",
    rating: 1,
  });
  const [userId, setUserId] = useState<number | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Fetch session data to get user ID
  useEffect(() => {
    const fetchSession = async () => {
      try {
        const res = await fetch("/api/session", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();
          setUserId(data.userId);
        } else {
          setUserId(null);
        }
      } catch (error) {
        console.error("Error fetching session:", error);
        setUserId(null);
      } finally {
        setIsLoading(false);
      }
    };
    fetchSession();
  }, []);

  const handleOpenForm = (wineId: number) => {
    setOpenFormWineId(wineId);
    setFormData({ title: "", content: "", rating: 1 });
  };

  const handleCloseForm = () => {
    setOpenFormWineId(null);
    setFormData({ title: "", content: "", rating: 1 });
  };

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]:
        name === "rating" ? Math.min(5, Math.max(1, parseInt(value))) : value,
    }));
  };

  const handleSubmit = async (wineId: number) => {
    if (!userId) {
      alert("You must be logged in to submit a review.");
      return;
    }

    const review: Review = {
      ID: Math.floor(Math.random() * 10000), //TODO: Mock ID; replace with server-generated ID
      WineID: wineId,
      UserID: userId,
      Title: formData.title,
      Content: formData.content,
      Rating: formData.rating,
      ReviewDate: new Date().toISOString().split("T")[0],
      ReviewDateTime: new Date().toISOString(),
    };

    try {
      // Mock API call; replace with actual endpoint
      const res = await fetch("/api/reviews", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(review),
      });
      if (!res.ok) {
        throw new Error("Failed to submit review");
      }
      handleCloseForm();
    } catch (error) {
      console.error("Error submitting review:", error);
      alert("Failed to submit review.");
    }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex-1 flex flex-col items-center overflow-y-auto py-6">
      <h1 className="text-2xl font-bold mb-4">Wine Dashboard</h1>
      <div className="w-full max-w-4xl bg-white shadow-md rounded-lg p-6 mx-4">
        {wines.map((wine: Wine) => (
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
                {wine.reviews &&
                  wine.reviews.map((review: Review) => (
                    <div key={review.ID} className="text-sm text-gray-500">
                      <h3 className="font-medium">{review.Title}</h3>
                      <p>{review.Content}</p>
                      <p>Rating: {review.Rating}/5</p>
                      <p>Date: {review.ReviewDate}</p>
                    </div>
                  ))}
              </div>
            </div>
            <button
              onClick={() => handleOpenForm(wine.id!)}
              disabled={!userId}
              className={`mt-2 px-4 py-2 rounded text-white ${
                userId
                  ? "bg-blue-500 hover:bg-blue-600"
                  : "bg-gray-400 cursor-not-allowed"
              }`}
            >
              Add Review
            </button>
            {openFormWineId === wine.id && (
              <div className="mt-4 p-4 bg-gray-50 rounded-lg">
                <h3 className="text-lg font-semibold mb-2">Add a Review</h3>
                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-700">
                      Title
                    </label>
                    <input
                      type="text"
                      name="title"
                      value={formData.title}
                      onChange={handleInputChange}
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                      required
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700">
                      Content
                    </label>
                    <textarea
                      name="content"
                      value={formData.content}
                      onChange={handleInputChange}
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                      rows={4}
                      required
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700">
                      Rating (1-5)
                    </label>
                    <input
                      type="number"
                      name="rating"
                      value={formData.rating}
                      onChange={handleInputChange}
                      min="1"
                      max="5"
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                      required
                    />
                  </div>
                  <div className="flex space-x-2">
                    <button
                      onClick={() => handleSubmit(wine.id!)}
                      className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                    >
                      Submit Review
                    </button>
                    <button
                      onClick={handleCloseForm}
                      className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-400"
                    >
                      Cancel
                    </button>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
