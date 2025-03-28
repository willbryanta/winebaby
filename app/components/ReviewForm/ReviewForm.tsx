import Form from "next/form";

export default function ReviewForm() {
  return (
    <>
      <Form action="/review">
        <label htmlFor="Title">Title</label>
        <input name="title"></input>
        <label htmlFor="Date">Date</label>
        <input name="date"></input>
        <label htmlFor="Review">Review</label>
        <input name="review"></input>
      </Form>
    </>
  );
}
