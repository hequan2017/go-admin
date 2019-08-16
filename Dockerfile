FROM centos:latest

MAINTAINER hequan ""


RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup

RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
RUN  curl -o  /etc/yum.repos.d/epel-7.repo http://mirrors.aliyun.com/repo/epel-7.repo
RUN yum clean all
RUN yum makecache


RUN mkdir -p /data/go
RUN yum install golang -y
RUN echo export GO111MODULE=on >> /etc/profile
RUN echo export GOPATH=/data/go >> /etc/profile
RUN echo export GOPROXY=https://goproxy.io >> /etc/profile

RUN source /etc/profile

WORKDIR /data

RUN  git clone https://github.com/hequan2017/go-admin

WORKDIR /data/go-admin

RUN go build .

EXPOSE 80

ENTRYPOINT ["go","run","/data/go-admin/main.go"]
