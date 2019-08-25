import {IItem} from "../../pages/ItemsPage";

export interface ICartState {
    cartItems: Array<IItem>
}

export type LoadedCartPayload = {
    cartItems: Array<IItem>
}

export type AddCartPayload = {
    item: IItem
}