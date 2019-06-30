import React, {useEffect} from 'react';
import {useDispatch, useSelector} from "react-redux";
import {AppState} from "../../store/store";
import {loadUser} from "../../reducers/userReducer";

const userSelector = (state: AppState) => state.users;

const App: React.FC = () => {
    const dispatch = useDispatch()
    const userState = useSelector(userSelector)

    useEffect(() => {
        dispatch(loadUser())
    }, [dispatch])

    const listItem = userState.users.map((user) => <li>{user}</li>)

    return (
        <div className="App">
            <ul>
                {listItem}
            </ul>
        </div>
    );
}

export default App;
