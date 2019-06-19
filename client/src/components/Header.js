import React, { Component } from 'react';
import logo from '../logo.svg';

class Header extends Component {
  render() {
    return (
      <header>
        <nav className="navbar navbar-dark bg-dark">
          <div className="container">
            <img src={logo} className="App-logo" alt="logo" />
            <ul className="navbar-nav">
              <li className="nav-item">
                <a className="nav-link" href="/">
                  Главная страница
                </a>
              </li>
            </ul>
          </div>
        </nav>
      </header>
    );
  }
}

export default Header;