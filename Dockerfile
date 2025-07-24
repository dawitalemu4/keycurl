ARG GO_VERSION=1.22.2

# for you

# FROM dawitalemu4/keycurl:latest

# COPY .env .

# CMD ["keycurl"]



# for me (push to docker hub)

# FROM golang:${GO_VERSION} AS builder

# COPY . .

# RUN go build -o /keycurl


# FROM debian:stable-20250721-slim

# RUN apt-get update && \
#     apt-get install -y curl && \
#     rm -rf /var/lib/apt/lists/*

# COPY --from=builder /keycurl /usr/local/bin
# COPY --from=builder /go/views /views

# docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t dawitalemu4/keycurl:latest --push .



# for me (test locally)

FROM golang:1.22.2 AS builder

COPY . .

RUN go build -o /keycurl


FROM debian:stable-20250721-slim

RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /keycurl /usr/local/bin
COPY --from=builder /go/views /views

COPY .env .

CMD ["keycurl"]
