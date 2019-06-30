MAINTAINER makemoney
FROM node

RUN git clone git@github.com:make-money-sysu/Make-Money.git \
    && cd Make-Money && npm install \
    && npm run build

FROM golang

RUN go get -u "github.com/astaxie/beego" \
    && go get -u "github.com/bitly/go-simplejson" \
    && go get -u "github.com/go-sql-driver/mysql" \
    && go get -u "github.com/make-money-sysu/server"


WORKDIR /go/src/github.com/make-money-sysu/server

CMD ["go", "run", "main.go"]