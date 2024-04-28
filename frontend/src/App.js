// App.js
import React, { Component } from 'react';
import './App.css';
import { Route, Switch } from 'react-router-dom';
import Header from './components/Header/Header';
import ChatPage from './components/Chat/ChatPage';
import LoginPage from './components/Login/LoginPage';
import { SecureRoute } from './authorization/secure.route';

class App extends Component {

  render() {
    return (
      <div className="App">
        <Header />
        <Switch>
          <Route exact path="/" component={LoginPage} />
          <SecureRoute exact path="/chat" component={ChatPage} />
          <Route path="*" component = {() => "4o4 NOT FOUND (plis no masti)"} />
        </Switch>
      </div>
    );
  }
}

export default App;
