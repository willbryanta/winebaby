"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";
import { UserProfile } from "@/app/api/auth/types/page";

export default function UserProfilePage() {
  const { username } = useParams();
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!username) return;

    const fetchProfile = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/users/${username}`, {
          headers: { "Content-Type": "application/json" },
        });
        if (!res.ok) {
          throw new Error("Failed to fetch user profile");
        }
        const data: UserProfile = await res.json();
        setProfile(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "An error occurred");
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, [username]);

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
