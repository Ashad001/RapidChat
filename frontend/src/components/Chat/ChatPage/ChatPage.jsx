import React, { Component } from "react";
import { Redirect } from "react-router-dom";
import ChatSocket from "../../../api/ChatSocket";
import "./ChatPage.scss";
import ChatHistory from "../ChatHistory/ChatHistory";
import ChatInput from "../ChatInput";
import UserList from "../../UserList";
import Room from "../../Room";
import auth from "../../../authorization/auth";

class ChatPage extends Component {
  _chatSocket;

  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
      userList: [],
      roomColor: "",
    };
  }

  componentDidMount() {
    if (auth.isAuthenticated()) {
      this._chatSocket = new ChatSocket(
        "ws://localhost:8080/ws",
        auth.getUserName(),
        auth.getUserId(),
        auth.getRoomName(),
        true
      );
      this._chatSocket.connect((event) => {
        this.handleSocketEvent(event);
      });
    }
  }

  componentWillUnmount() {
    this._chatSocket.closeSocket();
  }

  handleSocketEvent(event) {
    switch (event.type) {
      case "close":
        // Handle WebSocket close event (e.g., notify user about logout)
        this.handleLogout();
        break;
      case "message":
        this.handleMessage(event);
        break;
      default:
        console.log("Unknown WebSocket event type");
    }
  }

  handleLogout() {
    auth.logout(() => {
      this.props.history.push("/");
    });
  }

  handleMessage(event) {
    const msgData = JSON.parse(event.data);

    switch (msgData.type) {
      case 0:
        this.setState({
          userList: msgData.clientList,
        });
        break;
      case 1:
        this.setState((prevState) => ({
          chatHistory: [...prevState.chatHistory, event],
        }));
        break;
      default:
        console.log("Unknown message type");
    }
  }

  send(event) {
    if (event.keyCode === 13 && event.target.value !== "") {
      console.log("Sending message >> ", event.target.value);
      this._chatSocket.sendMessage(event.target.value, auth.getUserId());
      event.target.value = "";
    }
  }

  render() {
    if (!auth.isAuthenticated()) {
      return <Redirect to="/" />;
    }

    return (
      <div className="ChatPage">
        <Room roomName={auth.getRoomName()} />
        <UserList userList={this.state.userList} />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={(e) => this.send(e)} />
      </div>
    );
  }
}

export default ChatPage;
