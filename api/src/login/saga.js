import { select, takeEvery, all, call, put } from "redux-saga/effects";
import * as action from './actions.js';
import * as actionType from './actionTypes.js';

export function* loggingInUser() {

    const loginState = yield select(state => state.login);

    const response = yield fetch('http://localhost:8080/go/api/login', {
        method: 'POST',
        credentials: 'include',
        headers: {'Content-type':'application/json'},
        body: JSON.stringify({
            email: loginState.email,
            password: loginState.password
        })
    });

    const data = yield response.json();

    if (data.email) {
        console.log(data);
        yield put(action.loginSuccess(data.token));
        sessionStorage.setItem('token', data.token);
        sessionStorage.setItem('email', data.email);
    } else {
        yield put(action.loginFailure(data.code, data.message));
    }
}

export function* loginUser () {
    yield takeEvery(actionType.LOGIN_SUBMIT, loggingInUser);
}

export default function* root(){
    yield all([call(loginUser)])
}