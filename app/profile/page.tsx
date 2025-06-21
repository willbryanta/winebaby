"use client";

import React, { useEffect, useState } from "react";
import NavBar from "../api/components/NavBar";
import { checkSession } from "../api/auth/services/sessionService";
import { useRouter } from "next/navigation";

export default function ProfilePage() {
  const [isAuth, setIsAuth] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const session = await checkSession();
        if (session) {
          setIsAuth(true);
        } else {
          setIsAuth(false);
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

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-screen">
        Loading...
      </div>
    );
  }

  return (
    <div>
      <NavBar isAuth={isAuth} />
      <div className="flex items-center justify-center h-screen">
        <h1 className="text-4xl font-bold">Currently still building...</h1>
      </div>
    </div>
  );
}
