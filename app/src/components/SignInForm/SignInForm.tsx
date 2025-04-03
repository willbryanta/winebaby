"use client";

import Form from "next/form";

export default function SignInForm() {
  return (
    <Form
      action="/signin"
      className="bg-white p-6 rounded-lg shadow-md w-80 space-y-4"
    >
      <label htmlFor="username" className="block text-gray-700">
        Username
      </label>
      <input name="username" placeholder="Username"></input>
      <label htmlFor="password" className="block text-gray-700">
        Password
      </label>
      <input name="password" placeholder="Password"></input>
      <button type="submit" className="btn">
        Sign in
      </button>
    </Form>
  );
}
