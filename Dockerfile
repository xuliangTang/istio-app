FROM golang:1.19-alpine as builder
RUN mkdir /src
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#RUN apk add build-base
ADD . /src
WORKDIR /src
RUN GOPROXY=https://goproxy.cn go build -o userapp cmd/UserApp.go && chmod +x userapp


FROM alpine:3.12
ENV ZONEINFO=/app/zoneinfo.zip
RUN mkdir /app
WORKDIR /app

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /app

COPY --from=builder /src/userapp /app

ENTRYPOINT  ["./userapp"]
EXPOSE 80