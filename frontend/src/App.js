import React, { Component } from 'react';
import "./App.css";
import { connect, sendMsg } from './api';
import Header from './components/Header/Header';
import ChatHistory from './components/Chat/ChatHistory/ChatHistory';
import ChatInput from './components/Chat/ChatInput';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    }
  }

  send(event) {
    if(event.key === 'Enter'){
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state)
    });
  
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;