import React, { useState, useEffect, useRef } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginForm from './LoginForm';
import Home from './Home';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const myRef = useRef();

  useEffect(() => {
    const storedToken = localStorage.getItem('authToken');
    if (storedToken) {
      setIsLoggedIn(true); // Replace with actual token validation
    }
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
    localStorage.setItem('authToken', 'your-token'); // Replace with actual token value
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    localStorage.removeItem('authToken');
  };

  return (
    <Router>
      <div className="App">
        <h1>Página de login</h1>
        {isLoggedIn ? (
          <div>
            <button onClick={handleLogout}>Logout</button>
            <Routes>
              <Route path="/" element={<Home />} />
            </Routes>
          </div>
        ) : (
          <LoginForm onLogin={handleLogin} />
        )}
      </div>
    </Router>
  );
}

export default App;
