FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o my-app .

FROM alpine:3.22
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/my-app ./my-app
USER 65532:65532
ENTRYPOINT ["./my-app"]
