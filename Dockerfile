FROM golang:alpine
WORKDIR /app

ADD go.mod .
ADD go.sum .
COPY module11/cmd/main.go .

RUN go build -o bin .
ENTRYPOINT ["/app/bin"]