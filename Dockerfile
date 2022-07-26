# install backend dependencies and build
FROM golang:1.18.4
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

# build the app
COPY . ./
RUN export PORT=~/go/bin
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /go/bin/ \
    -ldflags "-X main.version=$(git rev-parse HEAD) -X main.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

# deploy the app
CMD go-api