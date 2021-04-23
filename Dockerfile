FROM hewenyulucky/gobuild:1.16 as builder
WORKDIR /workspace/

LABEL Version=1.0
LABEL Name=simple

ENV GOPROXY=https://goproxy.cn,direct

ENV LANG en_US.UTF-8

COPY . .

RUN go build -o server .

FROM debian:stable-slim

LABEL Version=1.0
LABEL Name=simple
LABEL maintainer="hewenyu<yuebanlaosiji@outlook.com>"

WORKDIR /workspace/


ENV GIN_MODE="release"

ENV POSTGRES_HOST="localhost" \
    POSTGRES_USER="gorm" \
    POSTGRES_PWD="gorm" \
    POSTGRES_DB="gorm" \
    POSTGRES_PORT="9920" \
    POSTGRES_SSLMODE="disable"


# 拷贝可执行文件
COPY --from=builder  /workspace/gin-simple /workspace/

EXPOSE 8888

CMD ["./server"]