import NavBar from "../api/components/NavBar";

export default function TermsAndConditions() {
  return (
    <div>
      <NavBar />
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-3xl">
          <h1 className="text-2xl font-bold mb-4">Terms and Conditions</h1>
          <p className="mb-4">
            Welcome to our website! These terms and conditions outline the rules
            and regulations for the use of our website.
          </p>
          <h2 className="text-xl font-semibold mb-2">1. Acceptance of Terms</h2>
          <p className="mb-4">
            By accessing this website, you accept these terms and conditions in
            full. If you disagree with any part of these terms, you must not use
            this website.
          </p>
          <h2 className="text-xl font-semibold mb-2">2. Changes to Terms</h2>
          <p className="mb-4">
            We reserve the right to modify these terms at any time. Your
            continued use of the website after changes are made constitutes your
            acceptance of the new terms.
          </p>
          <h2 className="text-xl font-semibold mb-2">
            3. User Responsibilities
          </h2>
          <p className="mb-4">
            You agree to use this website only for lawful purposes and in a
            manner that does not infringe on the rights of others.
          </p>
          <h2 className="text-xl font-semibold mb-2">
            4. Limitation of Liability
          </h2>
          <p className="mb-4">
            We will not be liable for any damages arising from your use of this
            website or its content.
          </p>
        </div>
      </div>
    </div>
  );
}
