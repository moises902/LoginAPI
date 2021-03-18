import * as actionType from './actionTypes.js';

const initialState = {
    email: sessionStorage.getItem('email') || '',
    password: '',
    loginSuccess: sessionStorage.getItem('loginSuccess') || false,
    loginFailure: false,
    token: sessionStorage.getItem('token') || null
}

export default (state = initialState, action) => {
    switch(action.type) {
        case actionType.LOGIN_INPUT_CHANGE:
            return {
                ...state,
                [action.key]: action.value
            };
        case actionType.LOGIN_SUBMIT:
            return state;
        case actionType.LOGIN_SUCCESS:
            sessionStorage.setItem('loginSuccess', 'true')
            return {
                ...state,
                loginSuccess: true,
                token: action.payload
            }
        case actionType.LOGIN_FAILURE:
            return {
                ...state,
                loginFailure: true,
                error: action.payload
            }
        case actionType.LOGIN_RESET:
            sessionStorage.clear()
            state = {
                ...initialState,
                loginSuccess: false
            }
            return state;
        default:
            return state;
    }
}