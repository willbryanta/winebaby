import { useSession, signIn, signOut } from "next-auth/react";
import Link from "next/link";

export default function NavBar() {
  const { data: session } = useSession();
  return (
    <>
      <nav>
        <ul>
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
    </>
  );
}
