FROM golang:latest

LABEL maintainer="Pham Ngoc Phu <phumaster.dev@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE ${APP_PORT}
CMD ["./main"]