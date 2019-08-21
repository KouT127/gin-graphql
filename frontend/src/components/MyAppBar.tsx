import {AppBar, Button, createStyles, makeStyles, Theme, Toolbar, Typography} from "@material-ui/core";
import React from "react";
import Routes from "../app/routes";
import {Link} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {signOut, userSelector} from "../reducers/UserReducer";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        title: {
            flexGrow: 1,
        },
    }),
);

const MyAppBar: React.FC = () => {
    const classes = useStyles();
    const dispatch = useDispatch();
    const {user} = useSelector(userSelector);
    const isLoggedIn = !!user;

    const handleSignOut = () => {
        dispatch(signOut())
    };

    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" className={classes.title}>
                    GraphQL Sample
                </Typography>
                {
                    isLoggedIn ?
                        <Button color="inherit"
                                component={Link} to={Routes.signIn()}
                                onClick={handleSignOut}>
                            SignOut
                        </Button> :
                        <Button color="inherit"
                                component={Link} to={Routes.signIn()}>
                            SingIn
                        </Button>
                }
            </Toolbar>
        </AppBar>
    )
};

export default MyAppBar;