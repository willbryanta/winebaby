"use client";

import React from "react";
import NavBar from "@/app/api/components/NavBar";
import WineCard from "../api/components/WineCard";
import { wines } from "../api/auth/data/mockWineData";
import { useAuth } from "@/app/context/AuthContext";

const Dashboard: React.FC = () => {
  const { isAuthenticated, isLoading } = useAuth();

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
      <NavBar />
      <WineCard wines={wines} />
    </div>
  );
};

export default Dashboard;
