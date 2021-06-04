FROM golang:1.15-alpine3.12 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.io/
# GOPROXY=https://mirrors.aliyun.com/goproxy/
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN go build -o app main.go
WORKDIR /dist
RUN cp /build/app ./app

FROM scratch
COPY --from=builder /dist/app /app
WORKDIR /
ENTRYPOINT ["/app"]