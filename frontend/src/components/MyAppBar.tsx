import {AppBar, Avatar, Badge, Button, createStyles, makeStyles, Theme, Toolbar, Typography} from "@material-ui/core";
import React from "react";
import Route from "../app/Route";
import {Link} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {signOut, userSelector} from "../reducers/UserReducer";
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import {deepOrange} from "@material-ui/core/colors";
import {cartSelector} from "../reducers/CartReducer";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        title: {
            flexGrow: 1,
            marginRight: 12,
        },
        cart: {
            marginRight: 12,
        },
        avatar: {
            backgroundColor: deepOrange[500],
        },
    }),
);

const MyAppBar: React.FC = () => {
    const classes = useStyles();
    const dispatch = useDispatch();
    const {user} = useSelector(userSelector);
    const {cartItems} = useSelector(cartSelector);
    const isLoggedIn = !!user;

    const handleSignOut = () => {
        dispatch(signOut())
    };

    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" className={classes.title}>
                    GraphQL Sample
                </Typography>
                {
                    isLoggedIn ?
                        <Button color="inherit"
                                component={Link} to={Route.signIn()}
                                onClick={handleSignOut}>
                            ログアウト
                        </Button> :
                        <Button color="inherit"
                                component={Link} to={Route.signIn()}>
                            ログイン
                        </Button>
                }
                <Badge badgeContent={cartItems.length} color="secondary"
                       className={classes.cart}>
                    <ShoppingCartIcon/>
                </Badge>
                <Avatar className={classes.avatar}>TK</Avatar>
            </Toolbar>
        </AppBar>
    )
};

export default MyAppBar;