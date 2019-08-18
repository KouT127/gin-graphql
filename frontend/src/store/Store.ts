import {applyMiddleware, combineReducers, createStore} from "redux";
import thunk from "redux-thunk";
import {userActionCreator, UserState, userStateReducer} from "../reducers/UserReducer";

export type AppState = {
    userState: UserState
}

export const store = createStore(
    combineReducers<AppState>({
        userState: userStateReducer
    }),
    applyMiddleware(thunk)
);

export const actionCreator = {
    userActions: userActionCreator
};
