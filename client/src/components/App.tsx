import React from 'react';
import PriceCalculatorForm from './Form';
import Result from './Result';
import styled from 'styled-components';

const Wrapper = styled.section`
  margin: 0 auto;
  padding: 1em;
  max-width: 1024px;
  text-align: center;
`;

const Title = styled.h1`
  margin: 0.5em 0;
  font-size: 2em;
`;

const Split30 = styled.section`
  display: block;
  float: left;
  width: 30%;
  @media (max-width: 768px) {
    width: 100%;
  }
`;

const Split70 = styled.section`
  display: block;
  float: left;
  width: 70%;
  @media (max-width: 768px) {
    width: 100%;
  }
`;

const App = () => (
  <Wrapper>
    <Title>Price Calculator</Title>
    <Split30>
      <PriceCalculatorForm />
    </Split30>
    <Split70>
      <Result />
    </Split70>
  </Wrapper>
);

export default App;
