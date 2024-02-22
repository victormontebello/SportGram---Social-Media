import React, { useState } from 'react';
import axios from 'axios'; // For API requests
import './LoginForm.css'; // Import your CSS file

function LoginForm() {
  const [email, setEmail] = useState('');
  const [Senha, setSenha] = useState('');
  const [errorMessage, setErrorMessage] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Validate input
    if (!email || !Senha) {
      setErrorMessage('Please enter both email and Senha.');
      return;
    }

    try {
      // Send login request to your API
      const response = await axios.post('/api/login', { email, Senha });

      if (response.data.success) {
        localStorage.setItem('authToken', response.data.token);
        window.location.href = '/home'; // Redirect to home page
      } else {
        setErrorMessage(response.data.message || 'Email ou Senha inválidos.');
      }
    } catch (error) {
      console.error(error);
      setErrorMessage('Email ou Senha inválidos.');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="login-form">
      {errorMessage && <div className="error-message">{errorMessage}</div>}
      <div className="form-group">
        <label htmlFor="email" className="form-label">Email:</label>
        <input
          type="text"
          id="email"
          className="form-input"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>
      <div className="form-group">
        <label htmlFor="Senha" className="form-label">Senha:</label>
        <input
          type="Senha"
          id="Senha"
          className="form-input"
          value={Senha}
          onChange={(e) => setSenha(e.target.value)}
        />
      </div>
      <button type="submit" className="login-button">Login</button>
    </form>
  );
}

export default LoginForm;