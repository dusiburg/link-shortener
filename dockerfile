FROM golang:1.25-alpine3.21 AS builder

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o bin/application ./cmd

FROM alpine:3.21 AS runner

WORKDIR /opt

COPY --from=builder /opt/bin/application ./

CMD ["./application"]