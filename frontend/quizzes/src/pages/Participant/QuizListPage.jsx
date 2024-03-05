import React from "react";
import { Link } from "react-router-dom";

const QuizListPage = ({ quizzes }) => {
  return (
    <div className="max-w-md mx-auto mt-8 p-4 bg-white rounded-md shadow-md">
      <h2 className="text-2xl font-semibold mb-4">Available Quizzes</h2>
      <ul>
        {quizzes.map((quiz) => (
          <li key={quiz.id} className="mb-2">
            <Link to={`/participant/quiz/${quiz.id}`} className="block py-2 px-4 border border-gray-300 rounded-md hover:bg-gray-100 transition">
              {quiz.title}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default QuizListPage;
