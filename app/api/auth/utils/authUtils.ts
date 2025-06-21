import { useRouter } from "next/navigation";

export async function handleProfileUpdate(
  event: React.FormEvent<HTMLFormElement>,
  buildPayload: (formData: FormData) => Record<string, unknown>,
  successMessage: string,
  setMessage: (msg: string | null) => void,
  setError: (msg: string | null) => void,
  router: ReturnType<typeof useRouter>
) {
  event.preventDefault();
  const formData = new FormData(event.currentTarget);
  const payload = buildPayload(formData);

  try {
    const response = await fetch("http://localhost:8080/api/users/profile", {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) throw new Error("Failed to update profile");

    setMessage(successMessage);
    setError(null);
    router.push("/userSettings");
  } catch (error) {
    console.error("Update error:", error);
    setError(
      error instanceof Error ? error.message : "An unexpected error occurred"
    );
    setMessage(null);
  }
}
