FROM golang:1.11.0-alpine3.8

WORKDIR /go/src/app
COPY . .

RUN apk add git && go get -d -v ./...
RUN go install -v ./...

CMD ["app"]