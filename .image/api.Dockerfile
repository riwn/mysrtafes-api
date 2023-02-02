FROM golang:1.20 as Builder

WORKDIR /go/src

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src ./

WORKDIR /go/src/cmd/api

ARG GOOS=linux
ARG GOASRCH=amd64
RUN go build -o /go/bin/main -ldflags '-w -s'

FROM alpine:3

COPY --from=Builder /go/bin/main /api/main

EXPOSE 80

ENTRYPOINT ["/api/main"]
