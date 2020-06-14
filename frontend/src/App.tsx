import React from "react";
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Shop from "./Shop";

const client = new ApolloClient();

function App() {
  return (
    <ApolloProvider client={client}>
      <Shop />
    </ApolloProvider>
  );
}

export default App;
