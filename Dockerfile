FROM golang:1.12.1-alpine3.9

RUN addgroup -g 1000 -S deployapp && \
    adduser -u 1000 -S deployapp -G deployapp

USER root
RUN mkdir webserver-go
COPY . webserver-go
WORKDIR webserver-go
RUN chown deployapp:deployapp -R /go/webserver-go

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh


USER deployapp
RUN export GO111MODULE=on
RUN go build .

EXPOSE 1323
CMD ["./webserver-go-templates"]