#指定构建镜像的基础镜像
FROM golang:1.21.8-alpine3.19

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#设置工作目录，后面的RUN,COPY等都基于在这个目录下工作
WORKDIR /goApp

# 复制项目文件到工作目录
COPY . .

# 编译应用程序成二进制文件app
RUN go build -o app

#容器暴露的端口
EXPOSE 8080

#启动容器时运行的命令
CMD ["./app"]