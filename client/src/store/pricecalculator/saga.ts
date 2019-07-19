import { takeEvery, call, put } from 'redux-saga/effects';
import { RESULTS_REQUESTED, RESULTS_LOADED, API_ERRORED } from './actions';

export default function* watcherSaga() {
  yield takeEvery(RESULTS_REQUESTED, workerSaga);
}

function* workerSaga(action: any) {
  try {
    const payload = yield call(getResults, action.payload);
    payload.submitted = action.payload;
    yield put({ type: RESULTS_LOADED, payload });
  } catch (e) {
    yield put({ type: API_ERRORED, payload: e });
  }
}

function getResults(bodyData: any) {
  let priceCalcApi = 'http://localhost:9096/v1/all_results'
  if (process.env.REACT_APP_PRICECALCULATOR_API) {
    priceCalcApi = process.env.REACT_APP_PRICECALCULATOR_API + '/all_results'
  }
  return fetch(priceCalcApi, {
    method: 'post',
    headers: { 'Content-Type': 'application/json' },
    body: bodyData
  }).then((response) => response.json());
}
