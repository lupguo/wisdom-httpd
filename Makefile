# 定义目标为编译可执行文件
build:
	go build -v -o wisdom-httpd .

install:
	go install -v .

# 定义目标为清理生成的可执行文件
clean:
	rm -f wisdom-httpd

# 定义目标为交叉编译 Linux 和 Windows 版本
cross-compile:
	GOOS=linux GOARCH=amd64 go build -o wisdom-httpd-linux-amd64 .
