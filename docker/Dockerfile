FROM golang:1.16-alpine as builder

WORKDIR /go/influxdb-admin

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirror.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    go env -w GOPROXY=https://goproxy.cn,direct

RUN apk add --no-cache git make && make


FROM alpine

COPY --from=builder /go/influxdb-admin/bin/influxdb-admin /usr/local/bin/

RUN sed -i 's/dl-cdn.alpinelinux.org/mirror.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache ca-certificates tzdata bash

EXPOSE 8083

COPY docker/entrypoint.sh /usr/local/bin/

ENTRYPOINT ["entrypoint.sh"]

CMD ["influxdb-admin"]
