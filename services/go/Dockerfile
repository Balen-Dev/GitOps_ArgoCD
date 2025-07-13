FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY api.go .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -a -o bin

FROM scratch
USER 1001
COPY --from=builder /app/bin /bin/api
EXPOSE 8080
ENTRYPOINT ["/bin/api"]
