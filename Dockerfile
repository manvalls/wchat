# Titan build

FROM golang:1.11-alpine AS wchat
RUN apk update
RUN set -ex
RUN apk add --no-cache --virtual .build-deps gcc libc-dev git

RUN go get github.com/manvalls/quicktemplate
RUN go install github.com/manvalls/quicktemplate/qtc

ADD . /go/src/github.com/manvalls/wchat
WORKDIR /go/src/github.com/manvalls/wchat
RUN go get
RUN go generate
RUN go install --ldflags '-extldflags "-static"'

# Final image

FROM debian
RUN apt-get update && apt-get install -y \
    fuse musl ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=wchat /go/bin/wchat /usr/local/bin

CMD ["/usr/local/bin/wchat"]