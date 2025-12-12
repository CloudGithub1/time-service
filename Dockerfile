# Stage 1: build
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git

WORKDIR /src

COPY app/go.mod ./app/go.mod

RUN cd app && go mod download || true

# Copy source from above
COPY app/ ./app/

WORKDIR /src/app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o time-service

# Stage 2: final minimal image
FROM scratch
COPY --from=builder /src/app/time-service /time-service

USER 10001
EXPOSE 8080
ENTRYPOINT ["/time-service"]

