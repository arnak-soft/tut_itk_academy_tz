FROM golang:1.20

WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./

# Download the modules needed for the build. This is done in a separate step so
# that if the go.mod file changes, the layer will be rebuilt, but if the go.mod
# file doesn't change, the layer will be reused, which can speed up the build
# process.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /wallet-app .cmd/main.go

# Command to run the executable
CMD [ "/wallet-app" ]