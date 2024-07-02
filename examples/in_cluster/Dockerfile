FROM busybox
LABEL authors="james"

# 创建并设置工作目录
WORKDIR /app
# 复制当前目录的所有文件到工作目录
COPY . .
# 给二级制文件添加执行权限
RUN chmod +x /app/in_cluster
ENTRYPOINT ["./in_cluster"]