#!/usr/bin/env bash

# 代码更新
git pull

# 编译安装
ls -al /data/go/bin/wisdom-httpd
go install -v .
ls -al /data/go/bin/wisdom-httpd

# 服务重启
systemctl restart wisdom-httpd.service
sleep 3
systemctl status wisdom-httpd.service
