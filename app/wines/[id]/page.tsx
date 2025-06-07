import React, { useState } from "react";

interface Review {
  id: number;
  content: string;
}

interface Wine {
  id: number;
  name: string;
  year: number;
  manufacturer: string;
  region: string;
  alcoholContent: number;
  servingTemp: number;
  servingSize: number;
  servingSizeUnit: string;
  price: number;
  rating?: number; // user's rating
  reviews: Review[];
  reviewCount: number;
  averageRating: number;
  type: string; // e.g., "red", "white", "sparkling"
  color: string; // e.g., "red", "white", "rose"
  imageUrl?: string; // optional image URL
}

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
      { id: 1, content: "Rich and full-bodied with hints of blackberry." },
      { id: 2, content: "A bit too dry for my taste." },
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
      { id: 3, content: "Smooth and buttery with a touch of oak." },
      { id: 4, content: "Perfect for a summer evening." },
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
            onClick={() => toggleExpand(wine.id)}
            style={{ cursor: "pointer" }}
          >
            {wine.name}
          </h2>
          {expandedWineId === wine.id && (
            <ul>
              {wine.reviews.map((review) => (
                <li key={review.id}>{review.content}</li>
              ))}
            </ul>
          )}
        </div>
      ))}
    </div>
  );
};

export default WineList;
