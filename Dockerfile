# 该镜像需要依赖的基础镜像
FROM golang:latest as builder

WORKDIR /var/go/src/testGo

#将容器外项目文件拷贝至容器中
COPY . .

RUN go build -tags netgo -mod=vendor -o testGo main.go

# 声明服务运行在8080端口
EXPOSE 8000

FROM alpine:latest
WORKDIR /var/go/src/testGo

COPY --from=builder /var/go/src/testGo/testGo  .

RUN chmod 777 ./testGo

CMD ["./testGo", "run"]
# 指定维护者的名字
MAINTAINER leixiaotian