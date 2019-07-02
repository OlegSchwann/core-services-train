FROM golang:1.12-alpine3.10

LABEL maintainer='OlegSchwann <cup.of.software.code@gmail.com>'

RUN apk add --no-cache git \
    && go get \
        'github.com/mailru/easyjson' \
        'github.com/pkg/errors' \
    && apk del git ;

COPY . "/go/src/github.com/OlegSchwann/core-services-train/3_categories_moscow"

RUN go build -o '/usr/local/go/bin/categories_location' 'github.com/OlegSchwann/core-services-train/3_categories_moscow' ;

CMD categories_location -location='moskva' -categories_path='/go/src/github.com/OlegSchwann/core-services-train/3_categories_moscow/categories.json';
