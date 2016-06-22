import React from 'react';
import ReactDom from 'react-dom';
import injectTapEventPlugin from 'react-tap-event-plugin';
import Root from './Router'; // Our custom react component
import './styles/main.less';

injectTapEventPlugin();

ReactDom.render(<Root />, document.getElementById('app'));
