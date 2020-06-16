const costFormatter = new Intl.NumberFormat(undefined, {
  style: "currency",
  currency: "USD",
});

export function formatCost(costCents: number) {
  return costFormatter.format(costCents / 100);
}
