import React, {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {AppState} from "../store/store";
import {authorizeUserPayload, signIn} from "../reducers/userReducer";
import {Button, Card, CardContent, makeStyles, TextField} from "@material-ui/core";
import useReactRouter from 'use-react-router';
import Routes from "../app/routes";

const userSelector = (state: AppState) => state.userState;

const useStyles = makeStyles({
    card: {
        display: `flex`,
        justifyContent: `center`,
        marginTop: 24,
        marginBottom: 24
    },
    cardContent: {
        display: `flex`,
        flexDirection: `column`,
        flexGrow: 1,
        marginTop: 8,
        marginBottom: 8
    },
    textField: {
        marginBottom: 10,
    },
    button: {
        marginTop: 10,
    },
    auth: {
        display: `grid`,
        gridTemplateRows: `2000px`,
        gridTemplateColumns: `1fr 40% 1fr`,
    },
    mainSection: {
        gridColumn: `2 / 3`
    }
});

interface AuthState {
    email: string,
    password: string,
}

const AuthorizationPage: React.FC = () => {
    const classes = useStyles();
    const {history} = useReactRouter();
    const dispatch = useDispatch();
    const userState = useSelector(userSelector);
    const [inputState, setInputState] = useState<AuthState>({email: "", password: ""});

    useEffect(() => {
        if (userState.user == null) {
            return
        }
        history.push(Routes.top())
    }, [userState]);

    const handleChange = (prop: keyof AuthState) => (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputState({...inputState, [prop]: event.target.value});
    };

    const handleSubmit = () => {
        const payload: authorizeUserPayload = {
            email: inputState.email,
            password: inputState.password
        };
        dispatch(signIn(payload))
    };


    return (
        <div className={classes.auth}>
            <div className={classes.mainSection}>
                <Card className={classes.card}>
                    <CardContent className={classes.cardContent}>
                        <TextField id='email'
                                   className={classes.textField}
                                   label='Email'
                                   value={inputState.email}
                                   onChange={handleChange('email')}
                        />
                        <TextField id='password'
                                   className={classes.textField}
                                   label='Password'
                                   type='password'
                                   onChange={handleChange('password')}
                        />
                        <Button variant="contained"
                                color={'primary'}
                                className={classes.button}
                                onClick={handleSubmit}>
                            SignIn
                        </Button>
                    </CardContent>
                </Card>
                <Card className={classes.card}>
                    <CardContent>
                        {userState.user === null ? (
                            <p>No User</p>
                        ) : (
                            <>
                                <p>{userState.user.id}</p>
                                <p>{userState.user.email}</p>
                            </>
                        )}
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default AuthorizationPage;
