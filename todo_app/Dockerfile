# stage 1: Build the Go binary
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

# copy dependency files first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the application source code
COPY . .

# build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-app ./cmd/api/main.go

# stage 2: Create a minimal production image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy the compiled binary from the builder stage
COPY --from=builder /app/todo-app .

# ensure .env exists so godotenv.Load() does not fail in container
RUN touch .env

EXPOSE ${HTTP_PORT}

# run the application
CMD ["./todo-app"]
