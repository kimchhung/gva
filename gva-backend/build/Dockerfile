FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /app

COPY ./go.* ./
ENV GO111MODULE=on
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# Copy the rest of your application
COPY . .

# Build your application to binary
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/api-binary ./api/main.go


# Copy necessary config file
FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/lang /app/lang
COPY --from=builder --chown=nonroot /app/env /app/env
ENTRYPOINT ["/bin/api-binary"]
EXPOSE 5500
