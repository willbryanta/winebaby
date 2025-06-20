"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { checkSession } from "../auth/services/sessionService";
import { signout } from "../auth/services/authService";
import { Session } from "../auth/types/page";

export default function NavBar() {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const [username, setUsername] = useState<string>("");
  const router = useRouter();

  useEffect(() => {
    checkSession().then((session: Session) => {
      setIsAuthenticated(session.isAuthenticated);
      if (session.username) {
        setUsername(session.username);
      }
    });
  }, [router]);

  const handleSignOut = async () => {
    const { success, error } = await signout();
    if (success) {
      setIsAuthenticated(false);
      router.push("/");
    } else {
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
            {isAuthenticated === null ? (
              <li>Loading...</li>
            ) : isAuthenticated ? (
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
