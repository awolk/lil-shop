import React from "react";
import { useAllItemsQuery } from "./generated/graphql";

function Items() {
  const { data, loading, error } = useAllItemsQuery();
  if (loading) {
    return <>Loading</>;
  }

  if (error) {
    return <>Error: {error.toString()}</>;
  }

  return (
    <ul>
      {data?.items.map((item) => (
        <li key={item.id}>
          {item.name}: {item.costCents}
        </li>
      ))}
    </ul>
  );
}

export default Items;
