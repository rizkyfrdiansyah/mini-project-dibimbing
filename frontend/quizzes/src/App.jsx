import React, { useState } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import LoginPage from "./pages/Auth/LoginPage";
import RegistrationPage from "./pages/Auth/RegistrationPage";
import DashboardPage from "./pages/Admin/DashboardPage";
import QuizListPage from "./pages/Participant/QuizListPage";
import QuizPage from "./pages/Participant/QuizPage";

function App() {
  const [user, setUser] = useState(null);

  const handleLogin = (userData) => {
    console.log("Login:", userData);
  };

  const handleRegister = (userData) => {
    console.log("Register:", userData);
  };

  const handleCreateQuiz = (quizData) => {
    console.log("Create Quiz:", quizData);
  };

  return (
    <Router>
      <div>
        <Switch>
          {/* Authentication Routes */}
          <Route path="/login">
            <LoginPage onLogin={handleLogin} />
          </Route>
          <Route path="/register">
            <RegistrationPage onRegister={handleRegister} />
          </Route>

          {/* Admin Routes */}
          <Route path="/admin/dashboard">
            <DashboardPage quizzes={[]} onCreateQuiz={handleCreateQuiz} />
          </Route>

          {/* Participant Routes */}
          <Route path="/participant/quizzes" exact>
            <QuizListPage quizzes={[]} />
          </Route>
          <Route path="/participant/quiz/:quizId">
            <QuizPage />
          </Route>

          {/* Default Route (Home or Login based on user state) */}
          <Route path="/" exact>
            {user ? <DashboardPage quizzes={[]} onCreateQuiz={handleCreateQuiz} /> : <LoginPage onLogin={handleLogin} />}
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
