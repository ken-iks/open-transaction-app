import React, { useState } from "react";
import { login } from "./api";
import "./LoginForm.css"; 

// Login scripts for basic OAuth

const LoginForm: React.FC<{ setIsLoggedIn: (isLoggedIn: boolean) => void }> = ({ setIsLoggedIn }) => {
  const [accountNumber, setAccountNumber] = useState("");
  const [routingNumber, setRoutingNumber] = useState("");
  const [message, setMessage] = useState("");

  const handleLogin = async () => {
    const data = await login(accountNumber, routingNumber);
    if (data.error) {
      setMessage(data.error);
    } else {
      setMessage("Login successful!");
      setIsLoggedIn(true);
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2>Login</h2>
        <input
          type="text"
          placeholder="Account Number"
          value={accountNumber}
          onChange={(e) => setAccountNumber(e.target.value)}
          className="login-input"
        />
        <input
          type="text"
          placeholder="Routing Number"
          value={routingNumber}
          onChange={(e) => setRoutingNumber(e.target.value)}
          className="login-input"
        />
        <button onClick={handleLogin} className="login-button">
          Login
        </button>
        <p className="login-message">{message}</p>
      </div>
    </div>
  );
};

export default LoginForm;
