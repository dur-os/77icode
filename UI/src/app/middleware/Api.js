import { normalize }	from 'normalizr';
import { camelizeKeys }	from 'humps';

const API_ROOT = 'http://127.0.0.1:8000';

function callApi(endpoint, formData) {
  const fullUrl = (endpoint.indexOf(API_ROOT) === -1) ? API_ROOT + endpoint : endpoint;
  const params = {};
  params.method = 'GET';
  if (typeof formData === 'object') {
    params.method = 'POST';
    if (formData instanceof FormData) {
      params.body = formData;
    } else {
      params.body = Object.keys(formData).map((key) => {
        if (formData.hasOwnProperty(key)) {
          const ukey = encodeURIComponent(key);
          const uval = encodeURIComponent(formData[key]);
          return `${ukey}=${uval}`;
        }
        return '';
      }).join('&');
      params.headers = {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
      };
    }
  }
  return fetch(fullUrl, params)
    .then(response =>
      response.json().then(json => ({ json, response }))
    ).then(({ json, response }) => {
      if (!response.ok) {
        return Promise.reject(json);
      }
      const camelizedJson = camelizeKeys(json);
      // const nextPageUrl = getNextPageUrl(response);
      return camelizedJson;
    });
}

// A Redux middleware that interprets actions with CALL_API info specified.
// Performs the call and promises when such actions are dispatched.
export default (CALL_API) => {
  const callFunc = store => next => action => {
    const callAPI = action[CALL_API];
    if (typeof callAPI === 'undefined') {
      return next(action);
    }

    let { endpoint } = callAPI;
    const { formData, schema, types } = callAPI;
    if (typeof endpoint === 'function') {
      endpoint = endpoint(store.getState());
    }

    if (typeof endpoint !== 'string') {
      throw new Error('Specify a string endpoint URL.');
    }
    if (!schema) {
      throw new Error('Specify one of the exported Schemas.');
    }
    if (!Array.isArray(types) || types.length !== 3) {
      throw new Error('Expected an array of three action types.');
    }
    if (!types.every(type => typeof type === 'string')) {
      throw new Error('Expected action types to be strings.');
    }

    function actionWith(data) {
      const finalAction = Object.assign({}, action, data);
      delete finalAction[CALL_API];
      return finalAction;
    }

    const [requestType, successType, failureType] = types;
    next(actionWith({ type: requestType }));

    return callApi(endpoint, formData).then(
      response => {
        if (response.code === 200) {
          return next(actionWith({
            data: normalize(response.data, schema),
            type: successType
          }));
        }
        return next(actionWith({
          error: response.msg,
          type: failureType
        }));
      },
      error => next(actionWith({
        type: failureType,
        error: error.message || 'Something bad happened'
      }))
    );
  };
  return callFunc;
};
