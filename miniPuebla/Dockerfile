# syntax=docker/dockerfile:1


#FROM golang:1.21.0
#FROM arm32v6/golang:1.21-alpine
#FROM arm32v6/golang:1.20-alpine
FROM golang:1.21.0-bullseye

WORKDIR /miniPuebla

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /miniPuebla

# Run

CMD ["./miniPuebla"]