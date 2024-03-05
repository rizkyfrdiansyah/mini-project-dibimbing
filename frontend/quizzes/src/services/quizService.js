const getActiveQuizzes = async () => {
  const response = await fetch("/api/quizzes/active");
  const data = await response.json();
  return data.quizzes;
};

const createQuiz = async (quizData) => {
  const response = await fetch("/api/quizzes", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(quizData),
  });

  if (response.ok) {
    const data = await response.json();
    return data.quiz;
  } else {
    throw new Error("Quiz creation failed");
  }
};

export { getActiveQuizzes, createQuiz };
