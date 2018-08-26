import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Dashboard from './components/Dashboard';

class App extends Component {
  render() {
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
}

export default App;
