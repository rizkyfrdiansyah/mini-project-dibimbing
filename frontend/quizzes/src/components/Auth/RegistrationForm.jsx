import React, { useState } from "react";
import { Transition } from "@headlessui/react";

const RegistrationForm = ({ onRegister }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleRegister = () => {
    onRegister({ username, password });
  };

  return (
    <div className="max-w-md mx-auto mt-4 p-4 bg-white rounded-md shadow-md">
      <Transition show={true} enter="transition-opacity duration-300" enterFrom="opacity-0" enterTo="opacity-100" leave="transition-opacity duration-300" leaveFrom="opacity-100" leaveTo="opacity-0">
        <h2 className="text-xl font-semibold mb-4">Register</h2>
      </Transition>
      <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} className="w-full border border-gray-300 p-2 mb-2 rounded-md focus:outline-none focus:ring focus:border-blue-500" />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} className="w-full border border-gray-300 p-2 mb-2 rounded-md focus:outline-none focus:ring focus:border-blue-500" />
      <button onClick={handleRegister} className="w-full bg-green-500 text-white p-2 rounded-md hover:bg-green-600 focus:outline-none focus:ring focus:border-green-700">
        Register
      </button>
    </div>
  );
};

export default RegistrationForm;
