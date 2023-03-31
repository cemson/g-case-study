# Dockerfile for dockerizing purpose
FROM golang:1.18

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build .

EXPOSE 11923

CMD ["/app/g-case-study", "prod"]