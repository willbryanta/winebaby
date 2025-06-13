"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function NavBar() {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const [username, setUsername] = useState<string>("");
  const router = useRouter();

  useEffect(() => {
    const checkSession = async () => {
      try {
        const res = await fetch("http://localhost:8080/verify-token", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          setIsAuthenticated(true);
          const data = await res.json();
          setUsername(data.username || "");
        } else {
          setIsAuthenticated(false);
          router.push("/signin");
        }
      } catch (error) {
        console.error("Error checking session:", error);
        setIsAuthenticated(false);
      }
    };
    checkSession();
  }, [router]);

  const handleSignOut = async () => {
    try {
      const res = await fetch("http://localhost:8080/signout", {
        method: "POST",
        credentials: "include",
      });

      document.cookie =
        "token=; Max-Age=0; path=/; domain=localhost; SameSite=Lax";
      if (res.ok) {
        setIsAuthenticated(false);
        router.push("/");
      } else {
        console.error("Sign out failed");
      }
    } catch (error) {
      console.error("Error signing out:", error);
    }
  };

  return (
    <nav className="bg-wine dark:bg-wine-dark text-white p-4 shadow-md">
      <ul className="flex space-x-6 items-center max-w-7xl mx-auto">
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
          <>
            <li>
              <Link
                href="/dashboard"
                className="hover:text-wine-light transition-colors duration-200 py-2 px-4"
              >
                Dashboard
              </Link>
            </li>
            <li>
              <button
                onClick={handleSignOut}
                className="bg-grape hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200"
              >
                Sign out
              </button>
            </li>
            <li>
              <Link
                href={`/users/${username}`}
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
          </>
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
    </nav>
  );
}
