FROM alpine:3.12.0
MAINTAINER qingmu 247687009@qq.com
ENV LANG=C.UTF-8 \
    TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
&& echo $TZ > /etc/timezone \
&& sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk add --update procps \
&& apk add --no-cache  ca-certificates ttf-dejavu tzdata tini bash
# 变动层
RUN apk add --no-cache  openjdk11-jre
ENV JAVA_HOME=/usr/lib/jvm/java-11-openjdk \
  PATH=$PATH:/usr/lib/jvm/java-11-openjdk/bin
RUN rm -rf /var/cache/apk/*



# imagemagick font-adobe-100dpi ttf-dejavu fontconfig
