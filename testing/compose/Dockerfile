# Allow go version to be set at build
ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine as build

RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep

COPY cmd /go/src/github.com/kfirbreger/vault-unsealer/cmd/
COPY internal /go/src/github.com/kfirbreger/vault-unsealer/internal/
COPY Gopkg* /go/src/github.com/kfirbreger/vault-unsealer/

WORKDIR /go/src/github.com/kfirbreger/vault-unsealer

RUN dep ensure --vendor-only
RUN go build -o unsealer cmd/unsealer/main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/src/github.com/kfirbreger/vault-unsealer/unsealer .

# Copying keys and config
COPY keys.txt /app/keys.txt
COPY config.toml /app/config.toml

CMD ["./unsealer", "-key-file", "keys.txt"]  
 
