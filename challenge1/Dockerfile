FROM golang:1.9.2 as builder
ARG SOURCE_LOCATION=/
WORKDIR ${SOURCE_LOCATION}
COPY main.go .
COPY repos.txt .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk update && apk add ca-certificates && apk add yarn dumb-init && rm -rf /var/cache/apk/*
ARG SOURCE_LOCATION=/
WORKDIR /root/
COPY --from=builder ${SOURCE_LOCATION} .
ENTRYPOINT ["/usr/bin/dumb-init","--"]