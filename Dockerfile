# 表示依赖 alpine 最新版
FROM ubuntu:20.04
ENV VERSION 1.0

RUN apt update && apt install libssl-dev ca-certificates -y && update-ca-certificates 

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 拷贝当前目录下 go_docker_demo1 可以执行文件
COPY tx_relayer /apps/tx_relayer

RUN chmod +x /apps/tx_relayer

# 设置编码
ENV LANG C.UTF-8

# 运行golang程序的命令
CMD ["./tx_relayer"]