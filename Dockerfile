FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o package-handler

FROM scratch

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/package-handler .

EXPOSE 8080

CMD ["./package-handler"]