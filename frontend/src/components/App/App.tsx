import React, {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {AppState} from "../../store/store";
import {authorizeUserPayload, connectAuth, signIn} from "../../reducers/userReducer";
import {BrowserRouter as Router, Link, Route, Switch} from "react-router-dom";
import styled from "styled-components";
import AuthorizationPage from "../../pages/AuthorizationPage";

const userSelector = (state: AppState) => state.userState;

const Header = styled.header`
  display: -webkit-flex;
  display: flex;
  justify-content: center;
  background-color: #f2f2f4;
  height: 60px;
`;

const Wrapper = styled.div`
  width: 50%
  height: inherit;
`;

const Title = styled.a`
  padding: 15px;
  font-size: 1.5em;
  color: black;
  display: inline-block;
  vertical-align:top;
`;

const Main = styled.main`
  height: 600px;
`;

const Column = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: inherit;
`;

const Row = styled.div`
  display: flex;
  flex-direction: Row;
  justify-content: center;
  height: inherit;
`;

const ButtonSection = styled.div`
  display: flex;
  align-items: stretch;
`;

const Button = styled.a`
  min-width: 100px
`;

const App: React.FC = () => {
    const dispatch = useDispatch();
    const userState = useSelector(userSelector);
    const [inputState, setInputState] = useState({email: '', password: ''});

    useEffect(() => {
        dispatch(connectAuth());
    }, [dispatch]);

    const handleEmail = (e: React.FormEvent<HTMLInputElement>) => {
        e.preventDefault();
        setInputState({password: inputState.password, email: e.currentTarget.value})
    };
    const handlePassword = (e: React.FormEvent<HTMLInputElement>) => {
        e.preventDefault();
        setInputState({password: e.currentTarget.value, email: inputState.email})
    };
    const handleSubmit = () => {
        console.log('test')
        const payload: authorizeUserPayload = {
            email: inputState.email,
            password: inputState.password
        };
        dispatch(signIn(payload))
    };

    return (
        <Router>
            <ul>
                <li>
                    <Link to="/auth/">Auth</Link>
                </li>
            </ul>
            {/*<Switch>*/}
            {/*    <div>*/}
            {/*        <Header>*/}
            {/*            <Wrapper>*/}
            {/*                <Title>TITLE</Title>*/}
            {/*            </Wrapper>*/}
            {/*        </Header>*/}
            {/*        <Main>*/}
            {/*            <Row>*/}
            {/*                <form>*/}
            {/*                    <input onChange={handleEmail}></input>*/}
            {/*                    <input onChange={handlePassword}/>*/}
            {/*                </form>*/}

            {/*                <button onClick={handleSubmit}>Submit</button>*/}
            {/*            </Row>*/}
            {/*        </Main>*/}
            {/*    </div>*/}
            {/*</Switch>*/}
            <Route path="/auth/" exact component={AuthorizationPage} />
        </Router>
    );
};

export default App;
