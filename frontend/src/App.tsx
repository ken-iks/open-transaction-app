import React, { useState, useEffect } from "react";
import "./App.css";
import { sendInput, checkSession, logout } from "./api";
import FetchMessage from "./FetchMessage";
import LoginForm from "./LoginForm";

function App() {
  const [input, setInput] = useState("");
  const [response, setResponse] = useState("");
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const verifySession = async () => {
      const sessionData = await checkSession();
      if (!sessionData.error) {
        setIsLoggedIn(true);
      }
    };
    verifySession();
  }, []);

  const handleSubmit = async () => {
    const data = await sendInput(input);
    if (data) setResponse(data.message);
  };

  const handleLogout = async () => {
    await logout();
    setIsLoggedIn(false);
  };

  if (!isLoggedIn) {
    return <LoginForm setIsLoggedIn={setIsLoggedIn} />;
  }

  return (
    <div className="app-container">
      <button onClick={handleLogout} className="logout-button">
        Logout
      </button>

      <h1>TRANSACTION API</h1>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Input message string"
        className="input-box"
      />
      <button onClick={handleSubmit} className="submit-button">
        Send
      </button>
      <p className="response-text">Response: {response}</p>
      <FetchMessage />
    </div>
  );
}

export default App;
