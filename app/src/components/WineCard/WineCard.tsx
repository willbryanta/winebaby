"use client";

import React from "react";
import Image from "next/image";

interface WineProperties {
  WineID: number;
  Title: string;
  Year: number;
  Winemaker: string;
  Type: string;
  Colour: string;
}

interface WineCardProps {
  wineProperties: WineProperties;
}

export default function WineCard({ wineProperties }: WineCardProps) {
  // May want to destructure properties individually i.e. Title, Manufacturer, WineID etc...
  const { WineID, Title, Year, Winemaker, Type, Colour } = wineProperties;

  return (
    <>
      <div id="wine-card">
        <Image src={`/filepathplaceholder/${WineID}`} alt={`${Title} image`} />
        <h2>Title: {Title}</h2>
        <div>Year: {Year}</div>
        <div>Winemaker: {Winemaker}</div>
        <div>Type: {Type}</div>
        <div>Colour: {Colour}</div>
      </div>
    </>
  );
}
