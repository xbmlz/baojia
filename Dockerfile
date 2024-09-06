# Build
FROM golang:1.21-alpine AS builder

WORKDIR /src

COPY . .

RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o baojia .

# Run
FROM alpine:latest


ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

RUN RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache bash ca-certificates tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo ${TIMEZONE} > /etc/timezone

WORKDIR /app

COPY --from=builder /src/baojia .
COPY --from=builder /src/public ./public
COPY --from=builder /src/templates ./templates

EXPOSE 8080

ENTRYPOINT ["./baojia"]