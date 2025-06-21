"use client";

import Link from "next/link";
import NavBar from "../api/components/NavBar";
import { checkSession } from "../api/auth/services/sessionService";
import { useEffect, useState } from "react";

export default function SettingsPage() {
  const [isAuth, setIsAuth] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const verifySession = async () => {
      try {
        const session = await checkSession();
        setIsAuth(!!session);
      } catch {
        setIsAuth(false);
      } finally {
        setIsLoading(false);
      }
    };
    verifySession();
  }, []);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-screen">
        Loading...
      </div>
    );
  }
  if (!isAuth) {
    return (
      <p className="flex items-center justify-center h-screen">Access Denied</p>
    );
  }
  return (
    <div>
      <NavBar isAuth={isAuth} />
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white p-6 rounded shadow-md w-96">
          <h1 className="text-2xl font-bold mb-4 text-wine">Settings</h1>
          <ul className="space-y-2 list-none">
            <li>
              <a href="/userSettings" className="text-blue-500 hover:underline">
                User Settings
              </a>
            </li>
            <li>
              <Link href="/reviews" className="text-blue-500 hover:underline">
                Reviews
              </Link>
            </li>
            <li>
              <Link
                href="/privacyPolicy"
                className="text-blue-500 hover:underline"
              >
                Privacy Policy
              </Link>
            </li>
            <li>
              <Link
                href="/termsAndConditions"
                className="text-blue-500 hover:underline"
              >
                Terms and Conditions
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
