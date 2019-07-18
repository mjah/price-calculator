export const required = (value: any) => value ? undefined : 'Required'

export const mustBeNumber = (value: any) => isNaN(value) ? 'Must be a number' : undefined

export const mustBeBool = (value: any) => typeof value === 'boolean' ? 'Must be yes or no' : undefined

export const composeValidators = (...validators: any) => (value: any) =>
  validators.reduce((error: any, validator: any) => error || validator(value), undefined)
