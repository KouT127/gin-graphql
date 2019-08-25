import React from "react";
import {createStyles, Grid, makeStyles, Theme} from "@material-ui/core";
import ItemCard from "../components/ItemCard";
import {useDispatch} from "react-redux";
import {AddCartPayload} from "../app/interfaces/Cart";
import {addCart} from "../reducers/CartReducer";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        items: {
            marginTop: 8
        },
        gridSection: {
            [theme.breakpoints.up('xl')]: {
                marginRight: 80,
                marginLeft: 80
            },
        },
    }),
);

export interface IItem {
    id: string,
    name: string,
    description: string,
    imageUrl: string,
    price: number,
}


const ItemsPage: React.FC = () => {
    const classes = useStyles();
    const dispatch = useDispatch();

    const handleAdd = (item: IItem) => (e: React.MouseEvent) => {
        const payload: AddCartPayload = {item: item};
        dispatch(addCart(payload))
    };

    const item: IItem = {
        id: '1',
        name: 'name',
        description: 'description',
        imageUrl: 'https://placehold.jp/350x350.png',
        price: 3000
    };
    const items: Array<IItem> = [item, item, item, item, item, item];
    return (
        <Grid container className={classes.items}>
            <Grid item xs={12} className={classes.gridSection}>
                <Grid container item justify='flex-start' spacing={2}>
                    {items.map((item, index) => {
                        return <ItemCard
                            key={index}
                            item={item}
                            handleAdd={handleAdd(item)}/>
                    })}
                </Grid>
            </Grid>
        </Grid>
    );
};
export default ItemsPage;
