import React from 'react';
import styled from 'styled-components';

const ButtonStyle = styled.button`
  margin: 0.5em;
  padding: 0.25em 1em;
  font-weight: bold;
  color: #fff;
  background-color: #777;
  border: unset;
  border-radius: 3px;

  &:hover {
    background-color: #333;
  }

  &:active {
    background-color: #000;
  }

  &:focus {
    outline: 0;
  }

  &:disabled {
    color: #ccc;
    background-color: #333;
  }
`;

const Button = ({ children, ...rest }: any) => {
  return <ButtonStyle {...rest}>{children}</ButtonStyle>;
};

export default Button;
