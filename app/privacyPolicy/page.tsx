import NavBar from "../api/components/NavBar";

const privacyPolicy: string[] = [
  "1. Introduction: This Privacy Policy explains how Winebaby ('we,' 'us,' or 'our') collects, uses, and protects your personal information when you use our service.",
  "2. Information We Collect: We may collect personal information such as your name, email address, and payment details when you register, make purchases, or interact with our service. We also collect usage data, such as IP addresses and browsing behavior, through cookies and similar technologies.",
  "3. How We Use Your Information: Your information is used to provide and improve our service, process transactions, send communications, and personalize your experience. We may also use it for analytics and to comply with legal obligations.",
  "4. Information Sharing: We do not sell your personal information. We may share it with trusted third-party service providers (e.g., payment processors) to operate our service, or as required by law.",
  "5. Your Rights: You have the right to access, correct, or delete your personal information, subject to applicable laws. You may also opt out of marketing communications at any time.",
  "6. Data Security: We implement reasonable security measures to protect your information, but no system is completely secure. You use our service at your own risk.",
  "7. Cookies: We use cookies to enhance your experience. You can manage cookie preferences through your browser settings.",
  "8. Changes to This Policy: We may update this Privacy Policy from time to time. Changes will be posted on our website, and your continued use of the service constitutes acceptance of the updated policy.",
  "9. Contact Us: For questions about this Privacy Policy, contact us at support@winebaby.com.",
];

export default function PrivacyPolicyPage() {
  return (
    <div>
      <NavBar />
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
