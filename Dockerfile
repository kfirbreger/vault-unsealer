# Allow go version to be set at build
ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine as build

RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/kfirbreger/vault-unsealer

WORKDIR /go/src/github.com/kfirbreger/vault-unsealer
RUN dep ensure --vendor-only
RUN go build cmd/unsealer/main.go
 
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/src/app/cmd/unsealer/unsealer .
CMD ["./unsealer"]  
 
