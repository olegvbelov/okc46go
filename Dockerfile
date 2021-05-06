FROM golang:1.15.11-alpine3.13
RUN mkdir /app
RUN mkdir /app/go
ADD . /app/go
WORKDIR /app/go
RUN go mod download
CMD ["/app/go/run.sh"]
