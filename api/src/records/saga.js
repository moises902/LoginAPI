import { select, takeEvery, all, call, put } from "redux-saga/effects";
import * as action from './actions.js';
import * as actionType from './actionTypes.js';

export function* gettingRecords() {

    const recordState = yield select(state => state.records);
    const url = new URL('http://127.0.0.1:8080/go/api/records');
    url.search = new URLSearchParams({
        city: recordState.city,
        state: recordState.state,
        specialty: recordState.specialty,
        drugName: recordState.drugName
    }).toString();

    const response = yield fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-type':'application/json',
            // Authorization: 'Bearer ' + loginState.token
        }
    })

    const data = yield response.json();
    yield put(action.recordsStore(data));
}

export function* gettingRecord() {

    const recordState = yield select(state => state.records);
    const url = new URL('http://127.0.0.1:8080/go/api/record');
    url.search = new URLSearchParams({
        id: recordState.id
    }).toString();

    const response = yield fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-type':'application/json',
            // Authorization: 'Bearer ' + loginState.token 
        }
    })

    const data = yield response.json();
    
    if (data.code) {
        yield put(action.recordError());
    } else {
        yield put(action.singleRecord(data))
    }
}

export function* getRecords() {
    yield takeEvery(actionType.RECORDS_SUBMIT, gettingRecords);
    yield takeEvery(actionType.RECORD_SUBMIT, gettingRecord);
}

export default function* root() {
    yield all([call(getRecords)]);
}