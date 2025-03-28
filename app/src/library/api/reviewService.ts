const BACKEND_URL = import.meta.env.VITE_EXPRESS_BACKEND_URL as
  | string
  | undefined;
const TOKEN_KEY = import.meta.env.VITE_JWT_KEY as string | undefined;

interface ReviewResponse {
  data?: string; //Might have to double check whether this is right
  error?: string;
}

interface ErrorWithMessage {
  message: string;
}

const createReview = async (
  userAgent: string,
  review: Record<string, any>
): Promise<ReviewResponse> => {
  try {
    if (!BACKEND_URL) {
      throw new Error("Backend URL is not configured");
    }
    if (!TOKEN_KEY) {
      throw new Error("Token key is not configured");
    }
    const token =
      typeof window !== "undefined" ? localStorage.getItem(TOKEN_KEY) : null;
    const res = await fetch(`${BACKEND_URL}/reviews`, {
      method: "POST",
      headers: {
        ...(token && { Authorization: `Bearer ${token}` }),
        "Content-type": "application/json",
        "User-Agent": userAgent,
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
    const token =
      typeof window !== "undefined" ? localStorage.getItem(TOKEN_KEY) : null;
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
