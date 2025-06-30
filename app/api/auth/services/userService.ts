"use client";

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL as string | undefined;
const TOKEN_KEY = import.meta.env.VITE_JWT_KEY as string | undefined;

import { UserProfile } from "@/app/api/auth/types/page";

export const getUserProfile = async (): Promise<UserProfile | null> => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured");
    }
    const token = localStorage.getItem(TOKEN_KEY);
    const res = await fetch(`${BACKEND_URL}/api/users/profile`, {
      method: "GET",
      headers: {
        ...(token && { Authorization: `Bearer ${token}` }),
        "Content-type": "application/json",
      },
    });
    if (!res.ok) {
      throw new Error("Failed to fetch user profile");
    }
    const data: UserProfile = await res.json();
    return data;
  } catch (error) {
    console.error("Error fetching user profile:", error);
    return null;
  }
};
