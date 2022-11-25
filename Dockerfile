FROM golang:latest AS build

WORKDIR /app

ADD . /app

RUN go get -u -v github.com/kardianos/govendor \
    && govendor sync \
    && GOOS=linux GOARCH=386 go build -v -o wxbot

FROM scratch

COPY --from=build /app/wxbot .

ENTRYPOINT ["./wxbot"]

