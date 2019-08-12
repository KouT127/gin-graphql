import React, {useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {AppState} from "../store/store";
import {authorizeUserPayload, signIn} from "../reducers/userReducer";
import {Button, Card, CardContent, makeStyles, TextField} from "@material-ui/core";

const userSelector = (state: AppState) => state.userState;

const useStyles = makeStyles({
    card: {
        display: `flex`,
        justifyContent: `center`,
        width: `50%`,
    },
    cardContent: {
        width: `inherit`,
        display: `flex`,
        flexDirection: `column`
    }
});

interface AuthState {
    email: string,
    password: string,
}

const AuthorizationPage: React.FC = () => {
    const classes = useStyles();
    const dispatch = useDispatch();
    const userState = useSelector(userSelector);
    const [inputState, setInputState] = useState<AuthState>({email: "", password: ""});

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
        <>
            <Card className={classes.card}>
                <CardContent className={classes.cardContent}>
                    <TextField id='email'
                               label='Email'
                               value={inputState.email}
                               onChange={handleChange('email')}
                    />
                    <TextField id='password'
                               label='Password'
                               type='password'
                               onChange={handleChange('password')}
                    />
                    <Button variant="contained"
                            onClick={handleSubmit}>
                        SignIn
                    </Button>
                </CardContent>
            </Card>
            <Card>
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
        </>
    );
};

export default AuthorizationPage;
