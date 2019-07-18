import React from 'react'
import { Field } from 'react-final-form'

const InputField = ({ fieldDetails }: any) => {
  switch(fieldDetails.type) {
    case "yesno_dropdown": {
      return (
        <Field name={fieldDetails.id} validate={fieldDetails.validate}>
          {({ input, meta }) => (
            <div>
              <label>{fieldDetails.label}</label>
              <select name={input.name} onChange={(value) => input.onChange(value)}>
                <option value="false">No</option>
                <option value="true">Yes</option>
              </select>
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
              <input {...input} type="text" placeholder={fieldDetails.label} />
              {meta.error && meta.touched && <span>{meta.error}</span>}
            </div>
          )}
        </Field>
      )
    }
  }
}

export default InputField
