FROM golang:1.12.1

RUN mkdir /opt/webserver-go
COPY . /opt/webserver-go
WORKDIR /opt/webserver-go

RUN export GO111MODULE=on
RUN go build .

CMD ["./webserver-go-templates"]