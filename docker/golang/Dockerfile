FROM golang:1.8

RUN mkdir -p /go/src/pulse
WORKDIR /go/src/pulse

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# this will ideally be built by the ONBUILD below ;)
ENTRYPOINT ping 127.0.0.1