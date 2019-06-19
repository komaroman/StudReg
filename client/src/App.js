import React, { Component } from 'react';
import './App.css';
import Footer from './components/Footer';
import Header from './components/Header';
import InvokeForm from './components/InvokeForm';
import QueryForm from './components/QueryForm';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header />
        <div className="container mt-3 mb-3">
          <div className="row">
            <div className="col-md-6">
              <InvokeForm />
            </div>
            <div className="col-md-6">
              <QueryForm />
            </div>
          </div>
        </div>
        <Footer/>
      </div>
    );
  }
}

export default App;
