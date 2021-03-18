import { select, takeEvery, all, call, put } from "redux-saga/effects";
import * as action from "./actions.js";
import * as actionType from "./actionsTypes.js";

export function* registeringUser() {
    
    const registerState = yield select(state => state.register);

    const response = yield fetch('http://127.0.0.1:8080/go/api/register', {
        method: 'POST',
        headers: {'Content-type':'application/json'},
        body: JSON.stringify({
            email: registerState.email,
            password1: registerState.password1,
            password2: registerState.password2
            })
    });
    const data = yield response.json();

    if (data.code === 200) {
        yield put(action.registerSuccess(data.code,data.message));
    } else {
        yield put(action.registerError(data.code,data.message));
    }
};

export function* registerUser() {
    yield takeEvery(actionType.REGISTER_SUBMIT, registeringUser);
};

export default function* root() {
    yield all([call(registerUser)])
};