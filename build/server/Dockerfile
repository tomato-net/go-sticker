ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine as test

WORKDIR /workspace
RUN apk add --no-cache git

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go test -timeout=60s -coverprofile=cover.out ./...

FROM test as build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server cmd/server/main.go

FROM gcr.io/distroless/static:latest as release
WORKDIR /

USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /workspace/server .
ENTRYPOINT ["/server"]