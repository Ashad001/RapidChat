import React, { Component } from "react";
import "./ChatHistory.scss";
import Message from "../Message";

class ChatHistory extends Component {

    componentDidMount() {
        this.scrollToBottom();
    }

    componentDidUpdate() {
        this.scrollToBottom();
    }
    
    scrollToBottom = () => {
        this.el.scrollIntoView({ behavior: "smooth" });
    }

    render() {
        const messages = this.props.chatHistory.map((msg, index) => (<Message key={index} message={msg.data} />));
        console.log("Messages: ", messages);
        return (
            <div className="ChatHistory">
                <div id="chatHistory" className="disable-scrollbars">
                    <div id="history">
                        {messages}
                    </div>
                    <div style={{ float:"left", clear: "both" }}
                        ref={el => { this.el = el; }} />
                </div>
            </div>
        )
    }
}

export default ChatHistory;