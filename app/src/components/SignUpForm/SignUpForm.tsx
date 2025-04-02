"use client";

import Form from "next/form";

export default function SignUpForm() {
  return (
    <Form action="signup">
      <input name="email"></input>
      <input name="username"></input>
      <input name="password"></input>
      <input name="confirm-password"></input>
      <button type="submit">Sign up</button>
    </Form>
  );
}
