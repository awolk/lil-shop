import React, { FunctionComponent } from "react";
import { useAllItemsQuery } from "./generated/graphql";

interface Props {
  onAddToCart(itemID: string): void;
}

const Items: FunctionComponent<Props> = (props) => {
  const { data, loading, error } = useAllItemsQuery();
  if (loading) {
    return <>Loading</>;
  }

  if (error) {
    return <>Error: {error.toString()}</>;
  }

  return (
    <>
      Store:
      <ul>
        {data?.items.map((item) => (
          <li key={item.id}>
            {item.name}: {item.costCents / 100}
            &nbsp;
            <button onClick={() => props.onAddToCart(item.id)}>
              Add to Cart
            </button>
          </li>
        ))}
      </ul>
    </>
  );
};

export default Items;
