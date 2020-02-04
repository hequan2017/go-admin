FROM alpine:latest

USER root


RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && echo "export GO111MODULE=on" >> /etc/profile \
    && echo "export GOPATH=/root/go" >> /etc/profile \
    && echo "export GOPROXY=https://mirrors.aliyun.com/goproxy/" >> /etc/profile \
    && source /etc/profile

RUN apk update && apk add go git musl-dev xz binutils \
    && cd \
    && go get -u -v go-admin \
    && go install go-admin


RUN wget https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz \
    && xz -d upx-3.95-amd64_linux.tar.xz \
    && tar -xvf upx-3.95-amd64_linux.tar \
    && cd upx-3.95-amd64_linux \
    && chmod a+x ./upx \
    && mv ./upx /usr/local/bin/ \
    && cd /root/go/bin \
    && strip --strip-unneeded go-admin \
    && upx go-admin \
    && chmod a+x ./go-admin \
    && cp go-admin /usr/local/bin

FROM alpine:latest

COPY --from=0 /usr/local/bin/go-admin /usr/local/bin/

EXPOSE 80

CMD ["go-admin"]

