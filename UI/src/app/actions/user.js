import Schemas, { CALL_API }	from '../domain';

export const USER_INFO_REQUEST	= 'USER_INFO_REQUEST';
export const USER_INFO_SUCCESS	= 'USER_INFO_SUCCESS';
export const USER_INFO_FAILURE	= 'USER_INFO_FAILURE';


// Fetches a single user from Github API.
// Relies on the custom API middleware defined in ../middleware/api.js.
export function getUser() {
  return {
    [CALL_API]: {
      types: [USER_INFO_REQUEST, USER_INFO_SUCCESS, USER_INFO_FAILURE],
      endpoint: '/admin/getUser',
      schema: Schemas.USER
    }
  };
}

export const USER_LOGIN_REQUEST	= 'USER_LOGIN_REQUEST';
export const USER_LOGIN_SUCCESS	= 'USER_LOGIN_SUCCESS';
export const USER_LOGIN_FAILURE	= 'USER_LOGIN_FAILURE';
// Fetches a single user from Github API.
// Relies on the custom API middleware defined in ../middleware/api.js.
function login(userName, passWord) {
  console.log('test:', userName, passWord);
  return {
    [CALL_API]: {
      types: [USER_LOGIN_REQUEST, USER_LOGIN_SUCCESS, USER_LOGIN_FAILURE],
      endpoint: '/admin/login',
      formData: { userName, passWord },
      schema: Schemas.USER
    }
  };
}

export function loginUser(userName, passWord) {
  console.log('test1:', userName, passWord);
  return (dispatch, getState) => {
    console.log(getState());
    return dispatch(login(userName, passWord));
  };
}
