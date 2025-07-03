"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { signout } from "../auth/services/authService";
import { NavBarProps } from "../auth/types/page";
import { useState } from "react";

export default function NavBar({ isAuth = false, username = "" }: NavBarProps) {
  const router = useRouter();
  const [isAuthenticated, setAuthenticated] = useState(isAuth);

  const handleSignOut = async () => {
    const { success, error } = await signout();
    if (success) {
      setAuthenticated(false);
      router.push("/");
      router.refresh();
    } else {
      setAuthenticated(true);
      console.error("Sign out failed:", error);
    }
  };

  return (
    <nav className="bg-wine dark:bg-wine-dark text-white p-4 shadow-md">
      <div className="flex items-center max-w-7xl mx-auto">
        <div className="flex-1">
          <ul className="flex space-x-6 items-center">
            <li>
              <Link
                href="/"
                className="hover:text-wine-light transition-colors duration-200"
              >
                Home
              </Link>
            </li>
            {isAuthenticated ? (
              <li>
                <Link
                  href="/dashboard"
                  className="hover:text-wine-light transition-colors duration-200 py-2 px-4"
                >
                  Dashboard
                </Link>
              </li>
            ) : (
              <>
                <li>
                  <Link
                    href="/signin"
                    className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
                  >
                    Sign in
                  </Link>
                </li>
                <li>
                  <Link
                    href="/signup"
                    className="hover:text-wine-light transition-colors duration-200 font-semibold py-2 px-4 rounded-md inline-block"
                  >
                    Sign up
                  </Link>
                </li>
              </>
            )}
          </ul>
        </div>
        <div className="flex-1 text-center">
          <span className="text-2xl font-bold text-wine-light">WineBaby</span>
        </div>
        {isAuthenticated && (
          <div className="flex-1">
            <ul className="flex space-x-6 items-center justify-end">
              <li>
                <Link
                  href={`/profile/${username}`}
                  className="hover:text-wine-light transition-colors duration-200 py-2 px-4"
                >
                  Profile
                </Link>
              </li>
              <li>
                <Link
                  href="/settings"
                  className="hover:text-wine-light transition-colors duration-200 py-2 px-4"
                >
                  Settings
                </Link>
              </li>
              <li>
                <button
                  onClick={handleSignOut}
                  className="hover:text-wine-light transition-colors duration-200 py-2 px-4"
                >
                  Sign out
                </button>
              </li>
            </ul>
          </div>
        )}
      </div>
    </nav>
  );
}
