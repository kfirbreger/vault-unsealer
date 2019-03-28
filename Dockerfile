FROM golang:1.12 as build
 
WORKDIR /go/src/app
COPY . .
 
RUN dep ensure
RUN GOOS=linux go build cmd/unsealer/main.go
 
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/src/app/cmd/unsealer/unsealer .
CMD ["./unsealer"]  
 