import * as actionType from './actionsTypes.js';

/* Action will take in the input from the register form */
export const registerInputChange = (e) => ({
    type: actionType.REGISTER_INPUT_CHANGE,
    key: e.target.id,
    value: e.target.value
})

export const registerSubmit = () => ({
    type: actionType.REGISTER_SUBMIT
})

/* Action occurs only when the registration was completed 
successfully */
export const registerSuccess = (code, message) => {
    return ({
        type: actionType.REGISTER_SUCCESS,
        payload: {code, message}
    })
}

export const registerError = (code, message) => {
    return({
        type: actionType.REGISTER_ERROR,
        payload: {code, message}
    });
}

export const registerReset = () => ({
    type: actionType.REGISTER_RESET
})