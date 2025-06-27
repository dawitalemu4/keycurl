# for you

# FROM dawitalemu4/keycurl:latest AS builder


# FROM golang:1.22.2

# RUN apt-get update && apt-get install -y curl

# COPY --from=builder /keycurl/keycurl
# COPY --from=builder /go/views /go/views

# COPY .env .

# CMD ["/keycurl"]


# for me (push to docker hub)

# FROM golang:1.22.2 AS builder

# COPY . .

# RUN go build -o /keycurl

# docker image build -t keycurl .
# docker image tag keycurl dawitalemu4/keycurl:latest
# docker push dawitalemu4/keycurl:latest


# for me (test locally)

FROM golang:1.22.2 AS builder

COPY . .

RUN go build -o /keycurl


FROM golang:1.22.2

RUN apt-get update && apt-get install -y curl

COPY --from=builder /keycurl /keycurl
COPY --from=builder /go/views /go/views

COPY .env .

CMD ["/keycurl"]


# for me (test published image)

# FROM dawitalemu4/keycurl:latest AS builder


# # change the next line to FROM --platform=linux/amd64 golang:1.22.2 if you are a mac user and getting this error: "rosetta error: failed to open elf at /lib64/ld-linux-x86-64.so.2"
# FROM golang:1.22.2

# RUN apt-get update && apt-get install -y curl

# COPY --from=builder /keycurl /keycurl
# COPY --from=builder /go/views /go/views

# COPY .env .

# CMD ["/keycurl"]
