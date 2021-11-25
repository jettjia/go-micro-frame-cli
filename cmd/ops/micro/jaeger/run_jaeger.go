package jaeger

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunJaeger() {
	mlog.Print("init jaeger:" + constant.JaegerVersion + " start...")

	// docker pull image
	_, err := gproc.ShellExec("sudo docker pull jaegertracing/all-in-one:" + constant.RedisVersion)
	if err != nil {
		mlog.Fatal("pull jaeger image err", err)
		return
	}

	// docker run
	_, err = gproc.ShellExec(`docker run -d --name ` + constant.JaegerName + ` \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:` + constant.JaegerVersion)
	if err != nil {
		mlog.Fatal("run jaeger err", err)
		return
	}

	mlog.Print("http://ip:16686/search")
	mlog.Print("done!")
}
