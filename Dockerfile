FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN go build -o receipt_processor

EXPOSE 8080
CMD ["./receipt_processor"]