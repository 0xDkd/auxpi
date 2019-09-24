FROM golang:1.13-stretch AS builder

ENV TZ Asia/Shanghai

ENV GO111MODULE on

ENV GOPROXY https://goproxy.cn

WORKDIR $GOPATH/src/github.com/auxpi

COPY . $GOPATH/src/github.com/auxpi

RUN make build

FROM debian:stretch-backports

ENV TZ Asia/Shanghai

COPY --from=builder /go/src/github.com/auxpi/auxpi /opt/go/auxpi

COPY static /opt/go/static

COPY views /opt/go/views

#COPY db /opt/go/db
RUN mkdir -p /opt/go/db /opt/go/conf

# 持久化数据库
VOLUME /opt/go/db
# 持久化配置
VOLUME /opt/go/conf

COPY conf /opt/go/confbak

COPY hack/run/entrypoint.sh /entrypoint.sh

RUN chmod 777 /opt/go/auxpi \
    && chmod 777 /entrypoint.sh

WORKDIR /opt/go

EXPOSE 2333

ENTRYPOINT ["/entrypoint.sh"]

CMD ["run"]
