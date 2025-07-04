"use client";

const BACKEND_URL =
  process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:3000";
const TOKEN_KEY =
  process.env.NEXT_PUBLIC_TOKEN_KEY || "default_token_key";

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
