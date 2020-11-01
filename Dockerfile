FROM hub.c.163.com/library/golang
ENV  GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/WeCanRun/gin-blog
COPY . $GOPATH/src/github.com/WeCanRun/gin-blog
# RUN go build .
EXPOSE 8000
ENTRYPOINT go build . && ./gin-blog
