FROM golang:1.12-alpine3.10

LABEL maintainer='OlegSchwann <cup.of.software.code@gmail.com>'

RUN apk add --no-cache git \
    && go get 'github.com/mailru/easyjson' \
    && apk del git ;

COPY . "${GOPATH}/src/github.com/OlegSchwann/core-services-train/1_locations_tree"

RUN go build -o '/usr/local/go/bin/locations_tree' 'github.com/OlegSchwann/core-services-train/1_locations_tree' ;

CMD ["locations_tree"]
