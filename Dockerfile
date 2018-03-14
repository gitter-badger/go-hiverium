# Build Geth in a stock Go builder container
FROM golang:1.10 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /happyuc-go
RUN cd /happyuc-go && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /happyuc-go/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp 30304/udp
ENTRYPOINT ["geth"]
