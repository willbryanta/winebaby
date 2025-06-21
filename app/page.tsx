"use client";
import React from "react";
import NavBar from "./api/components/NavBar";
import { checkSession } from "./api/auth/services/sessionService";
import { useEffect, useState } from "react";

export default function Dashboard() {
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
    <div className="w-screen h-screen flex flex-col bg-gray-100">
      <NavBar isAuth={isAuth} />
      <div className="flex-1 flex flex-col items-center justify-center">
        <h1 className="text-[4.8rem] font-bold text-wine">
          Welcome to WineBaby
        </h1>
        <p className="text-[2.2rem] text-gray-700">
          Your one-stop destination for wine enthusiasts.
        </p>
      </div>
    </div>
  );
}
