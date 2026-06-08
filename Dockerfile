# ── Build stage ─────────────────────────────────────────
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o travelsphere .

# ── Run stage ──────────────────────────────────────────
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary
COPY --from=builder /app/travelsphere .

# Copy config, templates, and static assets
COPY --from=builder /app/conf ./conf
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./travelsphere"]
