"use client";

import { UserProfile } from "@/app/api/auth/types/page";

export const getUserProfile = async (): Promise<UserProfile | null> => {
  try {
    const res = await fetch("http://localhost:8080/api/users/profile", {
      headers: { "Content-Type": "application/json" },
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
