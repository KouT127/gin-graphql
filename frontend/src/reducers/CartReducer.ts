import {AddCartPayload, ICartState, LoadedCartPayload} from "../app/interfaces/Cart";
import {AppState} from "../store/Store";
import {ThunkAction} from "redux-thunk";
import {Action, AnyAction, Dispatch} from "redux";

enum CartActionType {
    LoadedCartItems = "LOADED_CART_ITEMS",
}

// LoadedCartAction
interface LoadedCartAction extends Action {
    type: CartActionType.LoadedCartItems;
    payload: LoadedCartPayload;
}

const loadedCart = (payload: LoadedCartPayload): LoadedCartAction => {
    return {
        type: CartActionType.LoadedCartItems,
        payload
    };
};

// Reducer
const initialCartState: ICartState = {
    cartItems: []
};

const cartStateReducer = (state: ICartState = initialCartState, action: LoadedCartAction) => {
    switch (action.type) {
        case CartActionType.LoadedCartItems: {
            const cartItems = action.payload.cartItems;
            return {...state, cartItems};
        }
        default:
            return state;
    }
};

// Thunk Middleware
const addCart = (payload: AddCartPayload): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch, getState: () => AppState) => {
    const appState = getState();
    const carteState = {...appState.cartState};
    carteState.cartItems.push(payload.item);
    dispatch(loadedCart(carteState))
};

const cartActionCreators = {
    loadedCart,
};

// Selector
export const cartSelector = (state: AppState) => state.cartState;

export {
    addCart,
    cartStateReducer,
    cartActionCreators,
}
