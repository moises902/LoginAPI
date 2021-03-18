import * as actionType from './actionTypes.js';

export const recordsInputChange = (e) => ({
    type: actionType.RECORDS_INPUT_CHANGE,
    key: e.target.id,
    value: e.target.value
})

export const recordsSubmit = () => ({
    type: actionType.RECORDS_SUBMIT
})

export const recordSubmit = () => ({
    type: actionType.RECORD_SUBMIT
})

export const recordsStore = (recordsList) => {
    return ({
        type: actionType.RECORDS_STORE,
        payload: recordsList
    })
}

export const recordsReset = () => ({
    type: actionType.RECORDS_RESET
})

export const singleRecord = (data) => {
    return ({
        type: actionType.SINGLE_RECORD,
        payload: data
    })
}

export const recordError = () => ({
    type: actionType.RECORD_ERROR
})