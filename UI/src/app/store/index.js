import thunk	from 'redux-thunk';
import createLogger	from 'redux-logger';
import { createStore, applyMiddleware, compose }	from 'redux';
import api	from '../middleware/Api';
import { CALL_API }	from '../domain';
import rootReducer	from '../reducers';
import DevTools from '../containers/DevTools';
// Reducers
// import appReducer from '../reducers/reducers.js';

// Actions
// import { setViewport } from 'viewport.js';

const logger = createLogger();

export default function configureStore(preloadedState) {
  return createStore(
    rootReducer,
    preloadedState,
    compose(
         applyMiddleware(logger, thunk, api(CALL_API)),
         DevTools.instrument()
       )
  );
}
