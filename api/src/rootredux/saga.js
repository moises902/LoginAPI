import { all }  from "redux-saga/effects";
import registerSaga from "../register/saga.js";
import loginSaga from "../login/saga.js";
import recordSaga from "../records/saga.js"

function* rootSaga() {
    yield all([registerSaga(), loginSaga(), recordSaga()]);
}

export default rootSaga;