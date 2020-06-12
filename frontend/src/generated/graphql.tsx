import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/react-common';
import * as ApolloReactHooks from '@apollo/react-hooks';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Item = {
  __typename?: 'Item';
  id: Scalars['ID'];
  name: Scalars['String'];
  costCents: Scalars['Int'];
};

export type LineItem = {
  __typename?: 'LineItem';
  id: Scalars['ID'];
  item: Item;
  quantity: Scalars['Int'];
};

export type Cart = {
  __typename?: 'Cart';
  id: Scalars['ID'];
  lineItems: Array<LineItem>;
};

export type Query = {
  __typename?: 'Query';
  items: Array<Item>;
  cart: Cart;
};


export type QueryCartArgs = {
  id: Scalars['ID'];
};

export type Mutation = {
  __typename?: 'Mutation';
  newCart: Scalars['ID'];
  addItemToCart: Scalars['ID'];
};


export type MutationAddItemToCartArgs = {
  itemID: Scalars['ID'];
  quantity: Scalars['Int'];
  cartID: Scalars['ID'];
};

export type AllItemsQueryVariables = Exact<{ [key: string]: never; }>;


export type AllItemsQuery = (
  { __typename?: 'Query' }
  & { items: Array<(
    { __typename?: 'Item' }
    & Pick<Item, 'id' | 'name' | 'costCents'>
  )> }
);


export const AllItemsDocument = gql`
    query allItems {
  items {
    id
    name
    costCents
  }
}
    `;

/**
 * __useAllItemsQuery__
 *
 * To run a query within a React component, call `useAllItemsQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllItemsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllItemsQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllItemsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<AllItemsQuery, AllItemsQueryVariables>) {
        return ApolloReactHooks.useQuery<AllItemsQuery, AllItemsQueryVariables>(AllItemsDocument, baseOptions);
      }
export function useAllItemsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AllItemsQuery, AllItemsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<AllItemsQuery, AllItemsQueryVariables>(AllItemsDocument, baseOptions);
        }
export type AllItemsQueryHookResult = ReturnType<typeof useAllItemsQuery>;
export type AllItemsLazyQueryHookResult = ReturnType<typeof useAllItemsLazyQuery>;
export type AllItemsQueryResult = ApolloReactCommon.QueryResult<AllItemsQuery, AllItemsQueryVariables>;