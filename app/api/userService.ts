type User = {
  ID: number;
  Username: string;
  Email: string;
  Password: string;
};

interface ProfileResponse {
  data?: string;
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

  const token = localStorage.getItem(TOKEN_KEY);
  try {
    const res = await fetch(`${BACKEND_URL}/users/${user.ID}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status} ${res.statusText}`);
    }
    const data: ProfileResponse = await res.json();

    return data;
  } catch (error) {
    const err = error as ErrorWithMessage;
    return { error: err.message };
  }
};

export { getProfile };
