FROM golang:alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get firebase.google.com/go
RUN go get github.com/gin-gonic/gin
RUN go get google.golang.org/api/option
RUN go get github.com/dgrijalva/jwt-go

ENV SOURCES /go/src/firebase-custom-auth/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}firebase-custom-auth
EXPOSE 8081:8081