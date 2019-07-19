import React from 'react';
import PriceCalculatorForm from './Form';
import Result from './Result';
import styled from 'styled-components';

const Wrapper = styled.section`
  margin: 0 auto;
  padding: 1em;
  max-width: 1000px;
  text-align: center;
`;

const Title = styled.h1`
  margin: 0.5em 0;
  font-size: 2em;
`;

const App = () => (
  <Wrapper>
    <Title>Price Calculator</Title>
    <PriceCalculatorForm />
    <Result />
  </Wrapper>
);

export default App;
