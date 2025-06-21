import React, { useState } from "react";
import { Wine, Review } from "@/app/api/auth/types/page";
import { wines } from "@/app/api/auth/data/mockWineData";

const WineList: React.FC = () => {
  const [expandedWineId, setExpandedWineId] = useState<number | null>(null);

  const toggleExpand = (id: number) => {
    setExpandedWineId(expandedWineId === id ? null : id);
  };

  return (
    <div>
      {wines.map((wine: Wine) => (
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
