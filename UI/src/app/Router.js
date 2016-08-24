import React	from 'react';
import { Route }	from 'react-router';
import Login	from './views/Login';

export default (
  <Route path="/" component={ Login }>
    <Route
      path="/login"
      component={ Login }
    />
  </Route>
);
