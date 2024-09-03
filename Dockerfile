# 使用官方 Golang 镜像作为构建阶段
FROM golang:1.22-alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 创建并设置工作目录
WORKDIR /app

# 将项目的 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载所有依赖项
RUN go mod download

# 将项目的源代码复制到工作目录
COPY . .

# 编译 GoFrame 项目
RUN go build -o main .

# 使用最小化的镜像作为运行环境
FROM alpine:latest

# 设置时区（可选）
RUN apk --no-cache add tzdata
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译后的二进制文件和配置文件
COPY --from=builder /app /app

# 暴露服务端口
EXPOSE 8808

# 设置启动命令
CMD ["./main"]
