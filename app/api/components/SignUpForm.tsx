"use client";

import Form from "next/form";

export default function SignUpForm() {
  return (
    <Form
      action="signup"
      className="bg-white p-6 rounded-lg shadow-md w-80 space-y-4"
    >
      <input name="email" placeholder="Email"></input>
      <input name="username" placeholder="Username"></input>
      <input name="password" placeholder="Password"></input>
      <input name="confirm-password" placeholder="Confirm Password"></input>
      <button type="submit" className="btn">
        Sign up
      </button>
    </Form>
  );
}
