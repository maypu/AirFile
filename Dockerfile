# 使用官方的Golang镜像作为基础镜像
FROM dockerproxy.net/library/golang:1.19-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将Go模块文件和Go源文件复制到工作目录
COPY go.mod go.sum ./
RUN go mod download

# 复制所有Go源文件到工作目录
COPY . .

# 编译Go程序，并将输出文件命名为airfile
RUN go build -o airfile .

# 使用轻量级的Alpine镜像作为运行环境
FROM dockerproxy.net/library/alpine:latest

# 设置工作目录
WORKDIR /app

# 从builder阶段复制编译好的可执行文件
COPY --from=builder /app/airfile .
# 为airfile添加执行权限
RUN chmod +x /app/airfile

# 复制配置文件到工作目录
COPY ./configs.yaml .
# 复制web文件夹到工作目录
COPY ./web ./web
# 复制数据库文件到工作目录下的/database目录
COPY ./database/data.db ./database/
# 为database目录及其内容添加可写入权限
RUN chmod -R a+w /app/database

# 创建files文件夹存储文件
RUN mkdir -p /app/files
RUN chmod -R a+w /app/files

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 暴露端口
EXPOSE 8086

# 运行程序
CMD ["./airfile"]