"use client";

import NavBar from "../api/components/NavBar";

export default function AddReview() {
  const handleSubmitReview = async (
    event: React.FormEvent<HTMLFormElement>
  ) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const wineName = formData.get("wineName") as string;
    const reviewText = formData.get("reviewText") as string;
    try {
      const response = await fetch("/api/reviews", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ wineName, reviewText }),
      });
      if (!response.ok) {
        throw new Error("Failed to submit review");
      }
      const result = await response.json();
      console.log("Review submitted successfully:", result);
    } catch (error) {
      console.error("Error submitting review:", error);
    }
  };
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <NavBar />
      <div className="bg-white p-6 rounded shadow-md w-96">
        <h1 className="text-2xl font-bold mb-4 text-wine">Add Review</h1>
        <form onSubmit={handleSubmitReview} className="space-y-4">
          <label htmlFor="wineName" className="block text-gray-700">
            Wine Name
          </label>
          <input
            type="text"
            name="wineName"
            id="wineName"
            required
            className="w-full p-2 border border-gray-300 rounded"
          />
          <label htmlFor="reviewText" className="block text-gray-700">
            Review
          </label>
          <textarea
            name="reviewText"
            id="reviewText"
            required
            className="w-full p-2 border border-gray-300 rounded"
          ></textarea>
          <button type="submit" className="btn">
            Submit Review
          </button>
        </form>
      </div>
    </div>
  );
}
