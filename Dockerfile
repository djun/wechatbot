FROM golang:1.20.12-alpine AS builder

MAINTAINER hongyun

WORKDIR /work

COPY . /work

ENV GOPROXY=https://goproxy.cn,direct GO111MODULE=on

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

FROM alpine AS runner

WORKDIR /work

COPY --from=builder /work/main .

# 按需添加代理
# ENV http_proxy=http://192.168.0.101:7890 https_proxy=http://192.168.0.101:7890 all_proxy=socks5://192.168.0.101:7890

ENTRYPOINT ["./main"]