import thunk	from 'redux-thunk';
import createLogger	from 'redux-logger';
import { createStore, applyMiddleware, combineReducers }	from 'redux';
import { routerReducer as routing }	from 'react-router-redux';
import api	from '../middleware/Api';
import { CALL_API }	from '../domain';
// Reducers
// import appReducer from '../reducers/reducers.js';

// Actions
// import { setViewport } from 'viewport.js';

const logger = createLogger();

const rootReducer = combineReducers({
  routing
});

export default function configureStore(preloadedState) {
  return createStore(
    rootReducer,
    preloadedState,
    applyMiddleware(logger, thunk, api(CALL_API))
  );
}
