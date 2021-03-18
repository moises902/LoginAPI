import * as actionType from './actionTypes.js';

const initialState = {
    city: '',
    state: '',
    specialty: '',
    drugName: '',
    id: null,
    error: false
}

export default (state = initialState, action) => {
    switch(action.type) {
        case actionType.RECORDS_SUBMIT:
            return state;
        case actionType.RECORDS_INPUT_CHANGE:
            return {
                ...state,
                [action.key]: action.value
            }
        case actionType.RECORDS_STORE:
            return {
                ...state,
                listOfRecords: action.payload
            }
        case actionType.RECORDS_RESET:
            state = initialState;
            return state;
        case actionType.RECORD_SUBMIT:
            return {
                ...state,
                error: false
            }
        case actionType.SINGLE_RECORD:
            return {
                ...state,
                recordData: action.payload
            }
        case actionType.RECORD_ERROR:
            return {
                ...state,
                error: true
            }
        default:
            return state;
    }
}