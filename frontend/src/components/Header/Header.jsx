import React from 'react';
import "./Header.scss";
import auth from '../../authorization/auth';
import { useHistory } from 'react-router-dom';

const Header = () => {
    const history = useHistory();

    const handleLogout = () => {
        auth.logout(() => {
            history.push('/');
        });
    };

    return (
        <div className='header'>
            <h3>Rapid Chat</h3>
            <button className='logout-btn' onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default Header;
