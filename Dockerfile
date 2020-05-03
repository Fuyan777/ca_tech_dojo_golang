FROM golang:1.13.5-alpine
WORKDIR /go/src/game
EXPOSE 8080
RUN apk add --no-cache git && \
    go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql