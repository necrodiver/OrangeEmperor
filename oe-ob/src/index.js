import React from 'react';
import ReactDOM from 'react-dom';
import App from './Components/App/App';
import { BrowserRouter,HashRouter } from 'react-router-dom';
import './index.less';

ReactDOM.render(<HashRouter><App /></HashRouter>, document.getElementById('root'));
