FROM golang:1.12
 
WORKDIR /go/src/app
COPY . .
 
RUN dep ensure
RUN GOOS=linux go build cmd/unsealer/main.go
 
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/src/app/cmd/unsealer/unsealer .
CMD ["./unsealer"]  
 