# Stage 1: Build the Go binary
FROM golang:1.22.4-alpine

WORKDIR /post-api

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /post-api/app ./cmd/main.go

#run app
ENTRYPOINT [ "/post-api/app" ]