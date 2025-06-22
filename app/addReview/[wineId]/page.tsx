"use client";

import NavBar from "../../api/components/NavBar";
import { useParams, useRouter } from "next/navigation";
import { useState, useEffect } from "react";
import { wines } from "@/app/api/auth/data/mockWineData";
import { Wine, Review } from "../../api/auth/types/page";
import { checkSession } from "../../api/auth/services/sessionService";

export default function AddReview() {
  const { wineId } = useParams();
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const router = useRouter();

  const wine: Wine | undefined = wines.find(
    (wine) => wine.id === Number(wineId)
  );

  useEffect(() => {
    const verifySession = async () => {
      try {
        const session = await checkSession();
        if (!session) {
          router.push("/signin");
        } else {
          setIsAuthenticated(true);
        }
      } catch (error) {
        console.error("Session verification failed:", error);
        router.push("/signin");
      }
    };
    verifySession();
  }, [router]);

  useEffect(() => {
    if (success) {
      const timer = setTimeout(() => {
        router.push("/dashboard");
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [success, router]);

  const handleSubmitReview = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError(null);
    setSuccess(null);

    const formData = new FormData(event.currentTarget);
    const reviewText = formData.get("reviewText") as string;
    const rating = formData.get("rating") as string;
    const title = formData.get("title") as string;

    if (!wineId || !reviewText || !rating || !title) {
      setError("All fields are required");
      return;
    }

    if (!wine) {
      setError("Wine not found");
      return;
    }

    const parsedRating = Number(rating);
    if (isNaN(parsedRating) || parsedRating < 1 || parsedRating > 5) {
      setError("Rating must be a number between 1 and 5");
      return;
    }

    const newReview: Review = {
      ID: wines.flatMap((w) => w.reviews).length + 1,
      WineID: Number(wineId),
      Content: reviewText,
      ReviewDate: new Date().toISOString().split("T")[0],
      Title: title,
      Rating: parsedRating,
    };

    wine.reviews = [...wine.reviews, newReview];

    wine.reviewCount = wine.reviews.length;
    wine.averageRating =
      wine.reviews.reduce((sum, review) => sum + (review.Rating ?? 0), 0) /
      wine.reviewCount;

    setSuccess("Review submitted successfully!");
    event.currentTarget.reset();
  };

  if (isAuthenticated === null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <NavBar isAuth={isAuthenticated} />
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white p-6 rounded shadow-md w-96">
          <h1 className="text-2xl font-bold mb-4 text-wine">Add Review</h1>
          {error && <p className="text-red-500 mb-4">{error}</p>}
          {success && <p className="text-green-500 mb-4">{success}</p>}
          {wine ? (
            <div className="mb-4">
              <h2 className="text-lg font-semibold text-gray-700">
                {wine.name} ({wine.year})
              </h2>
            </div>
          ) : (
            <p className="text-red-500 mb-4">Wine not found</p>
          )}
          <form onSubmit={handleSubmitReview} className="space-y-4">
            <div>
              <label htmlFor="title" className="block text-gray-700">
                Review Title
              </label>
              <input
                type="text"
                name="title"
                id="title"
                required
                className="w-full p-2 border border-gray-300 rounded"
              />
            </div>
            <div>
              <label htmlFor="rating" className="block text-gray-700">
                Rating (1-5)
              </label>
              <input
                type="number"
                name="rating"
                id="rating"
                min="1"
                max="5"
                required
                className="w-full p-2 border border-gray-300 rounded"
              />
            </div>
            <div>
              <label htmlFor="reviewText" className="block text-gray-700">
                Review
              </label>
              <textarea
                name="reviewText"
                id="reviewText"
                required
                className="w-full p-2 border border-gray-300 rounded"
              ></textarea>
            </div>
            <button
              type="submit"
              className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
            >
              Submit Review
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}
