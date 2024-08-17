FROM golang:1.22.3-alpine3.18 AS builder

WORKDIR /app

COPY ./go.* ./
ENV GO111MODULE=on
RUN go mod download

# Copy the rest of your application
COPY . .

# Build your application to binary
RUN go build -o /bin/api-binary ./cmd/api/main.go


# Copy necessary config file
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/lang /app/lang
COPY --from=builder --chown=nonroot /app/env /app/env
ENTRYPOINT ["/bin/api-binary"]
EXPOSE 4000
