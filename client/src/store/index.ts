import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';
import { composeWithDevTools } from 'redux-devtools-extension';
import reducer from './pricecalculator/reducers';
import saga from './pricecalculator/saga';

const initialiseSagaMiddleware = createSagaMiddleware();

const store = createStore(reducer, composeWithDevTools(applyMiddleware(initialiseSagaMiddleware)));

initialiseSagaMiddleware.run(saga);

export default store;
