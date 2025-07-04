import { ReviewResponse, ErrorWithMessage } from "../types/page";

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL as string | undefined;
const TOKEN_KEY = import.meta.env.VITE_JWT_KEY as string | undefined;

const createReview = async (
  review: Record<string, unknown>
): Promise<ReviewResponse> => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured");
    }
    const token = localStorage.getItem(TOKEN_KEY);
    const res = await fetch(`${BACKEND_URL}/reviews`, {
      method: "POST",
      headers: {
        ...(token && { Authorization: `Bearer ${token}` }),
        "Content-type": "application/json",
      },
      body: JSON.stringify(review),
    });
    if (!res.ok) {
      throw new Error(`Request failed with status${res.status}`);
    }
    const data: ReviewResponse = await res.json();
    if (data.error) {
      throw new Error(data.error);
    }
    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};

const getReview = async (review: string): Promise<ReviewResponse> => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured");
    }
    const token = localStorage.getItem(TOKEN_KEY);
    const res = await fetch(
      `${BACKEND_URL}/reviews?review=${encodeURIComponent(review)}`,
      {
        method: "GET",
        headers: {
          ...(token && { Authorization: `Bearer ${token}` }),
          "Content-type": "application/json",
        },
      }
    );
    if (!res.ok) {
      throw new Error(`Request failed with status ${res.status}`);
    }
    const data: ReviewResponse = await res.json();

    if (data.error) {
      throw new Error(data.error);
    }
    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};

const updateReview = async (
  reviewId: number,
  review: Record<string, unknown>
) => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured correctly");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured properly");
    }
    const res = await fetch(`${BACKEND_URL}/reviews/${reviewId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${TOKEN_KEY}`,
      },
      body: JSON.stringify(review),
    });
    if (!res.ok) {
      throw new Error(`Request failed with status ${res.status}`);
    }
    const data: ReviewResponse = await res.json();
    if (data.error) {
      throw new Error(data.error);
    }
    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};

const deleteReview = async (
  reviewId: number,
  review: Record<string, unknown>
) => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured");
    }
    const token = localStorage.getItem(TOKEN_KEY);
    const res = await fetch(`${BACKEND_URL}/reviews/${reviewId}`, {
      method: "DELETE",
      headers: {
        ...(token && { Authorization: `Bearer ${token}` }),
        "Content-type": "application/json",
      },
      body: JSON.stringify(review),
    });
    if (!res.ok) {
      throw new Error(`Request failed with status ${res.status}`);
    }
    const data: ReviewResponse = await res.json();
    if (data.error) {
      throw new Error(data.error);
    }
    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};

export { createReview, getReview, updateReview, deleteReview };
