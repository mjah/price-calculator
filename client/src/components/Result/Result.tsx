import React, { useState } from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';
import { clearResults } from '../../store/pricecalculator/actions';
import Button from '../Button';

const SubTitle = styled.h2`
  margin: 0;
  font-size: 1.5em;
`;

const Box = styled.section`
  padding: 1em 0;
  margin: 0.5em;
  background-color: #111;
  border-radius: 0.75em;
`;

const Result = ({ results, error, clearResults }: any) => {
  const [toFixedDecimalPlace, setDecimalPlace] = useState(2);

  return (
    <>
      {error ? (
        <Box>
          <SubTitle>Error</SubTitle>
          <span>{String(error)}</span>
        </Box>
      ) : (
        <></>
      )}
      {results.length !== 0 ? (
        <Box>
          <SubTitle>All Results</SubTitle>
          <Button
            type="button"
            onClick={() => {
              clearResults();
            }}
          >
            Clear All Results
          </Button>
          <Button
            type="button"
            onClick={() => {
              setDecimalPlace(toFixedDecimalPlace > 0 ? toFixedDecimalPlace - 1 : 0);
            }}
            disabled={toFixedDecimalPlace === 0}
          >
            {toFixedDecimalPlace > 0 ? toFixedDecimalPlace - 1 : '-'} d.p.
          </Button>
          <Button
            type="button"
            onClick={() => {
              setDecimalPlace(toFixedDecimalPlace + 1);
            }}
          >
            {toFixedDecimalPlace + 1} d.p.
          </Button>
          {results
            .slice(0)
            .reverse()
            .map((result: any, i: any) => (
              <div key={results.length - i}>
                <h3>
                  ---
                  <br />#{results.length - i}
                </h3>
                <strong>Submitted:</strong>{' '}
                <div>{JSON.stringify(JSON.parse(result.submitted), null, 2)}</div>
                <br />
                <strong>Results:</strong>
                <div>
                  Fees Total: <span>{result.getFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                <div>
                  Channel Fees Total:{' '}
                  <span>{result.getChannelFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                <div>
                  Payment Fees Total:{' '}
                  <span>{result.getPaymentFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                <div>
                  Other Fees Total:{' '}
                  <span>{result.getOtherFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                <div>
                  Sales Tax Fees Total:{' '}
                  <span>{result.getSalesTaxFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                <div>
                  Profit Total: <span>{result.getProfitTotal.toFixed(toFixedDecimalPlace)}</span>
                </div>
                {result.isValidProfitRate ? (
                  <div>
                    Sell Price by Profit Rate:{' '}
                    <span>{result.getSellPriceByProfitRate.toFixed(toFixedDecimalPlace)}</span>
                  </div>
                ) : (
                  <div>Invalid Profit Rate.</div>
                )}
              </div>
            ))}
        </Box>
      ) : (
        <></>
      )}
    </>
  );
};

const mapStateToProps = (state: any) => ({
  results: state.results,
  error: state.error
});

const mapDispatchToProps = (dispatch: any) => ({
  clearResults: () => dispatch(clearResults())
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Result);
