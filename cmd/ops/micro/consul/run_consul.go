package consul

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunConsul() {
	mlog.Print("init consul:" + constant.ConsulVersion + " start...")

	// docker pull image
	has, _ := gproc.ShellExec("docker images -q consul:"+constant.ConsulVersion)
	if has == "" {
		_, err := gproc.ShellExec("sudo docker pull consul:" + constant.ConsulVersion)
		if err != nil {
			mlog.Fatal("pull consul image err", err)
			return
		}
	}

	// docker run
	_, err := gproc.ShellExec("sudo docker run --name " + constant.ConsulName + " -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent -dev -client=0.0.0.0")
	if err != nil {
		mlog.Fatal("run consul err", err)
		return
	}

	// yum install dns
	_, _ = gproc.ShellExec("sudo yum install bind-utils")

	mlog.Print("http://ip:8500")
	mlog.Print("done!")
}
