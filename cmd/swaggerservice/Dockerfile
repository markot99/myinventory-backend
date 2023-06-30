FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /swaggerservice cmd/swaggerservice/swaggerservice.go

ENV GIN_MODE=release
CMD ["/swaggerservice"]
