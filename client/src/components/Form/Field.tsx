import React from 'react'
import { Field } from 'react-final-form'

const InputField = ({ fieldDetails }: any) => {
  switch(fieldDetails.type) {
    case "checkbox": {
      return (
        <Field name={fieldDetails.id} type={fieldDetails.type} validate={fieldDetails.validate}>
          {({ input }) => (
            <div>
              <label>{fieldDetails.label}</label>
              <input {...input} type={fieldDetails.type} />
            </div>
          )}
        </Field>
      )
    }
    default: {
      return (
        <Field name={fieldDetails.id} validate={fieldDetails.validate}>
          {({ input, meta }) => (
            <div>
              <label>{fieldDetails.label}</label>
              <input {...input} type="number" placeholder="0.00" />
              {meta.error && meta.touched && <span>{meta.error}</span>}
            </div>
          )}
        </Field>
      )
    }
  }
}

export default InputField
