FROM golang:1.18-alpine AS build

WORKDIR /app

ADD . /app

# 配置goproxy
ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on

RUN GOOS=linux GOARCH=386 go build -o wxbot main.go

FROM scratch

# 构建镜像报错：X509: Certificate Signed by Unknown Authority
# 1、安装根证书，镜像 scratch(虚拟镜像) 不支持 apt-get 命令（无法解决）
# https://www.jianshu.com/p/97471c082b2f
#RUN apt-get update && apt-get install -y ca-certificates

# 2、使用 ADD 命令，将 cacert.pem 添加到镜像中
# wget https://curl.haxx.se/ca/cacert.pem
ADD cacert.pem /etc/ssl/certs/

# 3、不使用 https

COPY --from=build /app/wxbot .

CMD ["./wxbot"]

