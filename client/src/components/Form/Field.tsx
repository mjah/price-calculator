import React from 'react';
import { Field } from 'react-final-form';
import styled from 'styled-components';

const FieldStyle = styled.label`
  display: block;
  clear: both;
  margin-bottom: 0.5em;
`;

const Label = styled.label`
  display: block;
  clear: both;
  margin: 0 1em;
  font-weight: bold;
`;

const Input = styled.input`
  padding: 0.2em 0.4em;
  border: unset;
  border-radius: 0.25em;
  font-weight: bold;
  max-width: 80%;
`;

const InputField = ({ fieldDetails }: any) => {
  switch (fieldDetails.type) {
    case 'checkbox': {
      return (
        <Field name={fieldDetails.id} type={fieldDetails.type} validate={fieldDetails.validate}>
          {({ input }) => (
            <FieldStyle>
              <Label>{fieldDetails.label}</Label>
              <Input {...input} type={fieldDetails.type} />
            </FieldStyle>
          )}
        </Field>
      );
    }
    default: {
      return (
        <Field name={fieldDetails.id} validate={fieldDetails.validate}>
          {({ input, meta }) => (
            <FieldStyle>
              <Label>{fieldDetails.label}</Label>
              <Input {...input} type="number" placeholder="0.00" />
              {meta.error && meta.touched && <span>{meta.error}</span>}
            </FieldStyle>
          )}
        </Field>
      );
    }
  }
};

export default InputField;
