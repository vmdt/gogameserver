FROM golang:1.24.4-alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app

# Install dependencies
RUN apk --no-cache add ca-certificates tzdata

COPY . .

RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go mod download

RUN go build -o build cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata 

COPY --from=builder /app/build /release

ENTRYPOINT [ "/release" ]