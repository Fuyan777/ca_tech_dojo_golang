FROM golang:1.13.5-alpine
WORKDIR /go/src/game
EXPOSE 8080
RUN apk add --no-cache git && \
    go get github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/joho/godotenv && \
    go get github.com/dgrijalva/jwt-go && \
    go get github.com/go-sql-driver/mysql