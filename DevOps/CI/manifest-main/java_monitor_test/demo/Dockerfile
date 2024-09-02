FROM registry.cn-hangzhou.aliyuncs.com/s-ops/alpine-jdk8:latest
WORKDIR /opt
COPY target/*.jar /opt/app.jar
RUN apk add --no-cache tini && rm -rf /var/cache/apk/*
#ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/opt/app.jar"]
ENTRYPOINT ["sh","-c","exec java $java_ops -jar /opt/app.jar"]
#ENTRYPOINT ["/sbin/tini", "--", "java", "-jar", "/opt/app.jar"]
#ENTRYPOINT ["/sbin/tini", "--","sh","-c","exec java -Djava.security.egd=file:/dev/./urandom -jar /opt/app.jar"]
