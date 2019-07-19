import React from 'react';
import { connect } from 'react-redux';
import { Form } from 'react-final-form';
import styled from 'styled-components';
import InputField from './Field';
import { getResults } from '../../store/pricecalculator/actions';
import Button from '../Button';

const FormStyle = styled.form`
  padding: 1em 0;
  margin: 0.5em;
  background-color: #111;
  border-radius: 0.75em;
`;

const mustBeNumber = (value: any) =>
  isNaN(value) && typeof value !== 'undefined' ? 'Must be a number' : undefined;

const mustBeBool = (value: any) =>
  typeof value !== 'boolean' && typeof value !== 'undefined' ? 'Must be true or false' : undefined;

const fields = [
  {
    id: 'sell_price',
    label: 'Sell Price',
    validate: mustBeNumber
  },
  {
    id: 'free_delivery_price',
    label: 'Free Delivery Price',
    validate: mustBeNumber
  },
  {
    id: 'cost',
    label: 'Cost',
    validate: mustBeNumber
  },
  {
    id: 'fees.sales_tax.rate',
    label: 'Sales Tax Fee Rate',
    validate: mustBeNumber
  },
  {
    id: 'fees.payment.rate',
    label: 'Payment Fee Rate',
    validate: mustBeNumber
  },
  {
    id: 'fees.payment.fixed',
    label: 'Payment Fixed Fee',
    validate: mustBeNumber
  },
  {
    id: 'fees.channel.rate',
    label: 'Channel Fee Rate',
    validate: mustBeNumber
  },
  {
    id: 'fees.channel.fixed',
    label: 'Channel Fixed Fee',
    validate: mustBeNumber
  },
  {
    id: 'fees.channel.is_capped',
    label: 'Channel Fee Is Capped',
    type: 'checkbox',
    validate: mustBeBool
  },
  {
    id: 'fees.channel.capped_value',
    label: 'Channel Fee Capped Value',
    validate: mustBeNumber
  },
  {
    id: 'fees.other.rate',
    label: 'Other Fee Rate',
    validate: mustBeNumber
  },
  {
    id: 'fees.other.fixed',
    label: 'Other Fixed Fee',
    validate: mustBeNumber
  },
  {
    id: 'profit.rate',
    label: 'Select Profit Rate',
    validate: mustBeNumber
  }
];

const stringToFloatReplacer = (name: any, val: any) => {
  return typeof val === 'string' ? parseFloat(val) : val;
};

const PriceCalculatorForm = ({ getResults }: any) => {
  const onSubmit = (values: any) => {
    let bodyData = JSON.stringify(values, stringToFloatReplacer, 2);
    getResults(bodyData);
  };

  return (
    <Form
      onSubmit={onSubmit}
      render={({ handleSubmit, form, submitting, pristine }) => (
        <FormStyle onSubmit={handleSubmit}>
          {fields.map((field: any, i) => (
            <InputField key={i} fieldDetails={field} />
          ))}
          <div className="buttons">
            <Button type="submit" disabled={submitting}>
              Submit
            </Button>
            <Button type="button" onClick={form.reset} disabled={submitting || pristine}>
              Reset
            </Button>
          </div>
        </FormStyle>
      )}
    />
  );
};

const mapDispatchToProps = (dispatch: any) => ({
  getResults: (payload: any) => dispatch(getResults(payload))
});

export default connect(
  null,
  mapDispatchToProps
)(PriceCalculatorForm);
