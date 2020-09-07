FROM alpine:3.12.0
MAINTAINER qingmu.io 247687009@qq.com
ENV LANG=C.UTF-8 \
    TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
&& echo $TZ > /etc/timezone \
&& sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk add --update procps \
&& apk add --no-cache  ca-certificates ttf-dejavu tzdata tini bash
# 变动层
RUN apk add --no-cache  openjdk8-jre
ENV JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk/jre \
    PATH=$PATH:/usr/lib/jvm/java-1.8-openjdk/jre/bin:/usr/lib/jvm/java-1.8-openjdk/bin
RUN apk add --no-cache fontconfig font-adobe-100dpi
RUN rm -rf /var/cache/apk/*