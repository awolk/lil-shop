FROM golang:1.14.4

WORKDIR /app

RUN go get github.com/cespare/reflex && \
  go install github.com/cespare/reflex

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD reflex -r "(\.go$|go\.mod)" -s go run .