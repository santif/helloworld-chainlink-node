# Build stage
FROM golang:1.17-alpine AS build-env

ADD go.mod     /work/go.mod
ADD go.sum     /work/go.sum
ADD main.go    /work/main.go
ADD api.go     /work/api.go
ADD model.go   /work/model.go

RUN cd /work && go build -mod=readonly -o /work/service *.go

# Final stage
FROM alpine:3.13

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /work/service .
ENTRYPOINT [ "./service" ]
