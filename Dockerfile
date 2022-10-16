FROM golang:latest As buildStage
WORKDIR /go/src
COPY . /go/src
RUN cd /go/src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=buildStage /go/src/main /app/
CMD ["./main"]
