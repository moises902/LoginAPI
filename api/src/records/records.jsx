import React from "react";
import {useDispatch, useSelector} from 'react-redux';
import * as select from './selector.js';
import * as actions from './actions.js';
import styles from '../styles.js';
import RecordsTable from './recordTable.jsx';

const Records = () => {

    const records = useSelector(select.recordState);
    const dispatch = useDispatch();

    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(actions.recordsSubmit());
    }

    const recordsInputChange = (e) => {
        e.preventDefault();
        dispatch(actions.recordsInputChange(e));
    }

    return (
        <div style={styles}>
        <form onSubmit={handleSubmit}>
            <div>
                <label>City:</label>
                <input onChange={recordsInputChange} type="text" id="city" />
            </div>
            <div>
                <label>State:</label>
                <input onChange={recordsInputChange} type="text" id="state"/>
            </div>
            <div>
                <label>Specialty:</label>
                <input onChange={recordsInputChange} type="text" id="specialty"/>
            </div>
            <div>
                <label>Drug Name:</label>
                <input onChange={recordsInputChange} type="text" id="drugName"/>
            </div>
            <input type="submit" value="Search Records"/>
        </form>
        <div style={styles}>
            {records.listOfRecords && <RecordsTable records={records}/>}
        </div>
        </div>
    )
}

export default Records;