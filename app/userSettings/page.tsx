"use client";

import { useState, useEffect } from "react";
import NavBar from "@/app/api/components/NavBar";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function ChangeEmailPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    setIsLoading(false);
  }, []);

  const handleChangeEmail = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const newEmail = formData.get("newEmail") as string;

    try {
      const response = await fetch("http://localhost:8080/api/users/profile", {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: newEmail }),
      });

      if (!response.ok) {
        throw new Error("Failed to change email");
      }

      setMessage("Email changed successfully!");
      setError(null);
      router.push("/userSettings");
    } catch (error) {
      console.error("Error changing email:", error);
      setError(
        error instanceof Error ? error.message : "An unexpected error occurred"
      );
      setMessage(null);
    }
  };

  const handleChangeUsername = async (
    event: React.FormEvent<HTMLFormElement>
  ) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const newUsername = formData.get("newUsername") as string;

    try {
      const response = await fetch("http://localhost:8080/api/users/profile", {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username: newUsername }),
      });

      if (!response.ok) {
        throw new Error("Failed to change username");
      }

      setMessage("Username changed successfully!");
      setError(null);
      router.push("/userSettings");
    } catch (error) {
      console.error("Error changing username:", error);
      setError(
        error instanceof Error ? error.message : "An unexpected error occurred"
      );
      setMessage(null);
    }
  };

  const handleChangePassword = async (
    event: React.FormEvent<HTMLFormElement>
  ) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const oldPassword = formData.get("oldPassword") as string;
    const newPassword = formData.get("newPassword") as string;

    try {
      const response = await fetch("http://localhost:8080/api/users/profile", {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ oldPassword, password: newPassword }),
      });

      if (!response.ok) {
        throw new Error("Failed to change password");
      }

      setMessage("Password changed successfully!");
      setError(null);
      router.push("/userSettings");
    } catch (error) {
      console.error("Error changing password:", error);
      setError(
        error instanceof Error ? error.message : "An unexpected error occurred"
      );
      setMessage(null);
    }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <NavBar />
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 py-8">
        <div className="space-y-10 w-full max-w-md">
          <div>
            <h1 className="text-2xl font-bold mb-4 text-center">
              Change Email
            </h1>
            <form
              onSubmit={handleChangeEmail}
              className="bg-white shadow-md rounded-lg p-6"
            >
              <div className="mb-4">
                <label
                  htmlFor="newEmail"
                  className="block text-sm font-medium text-gray-700"
                >
                  New Email
                </label>
                <input
                  type="email"
                  name="newEmail"
                  id="newEmail"
                  required
                  className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              {message && <p className="text-green-600 mb-4">{message}</p>}
              {error && <p className="text-red-600 mb-4">{error}</p>}
              <button
                type="submit"
                className="w-full bg-blue-600 text-white font-medium py-1.5 px-3 rounded-md hover:bg-blue-700 transition-colors duration-200 text-sm"
              >
                Change Email
              </button>
            </form>
          </div>
          <div>
            <h1 className="text-2xl font-bold mb-4 text-center">
              Change Username
            </h1>
            <form
              onSubmit={handleChangeUsername}
              className="bg-white shadow-md rounded-lg p-6"
            >
              <div className="mb-4">
                <label
                  htmlFor="newUsername"
                  className="block text-sm font-medium text-gray-700"
                >
                  New Username
                </label>
                <input
                  type="text"
                  id="newUsername"
                  name="newUsername"
                  required
                  className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              {message && <p className="text-green-600 mb-4">{message}</p>}
              {error && <p className="text-red-600 mb-4">{error}</p>}
              <button
                type="submit"
                className="w-full bg-blue-600 text-white font-medium py-1.5 px-3 rounded-md hover:bg-blue-700 transition-colors duration-200 text-sm"
              >
                Change Username
              </button>
            </form>
          </div>
          <div>
            <h1 className="text-2xl font-bold mb-4 text-center">
              Change Password
            </h1>
            <form
              onSubmit={handleChangePassword}
              className="bg-white shadow-md rounded-lg p-6"
            >
              <div className="mb-4">
                <label
                  htmlFor="oldPassword"
                  className="block text-sm font-medium text-gray-700"
                >
                  Old Password
                </label>
                <input
                  type="password"
                  id="oldPassword"
                  name="oldPassword"
                  required
                  className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="newPassword"
                  className="block text-sm font-medium text-gray-700"
                >
                  New Password
                </label>
                <input
                  type="password"
                  id="newPassword"
                  name="newPassword"
                  required
                  className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              {message && <p className="text-green-600 mb-4">{message}</p>}
              {error && <p className="text-red-600 mb-4">{error}</p>}
              <button
                type="submit"
                className="w-full bg-blue-600 text-white font-medium py-1.5 px-3 rounded-md hover:bg-blue-700 transition-colors duration-200 text-sm"
              >
                Change Password
              </button>
            </form>
          </div>
          <div className="text-center">
            <Link
              href="/settings"
              className="text-blue-500 hover:underline text-sm"
            >
              Back to Settings
            </Link>
          </div>
        </div>
      </div>
    </>
  );
}
