"use client";

import Link from "next/link";

export default function SettingsPage() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-6 rounded shadow-md w-96">
        <h1 className="text-2xl font-bold mb-4 text-wine">Settings</h1>
        <li>
          <a href="/userProfile" className="text-blue-500 hover:underline">
            User Profile
          </a>
        </li>
        <li>
          <Link href="/reviews" className="text-blue-500 hover:underline">
            Reviews
          </Link>
        </li>
        <li>
          <Link href="/privacy" className="text-blue-500 hover:underline">
            Privacy Policy
          </Link>
        </li>
        <li>
          <Link
            href="/termsAndConditions"
            className="text-blue-500 hover:underline"
          >
            Terms and Conditions
          </Link>
        </li>
      </div>
    </div>
  );
}
