"use client";

import { useSession, signIn, signOut } from "next-auth/react";
import Link from "next/link";

export default function NavBar() {
  // const { data: session } = useSession(); uncomment once sessioning has been implemented
  const session = null;
  return (
    <nav>
      <ul className="bg-gray-800 text-white p-4 shadow-md">
        <li>
          <Link href="/">Home</Link>
        </li>
        {session ? (
          <>
            <li>
              <Link href="/dashboard">Dashboard</Link>
            </li>
            <li>
              <button onClick={() => signOut()}>Sign out</button>
            </li>
          </>
        ) : (
          <li>
            <button onClick={() => signIn()}>Sign in</button>
          </li>
        )}
      </ul>
    </nav>
  );
}
