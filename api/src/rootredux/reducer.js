import {combineReducers} from 'redux';
import register from '../register/reducer.js';
import login from '../login/reducer.js'
import records from '../records/reducer.js'

const rootReducer = combineReducers({register, login, records});

export default rootReducer;
