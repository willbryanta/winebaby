"use client";

import { useSession, signIn, signOut } from "next-auth/react";
import Link from "next/link";

export default function NavBar() {
  // const { data: session } = useSession(); uncomment once sessioning has been implemented
  const session = null;
  return (
    <nav className="bg-gray-800 p-4 shadow-md">
      <ul className="flex space-x-6 text-white">
        <li>
          <Link href="/" className="hover:text-gray-300">
            Home
          </Link>
        </li>
        {session ? (
          <>
            <li>
              <Link href="/dashboard" className="hover:text-gray-300">
                Dashboard
              </Link>
            </li>
            <li>
              <button onClick={() => signOut()} className="btn">
                Sign out
              </button>
            </li>
          </>
        ) : (
          <li>
            <button onClick={() => signIn()} className="btn">
              Sign in
            </button>
          </li>
        )}
      </ul>
    </nav>
  );
}
