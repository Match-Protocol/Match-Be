# 使用轻量级基础镜像
FROM alpine:3.18.0

ENV Private =be36199fb22c6dd529d451c6db01243fb0c07298e3a7f074391e2417e50961e8
# 设置工作目录
WORKDIR /app

# 复制可执行文件进容器
COPY match .
COPY config.toml .

# 暴露端口（如果你服务监听了某个端口，比如 8080）
EXPOSE 8099

# 设定容器启动时运行的命令
CMD ["./match"]
