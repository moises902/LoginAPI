import * as actionType from './actionTypes.js';

export const loginInputChange = (e) => ({
    type: actionType.LOGIN_INPUT_CHANGE,
    key: e.target.id,
    value: e.target.value
})

export const loginSubmit = () => ({
    type: actionType.LOGIN_SUBMIT
})

export const loginSuccess = (token) => {
    return ({
        type: actionType.LOGIN_SUCCESS,
        payload: token
    })
}

export const loginFailure = (code, message) => {
    return({
        type: actionType.LOGIN_FAILURE,
        payload: {code, message}
    });
}

export const loginReset = () => ({
    type: actionType.LOGIN_RESET
})