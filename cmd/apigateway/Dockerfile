FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /apigateway cmd/apigateway/apigateway.go

ENV GIN_MODE=release
CMD ["/apigateway"]