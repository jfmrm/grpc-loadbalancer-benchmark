FROM golang:1.19.1-bullseye as build

WORKDIR /app

COPY client/go.mod .
COPY client/go.sum .

RUN go mod download

COPY client/* .

RUN go build -o client-bin .

FROM golang:1.19.1-bullseye as run

WORKDIR /app

COPY --from=build /app/client-bin .

COPY helloworld .

CMD ["./client-bin"]
