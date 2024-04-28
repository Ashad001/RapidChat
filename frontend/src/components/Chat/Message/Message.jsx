import React from "react";
import "./Message.scss";

const Message = ({ message }) => {
  const parsedMessage = JSON.parse(message);
  const formattedTimestamp = displayTime(parsedMessage.timeStamp);

  return (
    <div className="Message">
      <span className="timeStamp">{formattedTimestamp}</span>
      <span className="userName" style={{ color: parsedMessage.color }}>
        {parsedMessage.user}&nbsp;
      </span>
      <span className="messageBody">{parsedMessage.body}</span>
    </div>
  );
};

const displayTime = (timeStamp) => {
  const localeTime = new Date(timeStamp).toLocaleTimeString();
  return `${localeTime.substring(
    0,
    localeTime.length - 6
  )}\u00A0${localeTime.substring(localeTime.length - 2, localeTime.length)}`;
};

export default Message;
