FROM alpine

# ENV GOPROXY https://goproxy.cn/

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

COPY ./webserver /webserver
COPY ./config/settings.yml /config/settings.yml

EXPOSE 8081
RUN  chmod +x /webserver
CMD ["/webserver","server","-c", "/config/settings.yml"]
