import React from "react";
import { useParams } from "react-router-dom";

const QuizPage = () => {
  const { quizId } = useParams();

  const quiz = {
    id: quizId,
    title: "Sample Quiz",
    questions: [
      {
        id: 1,
        text: "What is the capital of France?",
        options: ["Berlin", "Madrid", "Paris", "Rome"],
        correctAnswer: "Paris",
      },
    ],
  };

  const submitAnswers = (answers) => {
    console.log("Submitted Answers:", answers);
  };

  return (
    <div className="max-w-md mx-auto mt-8 p-4 bg-white rounded-md shadow-md">
      <h2 className="text-2xl font-semibold mb-4">{quiz.title}</h2>

      <form onSubmit={(e) => e.preventDefault()}>
        {quiz.questions.map((question) => (
          <div key={question.id} className="mb-4">
            <p className="text-lg font-medium mb-2">{question.text}</p>
            <div className="flex flex-col space-y-2">
              {question.options.map((option, index) => (
                <label key={index} className="flex items-center">
                  <input type="radio" name={`question_${question.id}`} value={option} className="mr-2" />
                  {option}
                </label>
              ))}
            </div>
          </div>
        ))}

        <button type="submit" onClick={() => submitAnswers({ quizId, answers: [] })} className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-700">
          Submit Answers
        </button>
      </form>
    </div>
  );
};

export default QuizPage;
