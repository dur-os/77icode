import { Schema }	from 'normalizr';


// Action key that carries API call info interpreted by this Redux middleware.
export const CALL_API = Symbol('Call API');

const BaseSchema = new Schema('BaseSchema');

const userSchema = new Schema('users', {
  idAttribute: 'id'
});

// Schemas for Github API responses.
export default {
  BaseSchema,
  USER: userSchema
};
