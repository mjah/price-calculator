import React from 'react'
import { Form } from 'react-final-form'
import InputField from './Field'

const mustBeNumber = (value: any) =>
  (isNaN(value) && typeof value !== 'undefined') ? 'Must be a number' : undefined

const mustBeBool = (value: any) =>
  (typeof value !== 'boolean' && typeof value !== 'undefined') ? 'Must be true or false' : undefined

const fields = [
  {
    id: "sell_price",
    label: "Sell Price",
    validate: mustBeNumber
  },
  {
    id: "free_delivery_price",
    label: "Free Delivery Price",
    validate: mustBeNumber
  },
  {
    id: "cost",
    label: "Cost",
    validate: mustBeNumber
  },
  {
    id: "fees.sales_tax.rate",
    label: "Sales Tax Fee Rate",
    validate: mustBeNumber
  },
  {
    id: "fees.payment.rate",
    label: "Payment Fee Rate",
    validate: mustBeNumber
  },
  {
    id: "fees.payment.fixed",
    label: "Payment Fixed Fee",
    validate: mustBeNumber
  },
  {
    id: "fees.channel.rate",
    label: "Channel Fee Rate",
    validate: mustBeNumber
  },
  {
    id: "fees.channel.fixed",
    label: "Channel Fixed Fee",
    validate: mustBeNumber
  },
  {
    id: "fees.channel.is_capped",
    label: "Channel Fee Is Capped",
    type: "checkbox",
    validate: mustBeBool
  },
  {
    id: "fees.channel.capped_value",
    label: "Channel Fee Capped Value",
    validate: mustBeNumber
  },
  {
    id: "fees.other.rate",
    label: "Other Fee Rate",
    validate: mustBeNumber
  },
  {
    id: "fees.other.fixed",
    label: "Other Fixed Fee",
    validate: mustBeNumber
  },
  {
    id: "profit.rate",
    label: "Select Profit Rate",
    validate: mustBeNumber
  }
]

const stringToFloat = (name: any, val: any) => {
  return typeof val === 'string' ? parseFloat(val) : val
}

const onSubmit = (values: any) => {
  window.alert(JSON.stringify(values, stringToFloat, 2))
}

const PriceCalculatorForm = () => (
  <Form
    onSubmit={onSubmit}
    render={({ handleSubmit, form, submitting, pristine }) => (
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
      </form>
    )}
  />
)

export default PriceCalculatorForm;
