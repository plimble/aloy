FROM golang:alpine

RUN apk --update upgrade && \
    apk add git curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

COPY app .
CMD ["./app"]
