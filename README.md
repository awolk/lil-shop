# Lil-Shop

A small bare-bones online shop. I wrote this to practice developing a Stripe integration and to explore some technologies that were new to me.

## Technology
The project is split into distinct frontend and backend components. They communicate via a GraphQL api defined in the `common` folder.

### Backend
The backend is written in Go. I use the [Ent Framework](https://entgo.io/) as an ORM and [gqlgen](https://gqlgen.com/) to generate GraphQL scaffolding code. Both of these tools use code generation in order to provide type-safe APIs, as strong type-safety was one of my goals while developing this project. 

I was curious to see how well that could be achieved despite Go's simpler type system, and I found the tools a joy to work with. It seems that there is a greater ecosystem developing around code-generation for Go, and I much prefer this to solutions that depend on reflection, as you lose the benefits of strong static typing with reflection-based APIs, making refactoring more difficult, and requiring tests to be more comprehensive.

Use the command `cd backend; go generate ./...` to regenerate any auto-generated code.

I use the [Stripe Go library](https://github.com/stripe/stripe-go) to handle payments processing with Stripe.

### Frontend
The frontend is a [create-react-app](https://github.com/facebook/create-react-app) Typescript React webapp.  I use [graphql-codegen](https://graphql-code-generator.com/) to generate well-typed React hooks for communicating with the backend, and [React Stripe.js](https://github.com/stripe/react-stripe-js) to handle transfering payment details.

## Environment variables
Make a file `.env` in the root directory containing the following environment variables:
- `STRIPE_PUBLISHABLE_KEY`
- `STRIPE_SECRET_KEY`

Then it's just `docker-compose up` to start the application on localhost:3000.
