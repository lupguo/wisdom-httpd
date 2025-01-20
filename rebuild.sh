#!/usr/bin/env bash

# 代理网络开启
proxy_command start

# 代码更新
git pull

# 编译安装
ls -al /data/go/bin/wisdom-httpd
go install -v .
ls -al /data/go/bin/wisdom-httpd

# 服务重启
systemctl restart wisdom-httpd.service

# 代理网络停止
proxy_command stop
