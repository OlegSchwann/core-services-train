FROM golang:1.12-alpine3.10

LABEL maintainer='OlegSchwann <cup.of.software.code@gmail.com>'

RUN apk add --no-cache git \
    && go get \
        'github.com/mailru/easyjson' \
        'golang.org/x/net/html' \
        'github.com/PuerkitoBio/goquery' \
        'github.com/pkg/errors' \
    && apk del git ;

COPY . "${GOPATH}/src/github.com/OlegSchwann/core-services-train/2_categories"

RUN go build -o '/usr/local/go/bin/categories' 'github.com/OlegSchwann/core-services-train/2_categories' ;

CMD ["categories"]
