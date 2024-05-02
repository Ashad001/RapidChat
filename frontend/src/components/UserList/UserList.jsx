import React, { Component } from "react";
import "./UserList.scss";

class UserList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      users: props.userList || []
    };
  }

  componentDidUpdate(prevProps) {
    if (prevProps.userList !== this.props.userList) {
      this.setState({ users: this.props.userList || [] });
    }
  }

  handleUserLeave = userId => {
    this.setState(prevState => ({
      users: prevState.users.filter(user => user.id !== userId)
    }));
  };

  render() {
    const { users } = this.state;
    return (
      <div className="UserList">
        <div className="user-list">
          <p className="usersLabel">Online</p>
          {users.map(user => (
            <p
              key={user.id}
              className={`user ${user.leaving ? "leave" : ""}`}
              style={{ color: user.color }}
              onAnimationEnd={() => {
                if (user.leaving) {
                  this.handleUserLeave(user.id);
                }
              }}
            >
              {user.name}
            </p>
          ))}
        </div>
      </div>
    );
  }
}

export default UserList;
