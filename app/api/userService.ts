import { User } from "next-auth";

interface ProfileResponse {
  data?: string; //Might have to double check whether this is right
  error?: string;
}

interface ErrorWithMessage {
  message: string;
}

const BACKEND_URL = import.meta.env.VITE_EXPRESS_BACKEND_URL as
  | string
  | undefined;
const TOKEN_KEY = import.meta.env.VITE_JWT_KEY as string | undefined;

const getProfile = async (user: User): Promise<ProfileResponse> => {
  if (!BACKEND_URL) {
    throw new Error("Backend URL is not configured properly");
  }
  if (!TOKEN_KEY) {
    throw new Error("Backend URL is not configured properly");
  }

  const token =
    typeof window !== "undefined" ? localStorage.getItem(TOKEN_KEY) : null;
  try {
    const res = await fetch(`${BACKEND_URL}/users/${user._id}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(TOKEN_KEY)}`,
      },
      user,
    });
    if (!res.ok) {
      throw new Error(res.error);
    }
    const data: ProfileResponse = await res.json();

    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};
