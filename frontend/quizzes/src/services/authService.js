const login = async (credentials) => {
  const response = await fetch("/api/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(credentials),
  });

  if (response.ok) {
    const data = await response.json();
    return data.token;
  } else {
    throw new Error("Login failed");
  }
};

const register = async (userData) => {
  const response = await fetch("/api/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  });

  if (response.ok) {
    const data = await response.json();
    return data.token;
  } else {
    throw new Error("Registration failed");
  }
};

export { login, register };
