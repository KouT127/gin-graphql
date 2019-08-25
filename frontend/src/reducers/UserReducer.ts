import {Action, AnyAction, Dispatch} from "redux";
import {ThunkAction} from "redux-thunk";
import * as firebase from "../config/firebase";
import * as FirebaseType from "firebase"
import {AppState} from "../store/Store";


export interface User {
    id: string | null,
    name: string | null,
    email: string | null,
    emailVerified: boolean,
    token: string | null,
}

export interface IUserState {
    user: User | null
}

export const userSelector = (state: AppState) => state.userState;

const initialUserState: IUserState = {
    user: null
};

export type authorizeUserPayload = {
    email: string,
    password: string,
}
export type loadedUserPayload = {
    user: User | null
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

export const userStateReducer = (state: IUserState = initialUserState, action: LoadedUserAction) => {
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
export const connectAuth = (payload: void): ThunkAction<void, AppState, any, AnyAction> => {
    return (dispatch: Dispatch) => {
        firebase.default.auth().onAuthStateChanged(async (user: FirebaseType.User | null) => {
                let userPayload: loadedUserPayload = {user: null};
                if (user !== null) {
                    const authUser = {
                        id: user.uid,
                        name: user.displayName,
                        email: user.email,
                        emailVerified: user.emailVerified,
                        token: await user.getIdToken()
                    };
                    userPayload.user = authUser
                }
                dispatch(loadedUser(userPayload))
            }
        );
    };
};

export const signIn = (payload: authorizeUserPayload): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    firebase.default.auth().signInWithEmailAndPassword(payload.email, payload.password).catch((e) => {
        console.log(e)
    })
};

export const signOut = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    firebase.default.auth().signOut().catch((e) => {
        console.log(e)
    })
};

export const userActionCreators = {
    connectAuth,
    signIn,
    signOut
};
