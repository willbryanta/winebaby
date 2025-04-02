"use client";

import Form from "next/form";

export default function SignInForm() {
  return (
    <Form action="/signin">
      <label htmlFor="username">Username</label>
      <input name="username"></input>
      <label htmlFor="password">Password</label>
      <input name="password"></input>
      <button type="submit">Sign in</button>
    </Form>
  );
}
