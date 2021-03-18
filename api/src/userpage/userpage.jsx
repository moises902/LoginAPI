import React from 'react';
import styles from '../styles.js';
import {useDispatch, useSelector} from 'react-redux';
import * as select from './selector.js';
import * as actions from '../login/actions.js'
import * as recordsActions from '../records/actions.js'
import Records from "../records/records.jsx"

const UserPage = (props) => {
    
    const login = useSelector(select.loginState);
    const dispatch = useDispatch();

    const onClick = (e) => {
        e.preventDefault();
        dispatch(actions.loginReset());
        dispatch(recordsActions.recordsReset());
    }

    if (!login.loginSuccess) {
        props.setView(true);
        return (null);
    } else return (
        <div>
        <div style={styles}>
            <h1>Hello {login.email}</h1> 
        </div>
        <div style={styles}>
            <Records/>
        </div>
        <div style={styles}>
            <button onClick={onClick}>Logout</button>
        </div>
        </div>
    )
};

export default UserPage;