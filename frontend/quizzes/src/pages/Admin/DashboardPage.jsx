import React from "react";
import QuizList from "../components/Admin/QuizList";
import QuizForm from "../components/Admin/QuizForm";
import { Transition } from "@headlessui/react";

const DashboardPage = ({ quizzes, onCreateQuiz }) => {
  return (
    <div className="max-w-3xl mx-auto mt-8 p-4 bg-white rounded-md shadow-md">
      <h2 className="text-2xl font-semibold mb-4">Admin Dashboard</h2>

      <Transition show={true} enter="transition-opacity duration-300" enterFrom="opacity-0" enterTo="opacity-100" leave="transition-opacity duration-300" leaveFrom="opacity-100" leaveTo="opacity-0">
        <QuizForm onSubmit={onCreateQuiz} />
      </Transition>

      <div className="mt-4">
        <QuizList quizzes={quizzes} />
      </div>
    </div>
  );
};

export default DashboardPage;
