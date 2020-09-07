FROM ghcr.io/goudai-projects/openjre:11.0.7
MAINTAINER qingmu.io 247687009@qq.com

ENV CATALINA_HOME /usr/local/tomcat
ENV PATH $CATALINA_HOME/bin:$PATH
RUN mkdir -p $CATALINA_HOME && cd /usr/local \
&& apk add wget && wget https://mirror.bit.edu.cn/apache/tomcat/tomcat-9/v9.0.37/bin/apache-tomcat-9.0.37.tar.gz \
&& apk del wget \
&& tar xf apache-tomcat-9.0.37.tar.gz \
&& rm -rf apache-tomcat-9.0.37.tar.gz \
&& mv apache-tomcat-9.0.37/* /usr/local/tomcat \
&& rm -rf apache-tomcat-9.0.37 \
&& rm -rf /var/cache/apk/*
WORKDIR $CATALINA_HOME
EXPOSE 8080
CMD ["catalina.sh", "run"]
