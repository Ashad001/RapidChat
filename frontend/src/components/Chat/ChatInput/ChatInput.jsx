import React from "react";
import "./ChatInput.scss";

const ChatInput = ({ send }) => {
  return (
    <div className="ChatInput">
      <input
        id="messageInput"
        onKeyDown={send}
        placeholder="Type your message..."
      />
    </div>
  );
};

export default ChatInput;
