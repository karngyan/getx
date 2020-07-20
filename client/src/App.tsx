import React from "react";
import logo from './assets/logo.svg';
import './App.css';
import { SearchURL } from "./components";

const App = () => {
  return (
    <div className="App">

      <main className="Content">
        <SearchURL />
      </main>
      <footer className="App-footer">
        <a href="/">
          <img src={logo} className="App-logo" alt="logo" width="70%"/>
        </a>
      </footer>
    </div>
  );
}
 
export default App;
