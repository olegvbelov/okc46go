FROM golang:1.15.11-alpine3.13
RUN mkdir /app
RUN mkdir /app/go
ADD . /app/go
WORKDIR /app/go
RUN go mod download
RUN go build -o okc46 cmd/web/*.go
CMD ["./okc46"]
