"use client";

import { useEffect, useState } from "react";
import NavBar from "../../src/components/NavBar/NavBar";
import Link from "next/link";

interface userSettings {
  username: string;
  email: string;
}

export default function UserSettingsPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [userSettings, setUserSettings] = useState<userSettings | null>(null);

  useEffect(() => {
    const fetchUserSettings = async () => {
      try {
        const response = await fetch("http://localhost:8080/user-settings", {
          method: "GET",
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error("Failed to fetch user settings");
        }

        const data = await response.json();
        setUserSettings(data);
      } catch (error) {
        console.error("Error fetching user settings:", error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserSettings();
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!userSettings) {
    return <div>No user settings found.</div>;
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <NavBar />
      <h1 className="text-2xl font-bold mb-4">User Settings</h1>
      <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-md">
        <h2 className="text-xl font-semibold mb-2">Settings</h2>
        <p>Username: {userSettings.username}</p>
        <p>Email: {userSettings.email}</p>
        <Link href="/userProfile" className="text-blue-500 hover:underline">
          View Profile
        </Link>
      </div>
    </div>
  );
}
