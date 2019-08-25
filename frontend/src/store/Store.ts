import {applyMiddleware, combineReducers, createStore} from "redux";
import thunk from "redux-thunk";
import {IUserState, userActionCreators, userStateReducer} from "../reducers/UserReducer";
import {ICartState} from "../app/interfaces/Cart";
import {cartActionCreators, cartStateReducer} from "../reducers/CartReducer";

export type AppState = {
    userState: IUserState
    cartState: ICartState
}

export const store = createStore(
    combineReducers<AppState>({
        userState: userStateReducer,
        cartState: cartStateReducer,
    }),
    applyMiddleware(thunk)
);

export const actionCreator = {
    userActions: userActionCreators,
    cartActions: cartActionCreators
};
