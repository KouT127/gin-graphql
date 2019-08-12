import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import App from './components/App/App';
import * as serviceWorker from './serviceWorker';
import {Provider} from "react-redux";
import {store} from "./store/store";


ReactDOM.render(
    <Provider store={store as any}>
        <App/>
    </Provider>,
    document.getElementById("root")
);

serviceWorker.unregister();
