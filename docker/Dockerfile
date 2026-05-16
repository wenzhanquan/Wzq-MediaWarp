FROM alpine:latest

# GoReleaser 会自动将编译好的二进制文件放在根目录，直接拷贝即可
COPY MediaWarp /MediaWarp

RUN chmod +x /MediaWarp

EXPOSE 9000
VOLUME ["/etc/localtime", "/etc/timezone", "/config", "/logs", "/custom"]
ENTRYPOINT ["/MediaWarp"]
