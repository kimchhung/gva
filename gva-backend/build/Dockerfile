FROM golang:1.22.3-alpine3.18 AS builder

WORKDIR /app

COPY ./go.* ./
ENV GO111MODULE=on
RUN go mod download

# Copy the rest of your application
COPY lottery-backend/. .

# Build your application to binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api-binary ./api/main.go

# Compress the binary
RUN apk add --no-cache upx
RUN upx -9 /bin/api-binary


# Copy necessary config file
FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/internal/geo2 /app/internal/geo2
COPY --from=builder --chown=nonroot /app/lang /app/lang
COPY --from=builder --chown=nonroot /app/env /app/env
ENTRYPOINT ["/bin/api-binary"]
EXPOSE 4000
