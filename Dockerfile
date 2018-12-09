FROM golang:alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get firebase.google.com/go
RUN go get github.com/gin-gonic/gin
RUN go get google.golang.org/api/option

ENV SOURCES /go/src/firebase-authenticator/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}firebase-authenticator
EXPOSE 8081:8081
