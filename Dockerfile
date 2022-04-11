FROM golang:1.18.0-bullseye as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server-bin ./server/

FROM golang:1.18.0-bullseye as run

WORKDIR /app

COPY --from=build /app/server-bin .

CMD ["./server-bin"]
