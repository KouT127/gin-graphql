import React, {useState} from "react";
import {Button, Card, CardContent, makeStyles, TextField} from "@material-ui/core";
import useReactRouter from 'use-react-router';
import {useDispatch, useSelector} from "react-redux";
import {signIn, userSelector} from "../reducers/UserReducer";

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
    const {user} = useSelector(userSelector);
    const dispatch = useDispatch();
    const {history} = useReactRouter();
    const [inputState, setInputState] = useState<AuthState>({email: "", password: ""});

    const handleChange = (prop: keyof AuthState) => (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputState({...inputState, [prop]: event.target.value});
    };

    const handleSubmit = () => {
        dispatch(signIn({email: inputState.email, password: inputState.password}))
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
                        {user === null ? (
                            <p>No User</p>
                        ) : (
                            <>
                                <p>{user.id}</p>
                                <p>{user.email}</p>
                            </>
                        )}
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default AuthorizationPage;
