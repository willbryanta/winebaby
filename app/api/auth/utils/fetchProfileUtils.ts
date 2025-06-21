"use client";

import { UserProfile } from "@/app/api/auth/types/page";

export const fetchProfile = async (username: string) => {
  try {
    const response = await fetch(
      `http://localhost:8080/api/users/${username}`,
      {
        headers: { "Content-Type": "application/json" },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch user profile");
    }

    const data: UserProfile = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching profile:", error);
    throw error;
  }
};
