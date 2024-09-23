# container for building the app
FROM golang:1.23.1-bookworm AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -buildvcs=false -ldflags "-w -s" -o app

# ----------------------------

# container for running the app in production
FROM debian:bookworm-slim AS deploy

RUN apt update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ----------------------------

# container for running the app in development
FROM golang:1.23.1 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]
