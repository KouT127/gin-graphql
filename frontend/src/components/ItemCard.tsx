import {Button, Card, CardContent, createStyles, Grid, makeStyles, Theme} from "@material-ui/core";
import React from "react";
import {IItem} from "../pages/ItemsPage";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        content: {
            width: `300px`,
            display: `flex`,
            flexDirection: `column`
        },
        description: {
            flexWrap: `wrap`
        },
        card: {
            padding: `12px 12px 12px 12px`
        }
    }),
);

type ItemCardProps = {
    item: IItem,
    handleAdd: (e: React.MouseEvent) => void,
}

const ItemCard = (props: ItemCardProps) => {
    const classes = useStyles();
    const {item, handleAdd} = props;

    return <Grid item>
        <Card className={classes.card}>
            <img src={item.imageUrl}/>
            <CardContent className={classes.content}>
                <h3>{item.name}</h3>
                <p className={classes.description}>{item.description}</p>
                <p>Price ￥{item.price}</p>
            </CardContent>
            <Button variant={'contained'}
                    color={'primary'}
                    fullWidth={true}
                    onClick={handleAdd}
            >
                Add to cart
            </Button>
        </Card>
    </Grid>
};

export default ItemCard;