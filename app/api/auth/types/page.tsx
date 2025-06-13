"use client";

export type Review = {
  ID: number;
  WineID?: number;
  Content: string;
  ReviewDate?: string;
  ReviewDateTime?: string;
  ReviewDateTimeUTC?: string;
  Title?: string;
  Rating?: number;
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

export interface UserProfile {
  username: string;
  email: string;
  favoriteWines: Wine[];
  reviews: Review[];
}

export interface Wine {
  id?: number;
  name: string;
  year: number;
  manufacturer?: string;
  region?: string;
  alcoholContent?: number;
  servingTemp?: number;
  servingSize?: number;
  servingSizeUnit?: string;
  price: number;
  rating?: number; // user's rating
  reviews: Review[];
  reviewCount?: number;
  averageRating: number;
  type?: string;
  color?: string;
  imageUrl?: string;
}
