"use client";

import React, { useEffect, useState } from "react";
import NavBar from "../api/components/NavBar";
import { wines } from "../api/auth/data/mockWineData";
import ReviewCard from "../api/components/ReviewCard";
import { checkSession } from "../api/auth/services/sessionService";
import { useRouter } from "next/navigation";

export default function ReviewsPage() {
  const [isAuth, setIsAuth] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const session = await checkSession();
        if (session) {
          setIsAuth(true);
        } else {
          setIsAuth(false);
          router.push("/signin");
        }
      } catch (error) {
        console.error("Error checking session:", error);
        setIsAuth(false);
        router.push("/signin");
      }
    };
    fetchSession();
  }, [router]);

  return (
    <div className="flex flex-col w-screen h-screen bg-gray-100 overflow-hidden">
      <NavBar isAuth={isAuth} />
      <div className="flex-1 overflow-y-auto py-6">
        {wines.map((wine) => (
          <ReviewCard
            key={wine.id}
            reviews={wine.reviews}
            wineName={wine.name}
            wineType={wine.type ?? ""}
          />
        ))}
      </div>
    </div>
  );
}
