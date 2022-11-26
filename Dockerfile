FROM golang:1.18-alpine AS build

WORKDIR /app

ADD . /app

# 配置goproxy
ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on

RUN GOOS=linux GOARCH=386 go build -o wxbot main.go

FROM scratch

COPY --from=build /app/wxbot .

ENTRYPOINT ["./wxbot"]

