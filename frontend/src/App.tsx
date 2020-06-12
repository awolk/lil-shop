import React from "react";
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Items from "./Items";

const client = new ApolloClient();

function App() {
  return (
    <ApolloProvider client={client}>
      <Items />
    </ApolloProvider>
  );
}

export default App;
