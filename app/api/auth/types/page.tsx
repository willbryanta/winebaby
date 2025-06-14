"use client";

//TODO - update optional fields based on UI requirements
export type Review = {
  ID: number;
  UserID?: number;
  WineID?: number;
  Content: string;
  ReviewDate?: string;
  ReviewDateTime?: string;
  Title?: string;
  Rating?: number;
};

export type ReviewCardType = {
  review: Review;
};

export type ReviewCardProps = {
  reviews: Review[];
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
  wines: Wine[];
  reviews?: Review[];
}

export interface UserProfile {
  username: string;
  email: string;
  favoriteWines: Wine[];
  reviews: Review[];
}

export type Wine = {
  id?: number;
  name: string;
  year: number;
  manufacturer?: string;
  region?: string;
  alcoholContent?: number;
  price: number;
  rating?: number; // user's rating
  reviews: Review[];
  reviewCount?: number;
  averageRating: number | null;
  type?: string;
  color?: string;
  imageUrl?: string;
};

export interface SignInResponse {
  message: string;
}

export interface SignInError {
  message: string;
}

export interface SignUpResponse {
  message: string;
}
