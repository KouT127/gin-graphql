import {Button, Card, CardContent, createStyles, Grid, makeStyles, Theme, Typography} from "@material-ui/core";
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


const ItemCard = (props: { item: IItem }) => {
    const classes = useStyles();
    const item = props.item;

    return <Grid item>
        <Card className={classes.card}>
            <img src={item.imageUrl}/>
            <CardContent className={classes.content}>
                <Typography>
                    {item.name}
                </Typography>
                <Typography className={classes.description}>
                    {item.description}
                </Typography>
                <Typography>
                    Price {item.price}
                </Typography>
            </CardContent>
            <Button variant={'contained'}
                    color={'primary'}
                    fullWidth={true}
            >
                Add to cart
            </Button>
        </Card>
    </Grid>
};

export default ItemCard;