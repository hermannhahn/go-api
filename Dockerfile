# syntax=docker/dockerfile:1
FROM golang:1.18.4
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy
COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=$(git tag) -X main.buildDate=$(date -u +'%d-%m-%Y|%H:%M:%S')" -o go-api .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
EXPOSE 8080
COPY --from=0 /app/ ./
CMD ["./go-api"]