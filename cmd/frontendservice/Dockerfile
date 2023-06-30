FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /frontendservice cmd/frontendservice/frontendservice.go

COPY cmd/frontendservice/frontend/ ./frontend/

ENV GIN_MODE=release
CMD ["/frontendservice"]