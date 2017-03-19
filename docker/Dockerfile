FROM golang:alpine

RUN apk --update upgrade && \
    apk add git curl ca-certificates mysql && \
    update-ca-certificates && \
    curl https://glide.sh/get | sh && \
    rm -rf /var/cache/apk/*

ADD gocover.sh /usr/local/bin
RUN chmod +x /usr/local/bin/gocover.sh

WORKDIR /go/src
ENTRYPOINT ["gocover.sh"]
