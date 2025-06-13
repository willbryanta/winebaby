import React, { useState } from "react";
import { Wine, Review } from "@/app/api/auth/types/page";

const wines: Wine[] = [
  {
    id: 1,
    name: "Cabernet Sauvignon",
    year: 2018,
    manufacturer: "Napa Valley Vineyards",
    region: "Napa Valley",
    alcoholContent: 14.5,
    servingTemp: 18,
    servingSize: 150,
    servingSizeUnit: "ml",
    price: 45,
    rating: 4.5,
    reviews: [
      { ID: 1, Content: "Rich and full-bodied with hints of blackberry." },
      { ID: 2, Content: "A bit too dry for my taste." },
    ],
    reviewCount: 2,
    averageRating: 4.0,
    type: "red",
    color: "red",
    imageUrl: "https://example.com/cabernet.jpg",
  },
  {
    id: 2,
    name: "Chardonnay",
    year: 2020,
    manufacturer: "Sonoma Estates",
    region: "Sonoma",
    alcoholContent: 13.5,
    servingTemp: 10,
    servingSize: 150,
    servingSizeUnit: "ml",
    price: 30,
    rating: 4.2,
    reviews: [
      { ID: 3, Content: "Smooth and buttery with a touch of oak." },
      { ID: 4, Content: "Perfect for a summer evening." },
    ],
    reviewCount: 2,
    averageRating: 4.1,
    type: "white",
    color: "white",
    imageUrl: "https://example.com/chardonnay.jpg",
  },
];

const WineList: React.FC = () => {
  const [expandedWineId, setExpandedWineId] = useState<number | null>(null);

  const toggleExpand = (id: number) => {
    setExpandedWineId(expandedWineId === id ? null : id);
  };

  return (
    <div>
      {wines.map((wine) => (
        <div
          key={wine.id}
          style={{ border: "1px solid #ccc", margin: "10px", padding: "10px" }}
        >
          <h2
            onClick={() => wine.id !== undefined && toggleExpand(wine.id)}
            style={{ cursor: "pointer" }}
          >
            {wine.name}
          </h2>
          {expandedWineId === wine.id && (
            <ul>
              {wine.reviews.map((review: Review) => (
                <li key={review.ID}>{review.Content}</li>
              ))}
            </ul>
          )}
        </div>
      ))}
    </div>
  );
};

export default WineList;
