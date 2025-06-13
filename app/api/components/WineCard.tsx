"use client";

import React from "react";
import Image from "next/image";
import { WineCardProps } from "@/app/api/auth/types/page";

export default function WineCard({ wineProperties }: WineCardProps) {
  const { WineID, Title, Year, Winemaker, Type, Colour } = wineProperties;

  return (
    <div
      id="wine-card"
      className="bg-white rounded-lg shadow-lg p-4 w-80 border border-gray-200"
    >
      <Image
        src={`/filepathplaceholder/${WineID}`}
        alt={`${Title} image`}
        width={300}
        height={200}
        className="rounded-md"
      />
      <h2 className="text-lg font-semibold mt-2">Title: {Title}</h2>
      <p className="text-gray-600">Year: {Year}</p>
      <p className="text-gray-600">Winemaker: {Winemaker}</p>
      <p className="text-gray-600">Type: {Type}</p>
      <p className="text-gray-600">Colour: {Colour}</p>
    </div>
  );
}
