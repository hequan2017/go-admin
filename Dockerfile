#源镜像
FROM golang:latest

MAINTAINER hequan ""

WORKDIR /go/src

RUN  go get github.com/hequan2017/go-admin
#go构建可执行文件
RUN go build .

#暴露端口
EXPOSE 80

ENTRYPOINT ["go","run","main.go"]
