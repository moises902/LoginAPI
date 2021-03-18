import React from 'react';
import styles from '../styles'
import {useDispatch, useSelector} from 'react-redux';
import * as select from './selector.js'
import * as actions from './actions.js'

const Register = (props) => {
    
    const register = useSelector(select.registerState);
    const dispatch = useDispatch();

    const registerInputChange = (e) => {
        e.preventDefault();
        dispatch(actions.registerInputChange(e));
    }

    const registerSubmit = (e) => {
        e.preventDefault();
        dispatch(actions.registerSubmit());
    }

    const handleClick = (e) => {
        e.preventDefault;
        props.setToggle(true);
        dispatch(actions.registerReset());
    }

    return (
        <div>
            <form onSubmit={registerSubmit}>
                <div>
                    <label>Email:</label>
                    <input onChange={registerInputChange} type="text" id="email"  required/>
                </div>
                <div>
                    <label>Password 1:</label>
                    <input onChange={registerInputChange} type="text" id="password1"  required/>
                </div>
                <div>
                    <label>Password 2:</label>
                    <input onChange={registerInputChange} type="text" id="password2" required/>
                </div>
                <div>
                    {register.registerSuccess &&
                        <div>
                            <p style={{color:'green', textAlign: 'center'}}>{register.success.code}:{register.success.message}</p>
                            <p style={{color:'green', textAlign: 'center'}}>You may now close the registration form</p>    
                        </div>}
                    {register.registerError &&
                        <p style={{color:'indianred', textAlign: 'center'}}>{register.error.code}:{register.error.message}</p>}
                </div>
                <div style={styles}>
                    <input type="submit" value="Submit"/>
                </div> 
            </form>
            <div style={styles}>
                <button onClick={handleClick}>Close</button>
            </div>
        </div>
    )
}

export default Register;