import {Action, AnyAction, Dispatch} from "redux";
import {ThunkAction} from "redux-thunk";
import {AppState} from "../store/store";
import * as firebase from "../config/firebase";

export interface UserState {
    user: any,
}

const initialUserState: UserState = {
    user: null
};

export type authorizeUserPayload = {
    email: string,
    password: string,
}
export type loadedUserPayload = {
    user: any
}

export interface LoadedUserAction extends Action {
    type: "LOADED_USERS";
    payload: loadedUserPayload;
}

export const loadedUser = (payload: loadedUserPayload): LoadedUserAction => {
    return {
        type: "LOADED_USERS",
        payload
    };
};

export const userStateReducer = (state: UserState = initialUserState, action: LoadedUserAction) => {
    switch (action.type) {
        case "LOADED_USERS": {
            const user = action.payload.user;
            return {...state, user};
        }
        default:
            return state;
    }
};

//Thunk-Actionの定義
export const connectAuth = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    firebase.default.auth().onAuthStateChanged((user) => {
            const userPayload: loadedUserPayload = {user: user}
            dispatch(loadedUser(userPayload))
        }
    )
};

export const signIn = (payload: authorizeUserPayload): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    firebase.default.auth().signInWithEmailAndPassword(payload.email, payload.password).catch((e) => {
        console.log(e)
    })
};

export const userActionCreator = {
    connectAuth,
    signIn,
};