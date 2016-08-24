import React	from 'react';
import { render }	from 'react-dom';
import injectTapEventPlugin from 'react-tap-event-plugin';
import { browserHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import configureStore from './store';
import Root from './Root';
import './styles/main.less';
const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

injectTapEventPlugin();

render(<Root store={ store } history={ history } />,
	document.getElementById('app'));
