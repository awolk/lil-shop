import React, { FunctionComponent } from "react";
import { Cart } from "./generated/graphql";

interface Props {
  cart: Cart;
}

const CartView: FunctionComponent<Props> = (props) => {
  return (
    <>
      Cart:
      <ul>
        {props.cart.lineItems.map((lineItem) => (
          <li key={lineItem.id}>
            {lineItem.quantity} {lineItem.item.name}
          </li>
        ))}
      </ul>
    </>
  );
};

export default CartView;
