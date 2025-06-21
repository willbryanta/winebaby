"use client";

import React, { useEffect, useState } from "react";
import NavBar from "@/app/api/components/NavBar";
import WineCard from "../api/components/WineCard";
import { wines } from "../api/auth/data/mockWineData";
import { checkSession } from "../api/auth/services/sessionService";

const Dashboard: React.FC = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const verifySession = async () => {
      try {
        const session = await checkSession();
        setIsAuthenticated(!!session);
      } catch {
        setIsAuthenticated(false);
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
  if (!isAuthenticated) {
    return (
      <p className="flex items-center justify-center h-screen">Access Denied</p>
    );
  }

  return (
    <div className="flex flex-col w-screen h-screen bg-gray-100 overflow-hidden">
      <NavBar isAuth={isAuthenticated} />
      <WineCard wines={wines} />
    </div>
  );
};

export default Dashboard;
