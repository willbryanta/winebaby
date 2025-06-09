"use client";

import { useState, FormEvent, ChangeEvent } from "react";
import { useRouter } from "next/navigation";
import { FormData } from "@/app/api/auth/types/page";

export default function ReviewForm() {
  const router = useRouter();

  const [formData, setFormData] = useState<FormData>({
    WineID: 0,
    Winemaker: "",
    Title: "",
    Description: "",
    Rating: 0,
  });

  const onFormInputChange = (
    event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;

    const updatedValue =
      name === "WineID"
        ? value === ""
          ? null
          : Number(value)
        : name === "Rating"
        ? Number(value)
        : value;
    setFormData((prevData) => ({
      ...prevData,
      [name]: updatedValue,
    }));
  };

  const isFormValid = (): boolean => {
    const { WineID, Winemaker, Title, Description, Rating } = formData;
    return (
      (WineID !== null || WineID !== 0) &&
      Winemaker.trim() !== "" &&
      Title.trim() !== "" &&
      Description.trim() !== "" &&
      Rating >= 0 &&
      Rating <= 5
    );
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!isFormValid()) {
      alert("Please fill out all fields correctly.");
      return;
    }
    try {
      // Update once the review service has been built
      const res = await fetch("/api/review", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!res.ok) throw new Error("Failed to submit review");
      alert("Review submitted successfully!");
      router.push("/");
    } catch (err) {
      console.error("Submission error:", err as unknown);
      alert("An error occured while submitting your review");
    }
  };

  //REVIEW - Revisit this, the user shouldn't be passing the title or winemaker of the wine, it should be automatically populated from the DB based on WineID
  return (
    <>
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded-lg shadow-md w-96 space-y-4"
      >
        <label htmlFor="Winemaker">Winemaker</label>
        <input
          id="Winemaker"
          name="Winemaker"
          className="input"
          value={formData.Winemaker}
          onChange={onFormInputChange}
          placeholder="Enter Winemaker"
        ></input>
        <label htmlFor="Title">Title</label>
        <input
          id="Title"
          name="Title"
          className="input"
          value={formData.Title}
          onChange={onFormInputChange}
          placeholder="Enter Title"
        ></input>
        <label htmlFor="Description">Description</label>
        <input
          id="Description"
          name="Description"
          className="input"
          value={formData.Description}
          onChange={onFormInputChange}
          placeholder="Enter Description"
        ></input>
        <label htmlFor="Rating">Rating</label>
        <input
          id="Rating"
          name="Rating"
          className="input"
          value={formData.Rating}
          onChange={onFormInputChange}
          placeholder="Enter Rating"
        ></input>
      </form>
    </>
  );
}
