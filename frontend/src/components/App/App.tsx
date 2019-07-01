import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { AppState } from "../../store/store";
import { loadUser } from "../../reducers/userReducer";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import styled from "styled-components";

const userSelector = (state: AppState) => state.users;

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

const Main =styled.main`
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

  useEffect(() => {
    dispatch(loadUser());
  }, [dispatch]);


  return (
    <Router>
      <Switch>

        <div>
          <Header>
            <Wrapper>
              <Title>TITLE</Title>
            </Wrapper>
          </Header>
          <Main>
            <Row>
              <Column>
                <h1>Description</h1>
                <ButtonSection>
                  <Button>SIGNIN</Button>
                  <Button>SIGNUP</Button>
                </ButtonSection>
              </Column>
            </Row>
          </Main>
        </div>
      </Switch>
    </Router>
  );
};

export default App;
