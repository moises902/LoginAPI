import * as actionType from './actionsTypes';

const initialState = {
    email: '',
    password1: '',
    password2: '',
    registerSuccess: false,
    registerError: false
}

export default (state = initialState, action) => {
    switch (action.type) {
        /* Changes the state by of username and password */
        case actionType.REGISTER_INPUT_CHANGE:
            return {
                ...state,
                [action.key]: action.value
            }
        case actionType.REGISTER_SUBMIT:
            return state;
        case actionType.REGISTER_SUCCESS:
            return {
                ...state,
                registerSuccess: true,
                registerError: false,
                success: action.payload
                };
        case actionType.REGISTER_ERROR:
            return {
                ...state,
                registerSuccess: false,
                registerError: true,
                error: action.payload
            }
        case actionType.REGISTER_RESET:
            state = initialState
            return state;
        default:
            return state;
    }
}