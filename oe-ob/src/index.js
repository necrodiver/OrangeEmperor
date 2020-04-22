import React from 'react';
import ReactDOM from 'react-dom';
import * as Sentry from '@sentry/browser';
import App from './Components/App/App';
import { BrowserRouter,HashRouter } from 'react-router-dom';
import './index.less';

Sentry.init({dsn: "https://99cb561cc36a45d6940aab6961b3489f@o381675.ingest.sentry.io/5209347"});

ReactDOM.render(<HashRouter><App /></HashRouter>, document.getElementById('root'));
