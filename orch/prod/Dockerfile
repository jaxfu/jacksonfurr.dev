# Stage 1
FROM golang:latest AS builder
WORKDIR /app
COPY repo/server/ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server.exe server.go
RUN mkdir logs
RUN touch logs/msgs.txt

# Stage 2
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/server.exe /
COPY --from=builder /app/logs /logs
COPY repo/client /client
ENTRYPOINT ["/server.exe"]
