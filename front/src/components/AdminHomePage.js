import React from 'react'
import { Grid } from "@material-ui/core";

import AdminVerificationRequestCard from './AdminVerificationRequestCard'

const AdminHomePage = () => {
    return (
        <div>
            <Grid container>
                <Grid item xs={3}></Grid>
                <Grid item xs={6}><AdminVerificationRequestCard></AdminVerificationRequestCard></Grid>
                <Grid item xs={3}></Grid>
            </Grid>
        </div>
    )
}

export default AdminHomePage
