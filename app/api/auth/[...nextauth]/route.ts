import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export const authOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials, req) {
        if (!credentials?.username || !credentials?.password) {
          throw new Error("Username and password are required");
        }

        // Make a POST request to your Golang backend
        const res = await fetch("http://localhost:8080/signin", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            username: credentials.username,
            password: credentials.password,
          }),
        });

        const data = await res.json();

        if (!res.ok || !data.token) {
          throw new Error(data.message || "Invalid username or password");
        }

        // Return user object with JWT token
        return {
          id: credentials.username, // Use username as ID (or fetch actual ID from backend)
          name: credentials.username,
          token: data.token, // Store the JWT token
        };
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.id = user.id;
        token.token = user.token; // Store JWT in token
      }
      return token;
    },
    async session({ session, token }) {
      session.user.id = token.id;
      session.user.token = token.token; // Pass JWT to session
      return session;
    },
  },
  pages: {
    signIn: "/signin", // Your sign-in page
  },
  secret: process.env.NEXTAUTH_SECRET, // Set in .env
};

export default NextAuth(authOptions);
