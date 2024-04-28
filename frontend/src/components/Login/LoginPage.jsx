import React, { useState } from "react";
import auth from "../../authorization/auth";
import "./LoginPage.scss";

const LoginPage = ({ history }) => {
  const [name, setName] = useState("");

  const handleChange = (e) => {
    setName(e.target.value);
  };

  const submitLogin = () => {
    auth.login(name, () => {
      history.push("/chat");
    });
  };

  const keyPressed = (event) => {
    if (event.key === "Enter") {
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
          <label className="form__label">
            Username
          </label>
        </div>

        <button className="login-button" onClick={submitLogin}>
          Login
        </button>
      </div>
    </div>
  );
};

export default LoginPage;
