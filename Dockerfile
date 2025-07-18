ARG GO_VERSION=1.22.2

# for you

# FROM dawitalemu4/keycurl:latest

# COPY .env .

# CMD ["keycurl"]



# for me (push to docker hub)

# FROM golang:${GO_VERSION} AS builder

# COPY . .

# RUN go build -o /keycurl


# FROM archlinux:base

# COPY --from=builder /keycurl /usr/local/bin
# COPY --from=builder /go/views /views

# docker image build -t keycurl .
# docker image tag keycurl dawitalemu4/keycurl:latest
# docker push dawitalemu4/keycurl:latest



# for me (test image before publish)

# FROM keycurl:latest

# COPY .env .

# CMD ["keycurl"]



# for me (test locally)

FROM golang:1.22.2 AS builder

COPY . .

RUN go build -o /keycurl


FROM archlinux:base

COPY --from=builder /keycurl /usr/local/bin
COPY --from=builder /go/views /views

COPY .env .

CMD ["keycurl"]
