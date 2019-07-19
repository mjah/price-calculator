import React, { useState } from 'react';
import { connect } from 'react-redux';
import { clearResults } from '../../store/pricecalculator/actions'

const Result = ({ results, error, clearResults }: any) => {
  const [toFixedDecimalPlace, setDecimalPlace] = useState(2)

  return (
    <>
      {error ? (
        <div>
          <h2>Error</h2>
          <span>{String(error)}</span>
        </div>
      ):(<></>)}
      {results.length !== 0 ? (
        <div>
          <h2>All Results</h2>
          <button type="button" onClick={() => {clearResults()}}>
            Clear All Results
          </button>
          <button type="button" onClick={() => { setDecimalPlace( toFixedDecimalPlace > 0 ? toFixedDecimalPlace-1 : 0 )}}
            >
            {toFixedDecimalPlace > 0 ? toFixedDecimalPlace-1 : 0} d.p.
          </button>
          <button type="button" onClick={() => {setDecimalPlace(toFixedDecimalPlace+1)}}>
            {toFixedDecimalPlace+1} d.p.
          </button>
          {results.slice(0).reverse().map((result: any, i: any) => (
            <div key={results.length-i}>
              <h3>#{results.length-i}</h3>
              <div>
                <strong>Submitted:</strong> {JSON.stringify(JSON.parse(result.submitted), null, 2)}
              </div>
              <div>
                <strong>Results:</strong>
                <div>Fees Total: <span>{result.getFeesTotal.toFixed(toFixedDecimalPlace)}</span>
                  <ul>
                    <li>Channel Fees Total: <span>{result.getChannelFeesTotal.toFixed(toFixedDecimalPlace)}</span></li>
                    <li>Payment Fees Total: <span>{result.getPaymentFeesTotal.toFixed(toFixedDecimalPlace)}</span></li>
                    <li>Other Fees Total: <span>{result.getOtherFeesTotal.toFixed(toFixedDecimalPlace)}</span></li>
                    <li>Sales Tax Fees Total: <span>{result.getSalesTaxFeesTotal.toFixed(toFixedDecimalPlace)}</span></li>
                  </ul>
                </div>
                <div>Profit Total: <span>{result.getProfitTotal.toFixed(toFixedDecimalPlace)}</span></div>
                {result.isValidProfitRate ? (
                  <div>Sell Price by Profit Rate: <span>{result.getSellPriceByProfitRate.toFixed(toFixedDecimalPlace)}</span></div>
                ) : (
                  <div>Invalid Profit Rate.</div>
                )}
              </div>
            </div>
          ))}
        </div>
      ):(<></>)}
    </>
  )
}

const mapStateToProps = (state: any) => ({
  results: state.results,
  error: state.error
})

const mapDispatchToProps = (dispatch: any) => ({
  clearResults: () => dispatch(clearResults())
})

export default connect(mapStateToProps, mapDispatchToProps)(Result);
