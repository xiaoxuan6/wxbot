FROM golang:1.18-alpine AS build

WORKDIR /app

ADD . /app

# 配置goproxy
ENV GOPROXY=https://goproxy.io,direct \
    GO111MODULE=on

# 时区
RUN apk --no-cache add tzdata

RUN GOOS=linux GOARCH=386 go build -o wxbot main.go

FROM scratch

COPY --from=build /app/wxbot .

# 证书
#ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 时区
ENV TZ=Asia/Shanghai
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

CMD ["./wxbot"]

