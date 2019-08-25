import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import App from './components/App';
import * as serviceWorker from './serviceWorker';
// @ts-ignore
import {store} from "./store/Store";
import {Provider} from "react-redux";
import {CssBaseline} from "@material-ui/core";


ReactDOM.render(
    <Provider store={store as any}>
        <CssBaseline/>
        <App/>
    </Provider>,
    document.getElementById("root")
);

serviceWorker.unregister();
