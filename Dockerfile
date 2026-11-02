# Build stage
FROM golang:1.25.6-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /kindle-go ./cmd/main

# Runtime stage
FROM alpine:3.20
COPY --from=build /kindle-go /kindle-go
EXPOSE 9010
CMD ["/kindle-go"]
