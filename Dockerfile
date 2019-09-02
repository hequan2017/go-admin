FROM alpine:latest

USER root

RUN apk update && apk add go git musl-dev

# 添加 go 环境变量
RUN echo "export PATH=\$PATH:/root/go/bin" >> /etc/profile \
    && echo "export GO111MODULE=on" >> /etc/profile \
    && echo "export GOPATH=/root/go" >> /etc/profile  \
    && echo "export GOPROXY=https://mirrors.aliyun.com/goproxy/" >> /etc/profile \
    && source /etc/profile

# 安装 micro
RUN cd \
    && go get -u -v github.com/hequan2017/go-admin \
    && go install github.com/hequan2017/go-admin

EXPOSE 80

ENTRYPOINT ["go","run","/go-admin/main.go"]
