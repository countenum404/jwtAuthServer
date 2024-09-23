FROM golang:1.22.6-alpine3.20

WORKDIR /go/app
COPY ./ ./
RUN go build -o application cmd/main.go

EXPOSE 8080

CMD ["./application"]