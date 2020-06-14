import React, { useEffect, useState } from "react";
import {
  useNewCartMutation,
  useCartLazyQuery,
  useAddItemToCartMutation,
} from "./generated/graphql";
import { ApolloError } from "apollo-boost";
import CartView from "./CartView";
import Items from "./Items";

export function useCart() {
  const [newCart] = useNewCartMutation();
  const [newCartError, setNewCartError] = useState<ApolloError | null>(null);
  const [fetchCart, { data: cart, error, refetch }] = useCartLazyQuery();

  const cartID = localStorage.getItem("cartID");

  useEffect(() => {
    if (!cartID || error) {
      newCart()
        .then(({ data }) => {
          const cartID = data!.newCart;
          localStorage.setItem("cartID", cartID);
          fetchCart({ variables: { id: cartID } });
        })
        .catch((err: ApolloError) => setNewCartError(err));
    } else {
      fetchCart({ variables: { id: cartID } });
    }
  }, [cartID, error, newCart, fetchCart]);

  if (newCartError) {
    return { loading: false, error: newCartError };
  }

  if (cart) {
    const refresh = () => {
      refetch({ id: cart.cart.id });
    };
    return { loading: false, cart: cart.cart, refresh };
  }

  return { loading: true };
}

const Shop = () => {
  const { loading, error, cart, refresh } = useCart();
  const [addItemToCart] = useAddItemToCartMutation();

  if (loading) {
    return <>Loading...</>;
  }
  if (error) {
    return <>Error: {error.toString()}</>;
  }

  const handleAddToCart = (itemID: string) => {
    addItemToCart({ variables: { itemID, cartID: cart!.id, quantity: 1 } })
      .then(() => refresh!())
      .catch(console.error);
  };

  return (
    <>
      <Items onAddToCart={handleAddToCart} />
      <CartView cart={cart!} />
    </>
  );
};

export default Shop;
