import React from "react";
import "./Room.scss";

const Room = ({ roomName }) => {
  return (
    <div className="RoomHeader">
      <h3 className="roomName">{roomName}</h3>
    </div>
  );
};

export default Room;
