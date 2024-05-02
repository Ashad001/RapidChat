import React, { useState } from "react";
import auth from "../../authorization/auth";
import "./LoginPage.scss";

const LoginPage = ({ history }) => {
  const [name, setName] = useState("");
  const [roomId, setRoomId] = useState(""); // Add state for room ID
  const [roomName, setRoomName] = useState(""); // Add state for room name

  const handleChange = (e) => {
    setName(e.target.value);
  };

  const submitLogin = () => {
    if (name !== "" && roomName !== "") {
      auth.login(name, roomName, () => { // Pass room ID and room name
        history.push("/chat");
      });
    }
  };

  const keyPressed = (event) => {
    if (event.key === "Enter" && name !== "" && name !== undefined) {
      submitLogin();
    }
  };

  return (
    <div className="LoginPage">
      <div className="loginContainer">
        <div className="form__group field">
          <input
            type="input"
            className="form__field"
            name="name"
            id="name"
            value={name}
            onChange={handleChange}
            onKeyPress={keyPressed}
            placeholder=" "
          />
          <label className="form__label">Username</label>
        </div>
        <div className="form__group field">
          <input
            type="input"
            className="form__field"
            name="roomName"
            id="roomName"
            value={roomName}
            onChange={(e) => setRoomName(e.target.value)}
            placeholder=" "
          />
          <label className="form__label">Room Name</label>
        </div>
        {/* End of input fields for room ID and room name */}
        <button className="login-button" onClick={submitLogin}>
          Login
        </button>
      </div>
    </div>
  );
};

export default LoginPage;
