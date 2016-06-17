import React, { Component, PropTypes } from 'react';
import { Provider } from 'react-redux';
import { Router, Route , browserHistory } from 'react-router';
import Login from './views/Login';
import { initializeStore } from './store';
import withMaterialUI from './decorators/withMaterialUI';

const store = initializeStore();


@withMaterialUI
export default class Root extends Component {
  render() {
    return (
        <div>
          <Provider store={store}>
            <Router history={browserHistory}>
              <Route path='/' component={Login} />
              <Route path='/login' component={Login}/>
            </Router>
          </Provider>
        </div>
    );
  }
};
