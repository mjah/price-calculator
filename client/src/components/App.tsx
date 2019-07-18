import React from 'react';
import PriceCalculatorForm from './Form';
import Result from './Result';
import styled from 'styled-components';

const Wrapper = styled.section`
  margin: 0 auto;
`;

const Title = styled.h1`
  font-size: 1.5em;
`;

const App = () => (
  <Wrapper>
    <Title>Price Calculator</Title>
    <PriceCalculatorForm />
    <Result />
  </Wrapper>
)

export default App;
