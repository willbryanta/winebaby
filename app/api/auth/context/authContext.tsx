"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { checkSession } from "../services/sessionService";

interface AuthContextType {
  isAuthenticated: boolean;
  username: string;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [username, setUsername] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const verifySession = async () => {
      try {
        const session = await checkSession();
        setIsAuthenticated(!!session);
        setUsername(session?.username || "");
      } catch {
        setIsAuthenticated(false);
      } finally {
        setIsLoading(false);
      }
    };
    verifySession();
  }, []);

  return (
    <AuthContext.Provider value={{ isAuthenticated, username, isLoading }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
