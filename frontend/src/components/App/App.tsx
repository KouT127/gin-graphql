import React from "react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import AuthorizationPage from "../../pages/AuthorizationPage";
import {makeStyles} from "@material-ui/core";
import TasksPage from "../../pages/TasksPage";
import {UserProvider} from "../Providers/UserProvider";
import MyAppBar from "../MyAppBar";

const useStyles = makeStyles({
    main: {
        background: `rgb(21, 32, 43)`
    },
});

const App: React.FC = () => {
    const classes = useStyles();
    return (
        <UserProvider>
            <div className={classes.main}>
                <Router>
                    <MyAppBar/>
                    <Route path="/" exact component={TasksPage}/>
                    <Route path="/users/signin/" exact component={AuthorizationPage}/>
                </Router>
            </div>
        </UserProvider>
    );
};

export default App;
