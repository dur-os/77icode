import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { Router, Route, browserHistory } from 'react-router';
import Login from './views/Login';
import { initializeStore } from './store';
import MaterialThemes from './MaterialThemes';

const store = initializeStore();

class Root extends Component {

  static childContextTypes = {
    muiTheme: React.PropTypes.object
  }

  getChildContext() {
    return {
      muiTheme: MaterialThemes
    };
  }

  render() {
    return (
      <div>
        <Provider store={ store }>
          <Router history={ browserHistory }>
            <Route path="/" component={ Login } />
            <Route path="/login" component={ Login } />
          </Router>
        </Provider>
      </div>
    );
  }
}

export default Root;
