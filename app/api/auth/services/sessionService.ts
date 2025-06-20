"use client";
import { verifyToken } from "./authService";

export const checkSession = async (): Promise<{
  isAuthenticated: boolean;
  username?: string;
}> => {
  try {
    const res = await verifyToken();
    if (!res) {
      return { isAuthenticated: false };
    }
    return {
      isAuthenticated: res.isAuthenticated,
      username: res.username || "",
    };
  } catch (error) {
    throw new Error("Error checking session: " + error);
  }
};
