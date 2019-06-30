import {Action, AnyAction, Dispatch} from "redux";
import {ThunkAction} from "redux-thunk";
import {AppState} from "../store/store";
import axios from 'axios';

export interface UserState {
    users: Array<string>,
}

const initialUserState: UserState = {
    users: []
};

export type loadedUserPayload = {
    users: Array<string>
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
            const users = action.payload.users;
            return {...state, users};
        }
        default:
            return state;
    }
};

//Thunk-Actionの定義
export const loadUser = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    axios.get("http://localhost:8080/users").then((res) => {
        const users: [any] = res.data
        const list = users.map((value) => value.toString())
        const payload: loadedUserPayload = {users: list}
        dispatch(loadedUser(payload))
    })

};

export const userActionCreator = {
    loadUser,
};