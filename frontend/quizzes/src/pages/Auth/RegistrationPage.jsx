import React from "react";
import RegistrationForm from "../components/Auth/RegistrationForm";
import { Transition } from "@headlessui/react";

const RegistrationPage = ({ onRegister }) => {
  return (
    <div className="max-w-md mx-auto mt-8 p-4 bg-white rounded-md shadow-md">
      <Transition show={true} enter="transition-opacity duration-300" enterFrom="opacity-0" enterTo="opacity-100" leave="transition-opacity duration-300" leaveFrom="opacity-100" leaveTo="opacity-0">
        <h2 className="text-2xl font-semibold mb-4">Registration</h2>
      </Transition>
      <RegistrationForm onRegister={onRegister} />
    </div>
  );
};

export default RegistrationPage;
