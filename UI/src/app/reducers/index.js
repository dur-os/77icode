import { combineReducers }	from 'redux';
import { routerReducer as routing }	from 'react-router-redux';

function errorMessage(state = null, action) {
  const { error } = action;
  if (error) {
    return action.error;
  }
  return null;
}

const rootReducer = combineReducers({
  errorMessage,
  routing
});

export default rootReducer;
