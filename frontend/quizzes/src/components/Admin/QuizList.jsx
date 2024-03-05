import React from "react";
import { Listbox, ListboxOption } from "@headlessui/react";

const QuizList = ({ quizzes }) => {
  return (
    <div className="max-w-md mx-auto mt-4 p-4 bg-white rounded-md shadow-md">
      <h2 className="text-xl font-semibold mb-4">Active Quizzes</h2>
      <Listbox as="ul">
        {quizzes.map((quiz) => (
          <ListboxOption key={quiz.id} value={quiz.title}>
            {({ active }) => <li className={`${active ? "bg-blue-500 text-white" : "hover:bg-blue-100"} py-2 px-4 cursor-pointer transition`}>{quiz.title}</li>}
          </ListboxOption>
        ))}
      </Listbox>
    </div>
  );
};

export default QuizList;
