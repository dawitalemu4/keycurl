FROM dawitalemu4/gokey:latest AS builder


# change the next line to FROM --platform=linux/amd64 golang:1.22.2 if you are a mac user and getting this error: "rosetta error: failed to open elf at /lib64/ld-linux-x86-64.so.2"
FROM golang:1.22.2

RUN apt-get update && apt-get install -y curl

COPY --from=builder /gokey-cURL /gokey-cURL
COPY --from=builder /go/views /go/views

COPY .env .

CMD ["/gokey-cURL"]
