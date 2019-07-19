export const RESULTS_REQUESTED = 'RESULTS_REQUESTED';
export const RESULTS_LOADED = 'RESULTS_LOADED';
export const RESULTS_CLEAR = 'RESULTS_CLEAR';
export const API_ERRORED = 'API_ERRORED';

export function getResults(payload: any) {
  return { type: RESULTS_REQUESTED, payload };
}

export function clearResults() {
  return { type: RESULTS_CLEAR };
}
