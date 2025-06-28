// Please rename this file from 'userService.ts' to 'userService.tsx' to support JSX syntax.

"use client";

import { UserProfile } from "@/app/api/auth/types/page";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";

export const UserProfilePage = () => {
  const { username } = useParams();
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isAuth, setIsAuth] = useState<boolean>(true);

  useEffect(() => {
    if (!username) return;

    const fetchProfile = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/users/${username}`, {
          headers: { "Content-Type": "application/json" },
        });
        if (!res.ok) {
          throw new Error("Failed to fetch user profile");
        }
        const data: UserProfile = await res.json();
        setProfile(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "An error occurred");
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, [username]);
};
