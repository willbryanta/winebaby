"use client";

import { Wine, WineCardProps, Review } from "../auth/types/page";
import Link from "next/link";

export default function WineCard({ wines }: WineCardProps) {
  return (
    <div className="flex-1 flex flex-col items-center overflow-y-auto py-6">
      <h1 className="text-2xl font-bold mb-4">Wine Dashboard</h1>
      <div className="w-full max-w-4xl bg-white shadow-md rounded-lg p-6 mx-4">
        {wines.map((wine: Wine) => (
          <div
            key={wine.id}
            className="mb-4 p-4 border-b last:border-b-0 flex flex-col md:flex-row gap-4"
          >
            <div className="flex-1">
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
            </div>
            <div className="md:w-48 flex-shrink-0 flex flex-col items-center">
              <img
                src={wine.imageUrl || "/placeholder-wine-bottle.png"}
                alt={`${wine.name} bottle`}
                className="w-full h-auto object-contain rounded-md"
                onError={(e) => {
                  e.currentTarget.src = "/placeholder-wine-bottle.png";
                }}
              />
              <Link href={`/addReview/${wine.id}`}>
                <button className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block">
                  Add Review
                </button>
              </Link>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
