import React, {useEffect, useState} from "react";
import {Button, Card, CardContent, makeStyles, TextField} from "@material-ui/core";
import useReactRouter from 'use-react-router';
import {useUserState} from "../components/Providers/UserProvider";

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
    const {userState, authConnect, signIn} = useUserState();
    const [inputState, setInputState] = useState<AuthState>({email: "", password: ""});

    useEffect(() => {
        authConnect()
    }, []);

    useEffect(() => {
        if (!userState || !userState.user) {
            return
        }
        // history.push(Routes.top())
    }, [history, userState]);

    const handleChange = (prop: keyof AuthState) => (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputState({...inputState, [prop]: event.target.value});
    };

    const handleSubmit = () => {
        signIn(inputState.email, inputState.password)
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
                        {userState === null ? (
                            <p>No User</p>
                        ) : (
                            <>
                                <p>{!userState.user ? '' : userState.user.id}</p>
                                <p>{!userState.user ? '' : userState.user.email}</p>
                            </>
                        )}
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default AuthorizationPage;
