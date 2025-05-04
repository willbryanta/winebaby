import NextAuth from "next-auth";
import { authOptions } from "./auth";

// Initialize the NextAuth handler
const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
export default NextAuth(authOptions);
