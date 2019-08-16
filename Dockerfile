#源镜像
FROM golang:latest

MAINTAINER hequan ""

RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*

WORKDIR /go/src

RUN  go get github.com/hequan2017/go-admin

#go构建可执行文件
RUN go build .
RUN go run main.go
#暴露端口
EXPOSE 80
#最终运行docker的命令
ENTRYPOINT  ["./mygohttp"]
