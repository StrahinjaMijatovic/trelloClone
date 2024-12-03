FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /projects-service

EXPOSE 8080

CMD ["/projects-service"]
