package es

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunEs() {
	mlog.Print("init es:" + constant.EsVersion + " start...")

	// pull image
	has, _ := gproc.ShellExec("docker images -q elasticsearch:"+constant.EsVersion)
	if has == "" {
		_, err := gproc.ShellExec("sudo docker pull elasticsearch:" + constant.EsVersion)
		if err != nil {
			mlog.Fatal("pull mysql image err: ", err)
			return
		}
	}

	// run
	_, err := gproc.ShellExec(`docker run -d --name `+constant.EsName+`  -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:`+constant.EsVersion)
	if err != nil {
		mlog.Printf("run es err: ", err)
	}

	mlog.Print("done!")
}
