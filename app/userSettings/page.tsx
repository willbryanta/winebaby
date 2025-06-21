"use client";

import { useState, useEffect } from "react";
import NavBar from "@/app/api/components/NavBar";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { handleProfileUpdate } from "@/app/api/auth/utils/authUtils";
import { checkSession } from "../api/auth/services/sessionService";

export default function UpdateUserCredentials() {
  const [isLoading, setIsLoading] = useState(true);
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isAuth, setIsAuth] = useState<boolean>(true);
  const router = useRouter();

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const session = await checkSession();
        if (session) {
          setIsAuth(true);
        } else {
          setIsAuth(false);
          router.push("/signin");
        }
      } catch (error) {
        console.error("Error checking session:", error);
        setIsAuth(false);
        router.push("/signin");
      } finally {
        setIsLoading(false);
      }
    };
    fetchSession();
  }, [router]);

  const handleChangeEmail = (event: React.FormEvent<HTMLFormElement>) =>
    handleProfileUpdate(
      event,
      (formData) => ({ email: formData.get("newEmail") as string }),
      "Email changed successfully!",
      setMessage,
      setError,
      router
    );

  const handleChangeUsername = (event: React.FormEvent<HTMLFormElement>) =>
    handleProfileUpdate(
      event,
      (formData) => ({ username: formData.get("newUsername") as string }),
      "Username changed successfully!",
      setMessage,
      setError,
      router
    );

  const handleChangePassword = (event: React.FormEvent<HTMLFormElement>) =>
    handleProfileUpdate(
      event,
      (formData) => ({
        oldPassword: formData.get("oldPassword") as string,
        password: formData.get("newPassword") as string,
      }),
      "Password changed successfully!",
      setMessage,
      setError,
      router
    );

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <NavBar isAuth={isAuth} />
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
              <div className="flex justify-center">
                <button
                  type="submit"
                  className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
                >
                  Change Email
                </button>
              </div>
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
              <div className="flex justify-center">
                <button
                  type="submit"
                  className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
                >
                  Change Username
                </button>
              </div>
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
              <div className="flex justify-center">
                <button
                  type="submit"
                  className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
                >
                  Change Password
                </button>
              </div>
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
