import { render } from "react-dom";
import React, {useState} from "react";
import {Provider} from 'react-redux';
import store from './rootredux/store.js';

import Login from "./login/login.jsx";
import styles from './styles.js';
import UserPage from './userpage/userpage.jsx'

const MainComp = () => {

    const [view, setView] = useState(true);
    
    return(
        <Provider store={store}>
            <div style={styles}>
                {view ? 
                    <Login setView={setView}/> :
                    <UserPage setView={setView}/>
                }
            </div>
        </Provider>
    )
}

render (<MainComp/>, document.getElementById('root'));