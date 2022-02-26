FROM golang:1.16-alpine

WORKDIR /app

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY=https://proxy.golang.org,direct

RUN go get github.com/beego/bee/v2

COPY . .

RUN ["chmod","+x","start.sh"]

CMD ["./start.sh"]

