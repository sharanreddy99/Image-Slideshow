FROM golang:1.16-alpine

WORKDIR /app

RUN apk update && apk add netcat-openbsd

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY=https://proxy.golang.org,direct

RUN go get github.com/beego/bee/v2

COPY ./start.sh ../
RUN ["chmod","+x","../start.sh"]

COPY . ./

EXPOSE 8080

CMD ["../start.sh"]

