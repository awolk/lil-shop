FROM golang:1.14.4

WORKDIR /app

# install reflex
RUN go get github.com/cespare/reflex && \
  go install github.com/cespare/reflex

# install stripe-cli
WORKDIR /usr/bin
RUN wget https://github.com/stripe/stripe-cli/releases/download/v1.4.4/stripe_1.4.4_linux_x86_64.tar.gz
RUN tar -xvf stripe_1.4.4_linux_x86_64.tar.gz
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ./entry.sh