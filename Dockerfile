FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/hid8/test-native-function
COPY . .
RUN go get -d -v

RUN go build -o /build/ntf

FROM alpine

COPY --from=builder /build/ntf /app/ntf

EXPOSE 8080
ENTRYPOINT ["/app/ntf"]