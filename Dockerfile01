FROM golang:alpine
WORKDIR /app

ADD go.mod /app
ADD go.sum /app
COPY module11/cmd/main.go /app

RUN go build -o /app/bin .

EXPOSE 8080
ENTRYPOINT ["/app/bin"]