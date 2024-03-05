import React from "react";
import LoginForm from "../components/Auth/LoginForm";
import { Transition } from "@headlessui/react";

const LoginPage = ({ onLogin }) => {
  return (
    <div className="max-w-md mx-auto mt-8 p-4 bg-white rounded-md shadow-md">
      <Transition show={true} enter="transition-opacity duration-300" enterFrom="opacity-0" enterTo="opacity-100" leave="transition-opacity duration-300" leaveFrom="opacity-100" leaveTo="opacity-0">
        <h2 className="text-2xl font-semibold mb-4">Login</h2>
      </Transition>
      <LoginForm onLogin={onLogin} />
    </div>
  );
};

export default LoginPage;
