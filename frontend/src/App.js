import React, { Component } from 'react';
import "./App.css";
import { connect, sendMsg } from './api';

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("Hello")
    sendMsg("Hello")
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <button onClick={this.send}>Send</button>
        </header>
      </div>
    );
  }
}

export default App;