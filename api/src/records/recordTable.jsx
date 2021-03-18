import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Record from './record.jsx'

const useStyles = makeStyles({
    table: {
        width: 650,
    }
})

const RecordsTable = (props) => {

    const data = props.records;
    const rows = data.listOfRecords;
    const classes = useStyles()

    return (
        <>
        <div>
        <TableContainer component={Paper}>
            <Table className={classes.table} size="small" aria-label="Records Data" fixed="true">
                <TableHead>
                    <TableRow>
                        <TableCell>ID</TableCell>
                        <TableCell align="right">City</TableCell>
                        <TableCell align="right">Drug Name</TableCell>
                        <TableCell align="right">First Name</TableCell>
                        <TableCell align="right">Generic Name</TableCell>
                        <TableCell align="right">Last Org. Name</TableCell>
                        <TableCell align="right">Specialty</TableCell>
                        <TableCell align="right">State</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.id}>
                            <TableCell component="th" scope="row">{row.id}</TableCell>
                            <TableCell align="right">{row.city}</TableCell>
                            <TableCell align="right">{row.drugName}</TableCell>
                            <TableCell align="right">{row.firstName}</TableCell>
                            <TableCell align="right">{row.genericName}</TableCell>
                            <TableCell align="right">{row.lastOrgName}</TableCell>
                            <TableCell align="right">{row.specialty}</TableCell>
                            <TableCell align="right">{row.state}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
        </div>
        <div>
            <Record/>
        </div>
        </>
    )
}

export default RecordsTable;