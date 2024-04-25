import React, { Component } from "react";
import "./ChatInput.scss";

class ChatInput extends Component {
    constructor(props) {
        super(props);
        // this.state = {
        //     message: ""
        // };
    }

    // handleChange = (e) => {
    //     this.setState({ message: e.target.value });
    // };

    // handleKeyPress = (e) => {
    //     if (e.key === 'Enter') {
    //         this.sendMessage();
    //     }
    // };

    // sendMessage = () => {
    //     if (this.state.message.trim() !== "") {
    //         this.props.send(this.state.message);
    //         this.setState({ message: "" });
    //     }
    // };

    render() {
        return (
            <div className="ChatInput">
                <input
                    type="text"
                    onKeyDown={this.props.send}
                    placeholder="Enter your message..."
                />
                <button onClick={this.props.send}>Send</button>
            </div>
        );
    }
}

export default ChatInput;
