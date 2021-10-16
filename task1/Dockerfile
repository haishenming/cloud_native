FROM golang:1.16 as builder

WORKDIR /task1
COPY . .
RUN export GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o service

FROM alpine:latest

WORKDIR /task1
COPY --from=builder /task1 .
EXPOSE 80

CMD ["./service"]