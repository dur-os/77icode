import React, { Component, PropTypes }	from 'react';
import { Provider }	from 'react-redux';
import routes	from './Router';
import { Router }	from 'react-router';
import MaterialThemes	from './MaterialThemes';

export default class Root extends Component {
  static childContextTypes = {
    muiTheme: React.PropTypes.object
  }

  getChildContext() {
    return {
      muiTheme: MaterialThemes
    };
  }

  componentWillMount() {
   // loadData();
  }

  render() {
    const { store, history } = this.props;
    return (
      <Provider store={ store }>
        <Router history={ history } routes={ routes } />
      </Provider>
    );
  }
}

Root.propTypes = {
  store: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired
};
