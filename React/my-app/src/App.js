import React from 'react';
// import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import RouterView from './Router'

function App() {
  return (
      <Router>
          <div>
              <ul>
                  <li>
                      <Link to="/">Home</Link>
                  </li>
                  <li>
                      <Link to="/foo">Foo</Link>
                  </li>
                  <li>
                      <Link to="/bar">Bar</Link>
                  </li>
              </ul>
              <hr />
              <RouterView></RouterView>
          </div>
      </Router>
  );
}

export default App;
