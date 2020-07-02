FROM golang:latest
MAINTAINER  "Vikram Ingawale" <vikram.ingawale91@gmail.com>
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]
