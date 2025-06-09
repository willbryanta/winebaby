"use client";

export type Review = {
  ID: number;
  WineID: number;
  Comment: string;
  ReviewDate: string;
  ReviewDateTime: string;
  ReviewDateTimeUTC: string;
  Title: string;
  Description: string;
  Rating: number;
};

export interface FormData {
  WineID: number | null;
  Winemaker: string;
  Title: string;
  Description: string;
  Rating: number;
}

export interface WineProperties {
  WineID: number;
  Title: string;
  Year: number;
  Winemaker: string;
  Type: string;
  Colour: string;
}

export interface WineCardProps {
  wineProperties: WineProperties;
}

export interface Wine {
  id: string;
  name: string;
  region: string;
}

export interface UserProfile {
  username: string;
  email: string;
  favoriteWines: Wine[];
  reviews: Review[];
}
