FROM alpine:latest

USER root

# 添加 go 环境变量 和 alpine 镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && echo "export GO111MODULE=on" >> /etc/profile \
    && echo "export GOPATH=/root/go" >> /etc/profile \
    && echo "export GOPROXY=https://mirrors.aliyun.com/goproxy/" >> /etc/profile \
    && source /etc/profile

# 安装 micro
RUN apk update && apk add go git musl-dev xz binutils \
    && cd \
    && go get -u -v github.com/hequan2017/go-admin \
    && go install github.com/hequan2017/go-admin

# 压缩 和 加壳
RUN wget https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz \
    && xz -d upx-3.95-amd64_linux.tar.xz \
    && tar -xvf upx-3.95-amd64_linux.tar \
    && cd upx-3.95-amd64_linux \
    && chmod a+x ./upx \
    && mv ./upx /usr/local/bin/ \
    && cd /root/go/bin \
    && strip --strip-unneeded micro \
    && upx micro \
    && chmod a+x ./micro \
    && cp micro /usr/local/bin


COPY --from=0 /usr/local/bin/micro /usr/local/bin/

EXPOSE 80

CMD ["go-admin"]

