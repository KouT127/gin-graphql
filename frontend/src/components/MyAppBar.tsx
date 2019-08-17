import {AppBar, Button, createStyles, makeStyles, Theme, Toolbar, Typography} from "@material-ui/core";
import React from "react";
import Routes from "../app/routes";
import {Link} from "react-router-dom";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        title: {
            flexGrow: 1,
        },
    }),
);

const MyAppBar: React.FC = () => {
    const classes = useStyles();
    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" className={classes.title}>
                    GraphQL Sample
                </Typography>
                <Button color="inherit"
                        component={Link} to={Routes.signIn()}
                >Login</Button>
            </Toolbar>
        </AppBar>
    )
};

export default MyAppBar;