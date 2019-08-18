import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import './nomalize.css'
import App from './components/App/App';
import * as serviceWorker from './serviceWorker';
import {store} from "./store/Store";
import {Provider} from "react-redux";


ReactDOM.render(
    <Provider store={store as any}>
        <App/>
    </Provider>,
    document.getElementById("root")
);

serviceWorker.unregister();
