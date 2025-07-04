"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { UserProfile } from "@/app/api/auth/types/page";
import { checkSession } from "../api/auth/services/sessionService";
import { getUserProfile } from "../api/auth/services/userService";

export default function UserProfilePage() {
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProfile = async (): Promise<void> => {
    try {
      const res = await getUserProfile();
      if (!res) {
        throw new Error("Failed to fetch user profile");
      }
      setProfile(res);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const session = await checkSession();
        if (!session.isAuthenticated) {
          setError("You must be logged in to view this page.");
          setLoading(false);
          return;
        }
        if (!session.username) {
          setError("Username not found in session.");
          setLoading(false);
          return;
        }
        await fetchProfile();
      } catch (error) {
        setLoading(false);
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError("An unexpected error occurred.");
        }
      }
    };
    fetchSession();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <p className="text-lg text-gray-700">Loading...</p>
      </div>
    );
  }

  if (error || !profile) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <p className="text-lg text-red-500">{error || "User not found"}</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-4xl mx-auto px-4">
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h1 className="text-3xl font-bold text-wine mb-2">
            {profile.username}
          </h1>
          <p className="text-gray-600">{profile.email}</p>
          <Link
            href="/dashboard"
            className="mt-4 inline-block text-wine-light hover:text-wine-dark transition-colors"
          >
            Back to Dashboard
          </Link>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-2xl font-semibold text-wine mb-4">
            Favorite Wines
          </h2>
          {profile.favoriteWines.length === 0 ? (
            <p className="text-gray-500">No favorite wines yet.</p>
          ) : (
            <ul className="space-y-4">
              {profile.favoriteWines.map((wine) => (
                <li key={wine.id} className="border-b pb-2">
                  <h3 className="text-lg font-medium text-gray-800">
                    {wine.name}
                  </h3>
                  <p className="text-gray-600">{wine.region}</p>
                </li>
              ))}
            </ul>
          )}
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-2xl font-semibold text-wine mb-4">
            Wine Reviews
          </h2>
          {profile.reviews.length === 0 ? (
            <p className="text-gray-500">No reviews yet.</p>
          ) : (
            <ul className="space-y-4">
              {profile.reviews.map((review) => (
                <li key={review.ID} className="border-b pb-2">
                  <h3 className="text-lg font-medium text-gray-800">
                    {review.Title}
                  </h3>
                  <p className="text-yellow-500">
                    {"★".repeat(review.Rating ?? 0)}{" "}
                    {"☆".repeat(5 - (review.Rating ?? 0))}
                  </p>
                  <p className="text-gray-600">{review.Content}</p>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}
