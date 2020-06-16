import React, { FunctionComponent, useEffect, useState } from "react";
import { loadStripe } from "@stripe/stripe-js";
import {
  Elements,
  useStripe,
  useElements,
  CardElement,
} from "@stripe/react-stripe-js";
import { useCheckoutCartMutation } from "./generated/graphql";
import { formatCost } from "./util";

const stripePromise = loadStripe(
  process.env.REACT_APP_STRIPE_PUBLISHABLE_KEY || ""
);

interface Props {
  cartID: string;
}

const InternalForm: FunctionComponent<Props> = (props) => {
  const stripe = useStripe();
  const elements = useElements();
  const [error, setError] = useState<any>(null);
  const [clientSecret, setClientSecret] = useState("");
  const [totalCost, setTotalCost] = useState<number | null>();
  const [processing, setProcessing] = useState(false);
  const [done, setDone] = useState(false);
  const [fetchPaymentIntent] = useCheckoutCartMutation();

  useEffect(() => {
    fetchPaymentIntent({ variables: { cartID: props.cartID } })
      .then(({ data }) => {
        setClientSecret(data!.checkoutCart.clientSecret);
        setTotalCost(data!.checkoutCart.totalCostCents);
      })
      .catch(setError);
  }, [fetchPaymentIntent, props.cartID]);

  const submit = async () => {
    setProcessing(true);
    const payload = await stripe!.confirmCardPayment(clientSecret, {
      payment_method: {
        card: elements!.getElement(CardElement)!,
      },
    });

    setProcessing(false);
    if (payload.error) {
      setError(payload.error);
    } else {
      setDone(true);
    }
  };

  if (error) {
    return <>Error: ${error.toString()}</>;
  }

  if (done) {
    return <>Payment Successful</>;
  }

  return (
    <>
      Total: {totalCost && formatCost(totalCost)}
      <br />
      <div style={{ maxWidth: "450px" }}>
        <CardElement />
      </div>
      {processing ? (
        "Payment Processing"
      ) : (
        <button onClick={submit}>Pay Now</button>
      )}
    </>
  );
};

const CheckoutForm: FunctionComponent<Props> = (props) => {
  return (
    <Elements stripe={stripePromise}>
      <InternalForm cartID={props.cartID} />
    </Elements>
  );
};

export default CheckoutForm;
