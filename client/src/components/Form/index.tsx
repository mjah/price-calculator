import React from 'react'
import { Form } from 'react-final-form'
import * as vd from './validation'
import InputField from './Field'

const fields = [
  {
    id: "sellPrice",
    label: "Sell Price",
    validate: vd.mustBeNumber
  },
  {
    id: "cost",
    label: "Cost",
    validate: vd.mustBeNumber
  },
  {
    id: "salesTaxFeeRate",
    label: "Sales Tax Fee Rate",
    validate: vd.mustBeNumber
  },
  {
    id: "paymentFeeRate",
    label: "Payment Fee Rate",
    validate: vd.mustBeNumber
  },
  {
    id: "paymentFixedFee",
    label: "Payment Fixed Fee",
    validate: vd.mustBeNumber
  },
  {
    id: "channelFeeRate",
    label: "Channel Fee Rate",
    validate: vd.mustBeNumber
  },
  {
    id: "channelFixedFee",
    label: "Channel Fixed Fee",
    validate: vd.mustBeNumber
  },
  {
    id: "channelFeeIsCapped",
    label: "Channel Fee Is Capped",
    validate: vd.mustBeNumber
  },
  {
    id: "channelFeeCappedValue",
    label: "Channel Fee Capped Value",
    type: "yesno_dropdown",
    validate: vd.mustBeBool
  },
  {
    id: "otherFeeRate",
    label: "Other Fee Rate",
    validate: vd.mustBeNumber
  },
  {
    id: "otherFixedFee",
    label: "Other Fixed Fee",
    validate: vd.mustBeNumber
  },
  {
    id: "selectProfitRate",
    label: "Select Profit Rate",
    validate: vd.mustBeNumber
  }
]

const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

const onSubmit = async (values: Promise<void>) => {
  await sleep(300)
  window.alert(JSON.stringify(values, null, 2))
}

const PriceCalculatorForm = () => (
  <Form
    onSubmit={onSubmit}
    render={({ handleSubmit, form, submitting, pristine, values }) => (
      <form onSubmit={handleSubmit}>
        {fields.map((field: any, i) => (
          <InputField key={i} fieldDetails={field} />
        ))}
        <div className="buttons">
          <button type="submit" disabled={submitting}>
            Submit
          </button>
          <button
            type="button"
            onClick={form.reset}
            disabled={submitting || pristine}
          >
            Reset
          </button>
        </div>
        <pre>{JSON.stringify(values, null, 2)}</pre>
      </form>
    )}
  />
)

export default PriceCalculatorForm;
