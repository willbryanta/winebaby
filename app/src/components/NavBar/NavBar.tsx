"use client";

import { useSession, signOut } from "next-auth/react";
import Link from "next/link";

export default function NavBar() {
  const { data: session, status } = useSession();

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
        {status === "loading" ? (
          <li>Loading...</li>
        ) : session ? (
          <>
            <li>
              <Link
                href="/dashboard"
                className="hover:text-wine-light transition-colors duration-200"
              >
                Dashboard
              </Link>
            </li>
            <li>
              <button
                onClick={() => signOut({ callbackUrl: "/" })}
                className="bg-grape hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200"
              >
                Sign out
              </button>
            </li>
          </>
        ) : (
          <li>
            <Link
              href="/signin"
              className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
            >
              Sign in
            </Link>
            <Link
              href="/signup"
              className="bg-wine hover:bg-wine-dark text-white font-semibold py-2 px-4 rounded-md transition-colors duration-200 inline-block"
            >
              Sign up
            </Link>
          </li>
        )}
      </ul>
    </nav>
  );
}
