# 方便Rocky-Linux机器快速代码更新和部署
install-rocky-linux:
	# 1. 代码更新
	git pull

	# 2. 编译安装
	ls -al /data/go/bin/wisdom-httpd
	go install -v .
	ls -al /data/go/bin/wisdom-httpd

	# 3. 服务重启
	systemctl restart wisdom-httpd.service
