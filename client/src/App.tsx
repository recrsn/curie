import React from 'react';
import logo from './images/logo.svg';
import './App.css';
import Dashboard from './components/dashboard/Dashboard';

function App() {
  return (
    <div className="App">
        <header>
          <img src={logo} className="App-logo" alt="logo" />
          <span className="App-title">cURIe</span>
        </header>
        <main>
          <Dashboard />
        </main>
      </div>
  );
}

export default App;
