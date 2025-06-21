import NavBar from "../api/components/NavBar";
import { privacyPolicy } from "../api/auth/data/privacyPolicy";
import { useEffect } from "react";
import { useRouter } from "next/router";
import { useState } from "react";
import { checkSession } from "../api/auth/services/sessionService";

export default function PrivacyPolicyPage() {
  const router = useRouter();
  const [isAuth, setIsAuth] = useState<boolean>(false);
  useEffect(() => {
    const verifySession = async () => {
      try {
        const session = await checkSession();
        if (!session) {
          setIsAuth(false);
          router.push("/signin");
        }
        setIsAuth(true);
      } catch (error) {
        console.error(`Session verification failed: ${String(error)}`);
        setIsAuth(false);
        router.push("/signin");
      }
    };
    verifySession();
  }, [router]);
  return (
    <div>
      <NavBar isAuth={isAuth} />
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white p-6 rounded shadow-md w-96">
          <h1 className="text-2xl font-bold mb-4 text-wine">Privacy Policy</h1>
          <p className="text-gray-700 mb-4">
            Welcome to our Privacy Policy page. Please read these terms
            carefully before using our service.
          </p>
          {privacyPolicy.map((term, index) => (
            <p key={index} className="text-gray-700 mb-2">
              {term}
            </p>
          ))}
        </div>
      </div>
    </div>
  );
}
