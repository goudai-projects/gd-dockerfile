FROM ghcr.io/goudai-projects/openjre:11.0.7
MAINTAINER qingmu.io 247687009@qq.com

ENV CATALINA_HOME /usr/local/tomcat
ENV PATH $CATALINA_HOME/bin:$PATH
RUN mkdir -p $CATALINA_HOME && cd /usr/local \
&& apk add wget && wget https://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.57/bin/apache-tomcat-8.5.57.tar.gz \
&& apk del wget \
&& tar xf apache-tomcat-8.5.57.tar.gz \
&& rm -rf apache-tomcat-8.5.57.tar.gz \
&& mv apache-tomcat-8.5.57/* /usr/local/tomcat \
&& rm -rf apache-tomcat-8.5.57 \
&& rm -rf /var/cache/apk/*
WORKDIR $CATALINA_HOME
EXPOSE 8080
CMD ["catalina.sh", "run"]
