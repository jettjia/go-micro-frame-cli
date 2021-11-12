package nacos

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunNacos()  {
	mlog.Print("init nacos:" + constant.NacosVersion + " start...")

	// docker pull image
	_, err := gproc.ShellExec("sudo docker pull nacos/nacos-server:" + constant.NacosVersion)
	if err != nil {
		mlog.Fatal("pull nacos image err", err)
		return
	}

	// docker run
	_, err = gproc.ShellExec("docker run --name "+constant.NacosName+" -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:"+constant.NacosVersion)
	if err != nil {
		mlog.Fatal("run nacos err", err)
		return
	}

	mlog.Print("http://ip:8848/nacos/index.html")
	mlog.Print("The nacos account password is nacos/nacos，Please keep it properly")
	mlog.Print("done!")
}