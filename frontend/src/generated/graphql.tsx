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

export type CartQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type CartQuery = (
  { __typename?: 'Query' }
  & { cart: (
    { __typename?: 'Cart' }
    & Pick<Cart, 'id'>
    & { lineItems: Array<(
      { __typename?: 'LineItem' }
      & Pick<LineItem, 'id' | 'quantity'>
      & { item: (
        { __typename?: 'Item' }
        & Pick<Item, 'id' | 'name' | 'costCents'>
      ) }
    )> }
  ) }
);

export type AddItemToCartMutationVariables = Exact<{
  itemID: Scalars['ID'];
  quantity: Scalars['Int'];
  cartID: Scalars['ID'];
}>;


export type AddItemToCartMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'addItemToCart'>
);

export type NewCartMutationVariables = Exact<{ [key: string]: never; }>;


export type NewCartMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'newCart'>
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
export const CartDocument = gql`
    query cart($id: ID!) {
  cart(id: $id) {
    id
    lineItems {
      id
      quantity
      item {
        id
        name
        costCents
      }
    }
  }
}
    `;

/**
 * __useCartQuery__
 *
 * To run a query within a React component, call `useCartQuery` and pass it any options that fit your needs.
 * When your component renders, `useCartQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCartQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useCartQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<CartQuery, CartQueryVariables>) {
        return ApolloReactHooks.useQuery<CartQuery, CartQueryVariables>(CartDocument, baseOptions);
      }
export function useCartLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<CartQuery, CartQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<CartQuery, CartQueryVariables>(CartDocument, baseOptions);
        }
export type CartQueryHookResult = ReturnType<typeof useCartQuery>;
export type CartLazyQueryHookResult = ReturnType<typeof useCartLazyQuery>;
export type CartQueryResult = ApolloReactCommon.QueryResult<CartQuery, CartQueryVariables>;
export const AddItemToCartDocument = gql`
    mutation addItemToCart($itemID: ID!, $quantity: Int!, $cartID: ID!) {
  addItemToCart(itemID: $itemID, quantity: $quantity, cartID: $cartID)
}
    `;
export type AddItemToCartMutationFn = ApolloReactCommon.MutationFunction<AddItemToCartMutation, AddItemToCartMutationVariables>;

/**
 * __useAddItemToCartMutation__
 *
 * To run a mutation, you first call `useAddItemToCartMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddItemToCartMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addItemToCartMutation, { data, loading, error }] = useAddItemToCartMutation({
 *   variables: {
 *      itemID: // value for 'itemID'
 *      quantity: // value for 'quantity'
 *      cartID: // value for 'cartID'
 *   },
 * });
 */
export function useAddItemToCartMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<AddItemToCartMutation, AddItemToCartMutationVariables>) {
        return ApolloReactHooks.useMutation<AddItemToCartMutation, AddItemToCartMutationVariables>(AddItemToCartDocument, baseOptions);
      }
export type AddItemToCartMutationHookResult = ReturnType<typeof useAddItemToCartMutation>;
export type AddItemToCartMutationResult = ApolloReactCommon.MutationResult<AddItemToCartMutation>;
export type AddItemToCartMutationOptions = ApolloReactCommon.BaseMutationOptions<AddItemToCartMutation, AddItemToCartMutationVariables>;
export const NewCartDocument = gql`
    mutation newCart {
  newCart
}
    `;
export type NewCartMutationFn = ApolloReactCommon.MutationFunction<NewCartMutation, NewCartMutationVariables>;

/**
 * __useNewCartMutation__
 *
 * To run a mutation, you first call `useNewCartMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useNewCartMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [newCartMutation, { data, loading, error }] = useNewCartMutation({
 *   variables: {
 *   },
 * });
 */
export function useNewCartMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<NewCartMutation, NewCartMutationVariables>) {
        return ApolloReactHooks.useMutation<NewCartMutation, NewCartMutationVariables>(NewCartDocument, baseOptions);
      }
export type NewCartMutationHookResult = ReturnType<typeof useNewCartMutation>;
export type NewCartMutationResult = ApolloReactCommon.MutationResult<NewCartMutation>;
export type NewCartMutationOptions = ApolloReactCommon.BaseMutationOptions<NewCartMutation, NewCartMutationVariables>;