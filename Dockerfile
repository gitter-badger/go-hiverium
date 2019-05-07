FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /mainchain
RUN cd /mainchain && make und

FROM alpine:latest

LABEL maintainer="etienne@tomochain.com"

WORKDIR /mainchain

COPY --from=builder /mainchain/build/bin/tomo /usr/local/bin/und

RUN chmod +x /usr/local/bin/und

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/und"]

CMD ["--help"]
