FROM hub.c.163.com/public/ubuntu:16.04
MAINTAINER wchuang

RUN apt-get update
RUN apt-get install -y git openssh-client vim tar wget curl supervisor \
    && rm -rf /var/lib/apt/lists/*

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p /var/www/html
WORKDIR /var/www/html

ENTRYPOINT /usr/bin/supervisord