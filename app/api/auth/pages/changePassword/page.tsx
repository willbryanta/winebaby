"use client";

import { useState, useEffect } from "react";
import NavBar from "@/app/api/auth/components/NavBar";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function ChangePasswordPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    setIsLoading(false);
  }, []);

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
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <NavBar />
      <h1 className="text-2xl font-bold mb-4">Change Password</h1>
      <form
        onSubmit={handleChangePassword}
        className="bg-white shadow-md rounded-lg p-6 w-full max-w-md"
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
        {message && <p className="text-green-600">{message}</p>}
        {error && <p className="text-red-600">{error}</p>}
        <button
          type="submit"
          className="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded-md hover:bg-blue-700 transition-colors duration-200"
        >
          Change Password
        </button>
      </form>
      <Link href="/userSettings" className="mt-4 text-blue-500 hover:underline">
        Back to Settings
      </Link>
    </div>
  );
}
