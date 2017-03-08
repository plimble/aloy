FROM golang:alpine

RUN apk --update upgrade && \
    apk add ca-certificates mysql && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

COPY app /
CMD ["/app"]
