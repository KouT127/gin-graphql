import React, {useEffect} from "react";
import {useDispatch} from "react-redux";
import {connectAuth} from "../../reducers/userReducer";
import {BrowserRouter as Router, Link, Route} from "react-router-dom";
import AuthorizationPage from "../../pages/AuthorizationPage";


const App: React.FC = () => {
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(connectAuth());
    }, [dispatch]);
    
    return (
        <Router>
            <ul>
                <li>
                    <Link to="/auth/">Auth</Link>
                </li>
            </ul>
            <Route path="/auth/" exact component={AuthorizationPage}/>
        </Router>
    );
};

export default App;
