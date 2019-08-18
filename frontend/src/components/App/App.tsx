import React, {useEffect} from "react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import AuthorizationPage from "../../pages/AuthorizationPage";
import {makeStyles} from "@material-ui/core";
import TasksPage from "../../pages/TasksPage";
import {UserProvider, useUserState} from "../Providers/UserProvider";
import MyAppBar from "../MyAppBar";
import ApolloClient from "apollo-client";
import {InMemoryCache} from "apollo-cache-inmemory";
import {ApolloProvider} from "@apollo/react-hooks";
import {createHttpLink} from "apollo-link-http";

const useStyles = makeStyles({
    main: {
        background: `rgb(21, 32, 43)`
    },
});

const App: React.FC = () => {
    const classes = useStyles();

    return (
        <UserProvider>
            <Main/>
        </UserProvider>
    );
};

const Main = () => {
    const classes = useStyles();
    const {userState, isLoggedIn, authConnect} = useUserState();

    useEffect(() => {
        authConnect()
    }, []);


    const token = userState.user && userState.user.token ? `Bearer ${userState.user.token}` : '';
    const client = new ApolloClient({
        link: createHttpLink({
            uri: 'http://localhost:8080/query',
            headers: {
                authorization: token
            }
        }),
        cache: new InMemoryCache(),
    });

    return (
        <ApolloProvider client={client}>
            <div className={classes.main}>
                <Router>
                    <MyAppBar/>
                    <Route path="/" exact component={TasksPage}/>
                    <Route path="/users/signin/" exact component={AuthorizationPage}/>
                </Router>
            </div>
        </ApolloProvider>
    )
};

export default App;
