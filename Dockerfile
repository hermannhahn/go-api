FROM golang:1.18.4
WORKDIR /app

#   ██████╗  ██████╗      █████╗ ██████╗ ██╗
#  ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║
#  ██║  ███╗██║   ██║    ███████║██████╔╝██║
#  ██║   ██║██║   ██║    ██╔══██║██╔═══╝ ██║
#  ╚██████╔╝╚██████╔╝    ██║  ██║██║     ██║
#   ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝ v0.1.1-alpha

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN echo "Installing dependencies..."
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

# build the app
COPY . ./
RUN export PORT=~/go/bin
RUN echo "Installing swagger..."
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN echo "Generating swagger files..."
RUN swag init --parseDependency --parseInternal
RUN echo "Building API..."
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /go/bin/ \
    -ldflags "-X main.version=$(git rev-parse HEAD) -X main.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

# deploy the app
CMD go-api