package docker

import (
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"strings"

	//"github.com/jettjia/go-micro-frame-cli/util"
)

var (
	oldStr = `# 表示依赖 alpine 最新版
FROM golang:1.16 AS build
MAINTAINER jettjia <jettjia@qq.com>

# 安装依赖包
RUN go mod tidy

# 复制源码并执行build，此处当文件有变化会产生新的一层镜像层
COPY . /go/src/srv-example/
## 拷贝配置文件到容器中
COPY grpc/config-dev.yaml /go/src/srv-example/config-dev.yaml
COPY grpc/config-prod.yaml /go/src/srv-example/config-prod.yaml

RUN go build -o /bin/srv-example

# 缩小到一层镜像
FROM busybox
COPY --from=build /bin/srv-example /bin/srv-example

# 暴露端口
EXPOSE 50051

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 运行golang程序的命令
ENTRYPOINT ["/bin/srv-example"]
`
)

// Run 当前项目生成 dockerfile 文件
func Run(args []string) {
	imgName := args[0]

	log.Println("create Dockerfile start ...")
	newStr := strings.Replace(oldStr, "srv-example", imgName, -1)

	gfile.PutContents("Dockerfile", newStr)
	log.Println("Done")
}
