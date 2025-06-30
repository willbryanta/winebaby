"use client";

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

export interface ReviewResponse {
  data?: string;
  error?: string;
}

export interface ErrorWithMessage {
  message: string;
}

export type ReviewCardType = {
  review: Review;
};

export type NavBarProps = {
  isAuth: boolean;
  username?: string;
};

export interface ReviewCardProps {
  reviews: Review[] | undefined;
  wineName: string;
  wineType: string;
}

export type Session = {
  isAuthenticated: boolean;
  username?: string;
};

export interface FetchProfileParams {
  username: string;
}

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
