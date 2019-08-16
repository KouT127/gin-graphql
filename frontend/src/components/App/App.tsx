import React, {useEffect} from "react";
import {useDispatch} from "react-redux";
import {connectAuth} from "../../reducers/userReducer";
import {BrowserRouter as Router, Link, Route} from "react-router-dom";
import AuthorizationPage from "../../pages/AuthorizationPage";
import {Button, makeStyles, Toolbar} from "@material-ui/core";
import TasksPage from "../../pages/TasksPage";
import Routes from "../../app/routes";

const useStyles = makeStyles({
    main: {
        background: `#7986cb`
    },
});

const App: React.FC = () => {
    const classes = useStyles();
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(connectAuth());
    }, [dispatch]);

    return (
        <div className={classes.main}>
            <Router>
                <Toolbar>
                    <Button variant="contained"
                            color="secondary"
                            component={Link} to={Routes.signIn()}
                    >
                        SignIn
                    </Button>
                </Toolbar>
                <Route path="/" exact component={TasksPage}/>
                <Route path="/users/signin/" exact component={AuthorizationPage}/>
            </Router>
        </div>
    );
};

export default App;
