import {applyMiddleware, combineReducers, createStore} from "redux";
import thunk from "redux-thunk";
import {userActionCreator, UserState, userStateReducer} from "../reducers/userReducer";

export type AppState = {
    users: UserState
}

export const store = createStore(
    combineReducers<AppState>({
        users: userStateReducer
    }),
    applyMiddleware(thunk)
);

export const actionCreator = {
    users: userActionCreator
};