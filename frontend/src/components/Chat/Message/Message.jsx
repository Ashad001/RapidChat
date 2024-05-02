import React from "react";
import "./Message.scss";
//import SelectEmoji from "../Emoji";

const Message = ({ message }) => {
  try {
    const parsedMessage = JSON.parse(message);
    const formattedTimestamp = displayTime(parsedMessage.timeStamp);
    console.log(parsedMessage);
    return (
      <div className="Message">
        <div className="Details">
          <span className="userName" style={{ color: parsedMessage.color }}>
            {parsedMessage.user}&nbsp;
          </span>
          <span className="timeStamp">{formattedTimestamp}</span>
        </div>
        <div className="messageBody">
          <span>{parsedMessage.body}</span>
        </div>
      </div>
    );
  } catch (error) {
    console.error("Error parsing message:", error);
    return null; // or handle the error in a different way
  }
};

const displayTime = (timeStamp) => {
  const localeTime = new Date(timeStamp).toLocaleTimeString();
  return `${localeTime.substring(
    0,
    localeTime.length - 6
  )}\u00A0${localeTime.substring(localeTime.length - 2, localeTime.length)}`;
};

export default Message;
