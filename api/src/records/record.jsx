import React from "react";
import {useDispatch, useSelector} from 'react-redux';
import * as select from './selector.js';
import * as actions from './actions.js';
import SingleRecordFormat from './singleRecord.jsx'

const Record = () => {

    const record = useSelector(select.recordState);
    const dispatch = useDispatch();

    const recordInputChange = (e) => {
        e.preventDefault();
        dispatch(actions.recordsInputChange(e));
    }
    
    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(actions.recordSubmit());
    }

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label> Enter ID (Number): </label>
                <input onChange={recordInputChange} type="text" id="id"/>
                <input type="submit" value="Search a Record"/>
            </div>
            <div>
                {record.error &&
                    <p style={{color:'indianred', textAlign: 'center'}}> INVALID VALUE </p>}
            </div>
            <div>
                {record.recordData && <SingleRecordFormat data={record.recordData}/> }
            </div>
        </form>
    )
}

export default Record;