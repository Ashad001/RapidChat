import React, { Component } from 'react';
import "./Header.scss";
import auth from '../../authorization/auth';
import { withRouter } from 'react-router-dom';

class Header extends Component {
    handleLogout() {
        auth.logout(() => {
            this.props.history.push('/');
        });
    }

    render() {
        return (
            <div className='header'>
                <h3>Rapid Chat</h3>
                <button className='logout-btn' onClick={() => {
                    this.handleLogout()
                }}>Logout</button>
            </div>
        );
    }
}

export default withRouter(Header);
