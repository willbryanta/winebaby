import React, { useState } from "react";

interface Review {
  id: number;
  content: string;
}

interface Wine {
  id: number;
  name: string;
  reviews: Review[];
}

const wines: Wine[] = [
  {
    id: 1,
    name: "Cabernet Sauvignon",
    reviews: [
      { id: 1, content: "Rich and full-bodied with hints of blackberry." },
      { id: 2, content: "A bit too dry for my taste." },
    ],
  },
  {
    id: 2,
    name: "Chardonnay",
    reviews: [
      { id: 3, content: "Smooth and buttery with a touch of oak." },
      { id: 4, content: "Perfect for a summer evening." },
    ],
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
