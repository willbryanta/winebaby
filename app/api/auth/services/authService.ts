"use client";

export interface SignInResponse {
  message: string;
}

export interface SignInError {
  message: string;
}

export interface SignUpResponse {
  message: string;
}

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
      credentials: "include", // Maintains session cookies
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
      credentials: "include", // Maintains session cookies
    });

    if (!response.ok) {
      return { success: false, error: "Failed to sign out" };
    }

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
