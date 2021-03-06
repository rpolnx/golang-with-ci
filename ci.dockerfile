FROM golang:1.17.3-buster as installer

WORKDIR /usr/app

RUN apt update && apt install -y build-essential

COPY go.* ./

RUN export GOOS=linux
RUN export GOARCH=amd64
RUN export CGO_ENABLED=1
RUN export GOMODCACHE="/go/pkg/mod"

RUN go mod download -x
RUN go install github.com/swaggo/swag/cmd/swag@latest


FROM golang:1.17.3-buster as builder

WORKDIR /usr/app

RUN apt update && apt install -y build-essential

COPY go.* ./

RUN export GOOS=linux
RUN export GOARCH=amd64
RUN export CGO_ENABLED=1
RUN export GOMODCACHE="/go/pkg/mod"

COPY --chown=1001:1001 --from=installer /go/pkg/mod /go/pkg/mod
COPY --chown=1001:1001 --from=installer /go/bin/swag /go/bin/swag

COPY *.go ./
COPY src ./src
COPY test ./test

RUN go test -v ./test/...

RUN /go/bin/swag init -g main.go
RUN go build main.go


FROM debian:buster-slim as final

ENV APP_ROOT=/usr/app

ARG USER=golang
ARG GROUP=golang

WORKDIR $APP_ROOT

RUN apt update && apt install build-essential ca-certificates -y

RUN groupadd -g 1001 $GROUP && \
    adduser --system --no-create-home --uid 1001 --ingroup $GROUP --disabled-password $USER && \
    chown 1001:1001 $APP_ROOT

USER $USER

COPY --chown=1001:1001 --from=builder /usr/app/main .
COPY --chown=1001:1001 --from=builder /usr/app/docs ./

EXPOSE 8080

CMD ["./main"]