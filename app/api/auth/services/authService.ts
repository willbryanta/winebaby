"use client";

import { SignInResponse } from "../types/page";

export const verifyToken = async (): Promise<{
  isAuthenticated: boolean;
  username?: string;
}> => {
  try {
    const response = await fetch("/verify-token", {
      method: "GET",
      credentials: "include", // Maintains session cookies
    });
    if (!response.ok) {
      return { isAuthenticated: false };
    }
    const data = await response.json();
    if (!data || typeof data.isAuthenticated !== "boolean") {
      return { isAuthenticated: false };
    }
    return {
      isAuthenticated: data.isAuthenticated,
      username: data.username || "",
    };
  } catch (error) {
    console.error("Token verification error:", error);
    return { isAuthenticated: false };
  }
};

export const signin = async (
  username: string,
  password: string
): Promise<{ success: boolean; error?: string }> => {
  const trimmedUsername = username.trim();
  const trimmedPassword = password.trim();

  if (!trimmedUsername || !trimmedPassword) {
    return { success: false, error: "Username and password are required" };
  }

  try {
    const response = await fetch("/signin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        username: trimmedUsername,
        password: trimmedPassword,
      }),
    });

    let data: SignInResponse;
    try {
      data = await response.json();
      if (!data || !data.message) {
        return { success: false, error: "Invalid response format" };
      }
    } catch {
      return { success: false, error: "Invalid response from server" };
    }

    if (!response.ok) {
      return { success: false, error: data.message || "Something went wrong" };
    }

    return { success: true };
  } catch (error) {
    console.error("Sign-in error:", error);
    return { success: false, error: "An unexpected error occurred" };
  }
};

export const signout = async (): Promise<{
  success: boolean;
  error?: string;
}> => {
  try {
    const response = await fetch("/signout", {
      method: "POST",
      credentials: "include",
    });

    if (!response.ok) {
      return { success: false, error: "Failed to sign out" };
    }
    document.cookie =
      "token=; Max-Age=0; path=/; domain=localhost; SameSite=Lax";

    return { success: true };
  } catch (error) {
    console.error("Sign-out error:", error);
    return { success: false, error: "An unexpected error occurred" };
  }
};

export const signup = async (
  username: string,
  password: string,
  email: string
): Promise<{ success: boolean; error?: string }> => {
  const trimmedUsername = username.trim();
  const trimmedPassword = password.trim();
  const trimmedEmail = email.trim();

  if (!trimmedUsername || !trimmedPassword || !trimmedEmail) {
    return { success: false, error: "All fields are required" };
  }

  try {
    const response = await fetch("/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include", // Maintains session cookies
      body: JSON.stringify({
        username: trimmedUsername,
        password: trimmedPassword,
        email: trimmedEmail,
      }),
    });

    let data;
    try {
      data = await response.json();
      if (!data || !data.message) {
        return { success: false, error: "Invalid response format" };
      }
    } catch {
      return { success: false, error: "Invalid response from server" };
    }

    if (!response.ok) {
      return { success: false, error: data.message || "Something went wrong" };
    }

    return { success: true };
  } catch (error) {
    console.error("Sign-up error:", error);
    return { success: false, error: "An unexpected error occurred" };
  }
};

export const changePassword = async (
  oldPassword: string,
  newPassword: string
): Promise<{ success: boolean; error?: string }> => {
  if (!oldPassword || !newPassword) {
    return { success: false, error: "Both passwords are required" };
  }

  try {
    const response = await fetch("/users/profile", {
      method: "PUT",
      credentials: "include", // Maintains session cookies
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ oldPassword, password: newPassword }),
    });

    if (!response.ok) {
      return { success: false, error: "Failed to change password" };
    }

    return { success: true };
  } catch (error) {
    console.error("Change password error:", error);
    return { success: false, error: "An unexpected error occurred" };
  }
};
export const changeUsername = async (
  newUsername: string
): Promise<{ success: boolean; error?: string }> => {
  if (!newUsername) {
    return { success: false, error: "Username is required" };
  }

  try {
    const response = await fetch("/users/profile", {
      method: "PUT",
      credentials: "include", // Maintains session cookies
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username: newUsername }),
    });

    if (!response.ok) {
      return { success: false, error: "Failed to change username" };
    }

    return { success: true };
  } catch (error) {
    console.error("Change username error:", error);
    return { success: false, error: "An unexpected error occurred" };
  }
};
