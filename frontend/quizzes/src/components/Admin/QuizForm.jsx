import React, { useState } from "react";

const QuizForm = ({ onSubmit }) => {
  const [title, setTitle] = useState("");

  const handleSubmit = () => {
    onSubmit({ title });
  };

  return (
    <div className="max-w-md mx-auto mt-4 p-4 bg-white rounded-md shadow-md">
      <input type="text" placeholder="Quiz Title" value={title} onChange={(e) => setTitle(e.target.value)} className="w-full border border-gray-300 p-2 mb-2 rounded-md focus:outline-none focus:ring focus:border-blue-500" />
      <button onClick={handleSubmit} className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-700">
        Create Quiz
      </button>
    </div>
  );
};

export default QuizForm;
