import { RESULTS_LOADED, API_ERRORED, RESULTS_CLEAR } from './actions';

const initialState = {
  results: [],
  error: false
};

function reducer(state = initialState, action: any) {
  if (action.type === RESULTS_LOADED) {
    return Object.assign({}, state, {
      results: state.results.concat(action.payload),
      error: false
    });
  }
  if (action.type === RESULTS_CLEAR) {
    return Object.assign({}, state, initialState);
  }
  if (action.type === API_ERRORED) {
    return Object.assign({}, state, {
      error: action.payload
    });
  }
  return state;
}

export default reducer;
