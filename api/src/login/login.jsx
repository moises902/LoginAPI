import React, {useState} from 'react';
import styles from '../styles.js';
import {useDispatch, useSelector} from 'react-redux';
import * as select from './selector.js';
import * as actions from './actions';
import Register from '../register/register.jsx'

const Login = (props) => {

    const login = useSelector(select.loginState);
    const dispatch = useDispatch();
    const [toggle, setToggle] = useState(true);

    const loginInputChange = (e) => {
        e.preventDefault();
        dispatch(actions.loginInputChange(e));
    }

    const loginSubmit = (e) => {
        e.preventDefault();
        dispatch(actions.loginSubmit());
    }
    
    if (login.loginSuccess){
        props.setView(false);
        return(null);
    } else return (
        <div>
            <form onSubmit={loginSubmit}>
                <div>
                    <label>Email:</label>
                    <input onChange={loginInputChange} type="text" id="email" required/>
                </div>
                <div>
                    <label>Password:</label>
                    <input onChange={loginInputChange} type="text" id="password" required/>
                </div>
                <div>
                {login.loginFailure &&
                        <p style={{color:'indianred', textAlign: 'center'}}>{login.error.code}:{login.error.message}</p>}
                </div>
                <div style={styles}>
                    <input type="submit" value="Login"/>
                </div>
            </form>
            <div style={styles}>
                {toggle ?
                    <button onClick={() => setToggle(false)}>Register</button> :
                    <Register setToggle={setToggle}/>}
            </div> 
        </div>
    )
}

export default Login;