"use client";

import { useState, useEffect } from "react";
import NavBar from "@/app/src/components/NavBar/NavBar";
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
      const response = await fetch("http://localhost:8080/change-email", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ newEmail }),
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

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <NavBar />
      <h1 className="text-2xl font-bold mb-4">Change Email</h1>
      <form
        onSubmit={handleChangeEmail}
        className="bg-white shadow-md rounded-lg p-6 w-full max-w-md"
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
        {message && <p className="text-green-600">{message}</p>}
        {error && <p className="text-red-600">{error}</p>}
        <button
          type="submit"
          className="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded hover:bg-blue-700 transition-colors duration-200"
        >
          Change Email
        </button>
      </form>
    </div>
  );
}
