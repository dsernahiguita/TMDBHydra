import React from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter as Router } from 'react-router-dom';
import ReactDOM from 'react-dom';
import 'typeface-roboto';
import * as serviceWorker from './serviceWorker';
import App from './App';
import { Header } from './pages/components/header'
import store from './store';
import './index.css';

/* The default language is needed in case that the browser language is not defined
into the messages object */
// const defaultLanguage = 'es';
// const language = defaultLanguage


const rootElement = document.getElementById('root');
ReactDOM.render(
  <Provider store={store}>
    <Header />
    <Router>
      <App />
    </Router>
  </Provider>,
  rootElement
);

serviceWorker.unregister();
