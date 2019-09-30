import React from 'react';
// import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Bar from "./views/Bar";
import Foo from "./views/Foo";
import Home from "./views/Home";

function App() {
  return (
      <Router>
          <div>
              <nav>
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
              </nav>

              {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
              <Switch>
                  <Route path="/foo">
                      <Foo />
                  </Route>
                  <Route path="/bar">
                      <Bar />
                  </Route>
                  <Route path="/">
                      <Home />
                  </Route>
              </Switch>
          </div>
      </Router>
  );
}

export default App;
