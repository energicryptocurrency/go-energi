# Build energi in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /energi
ENV GOFLAGS='-mod=mod -gcflags -dwarf=0 -ldflags "-s -w"'
RUN cd /energi && make geth

# Pull energi into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /energi/build/bin/energi3 /usr/local/bin/

EXPOSE 39796 39795 39797 39797/udp
ENTRYPOINT ["energi3"]
