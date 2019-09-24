#!/usr/bin/env bash

USERNAME=${USERNAME:-admin}
EMAIL=${EMAIL:-auxpi@example.com}
PASSWORD=${PASSWORD:-12345678}

# 检查是否已经初始化
[ -f "/opt/go/conf/app.conf" ] || cp -a /opt/go/confbak/app.conf /opt/go/conf/app.conf
[ -f "/opt/go/conf/siteConfig.json" ] || (
    /opt/go/auxpi migrate
    /opt/go/auxpi -mod=admin -name=${USERNAME} -email=${EMAIL} -pass=${PASSWORD}
)

if [ "$1" = "bash" -o "$1" = "sh" -o "$1" = "ash" ]; then
    exec /bin/bash
else
    exec /opt/go/auxpi "$@"
fi