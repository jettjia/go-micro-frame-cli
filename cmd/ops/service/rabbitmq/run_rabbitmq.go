package rabbitmq

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunRabbit() {
	mlog.Print("init rabbitmq:" + constant.RabbitmqVersion + " start...")


	// docker pull image
	has, _ := gproc.ShellExec("docker images -q rabbitmq:"+constant.RabbitmqVersion)
	if has == "" {
		_, err := gproc.ShellExec("sudo docker pull rabbitmq:" + constant.RabbitmqVersion)
		if err != nil {
			mlog.Fatal("pull rabbitmq image err", err)
			return
		}
	}

	_, err := gproc.ShellExec("docker run -d --name "+constant.RabbitmqName+" " +
		"-p 5672:5672 -p 15672:15672 -v /docker-data/rabbitmq:/var/lib/rabbitmq  -e RABBITMQ_DEFAULT_USER=admin " +
		"-e RABBITMQ_DEFAULT_PASS=123456 rabbitmq:"+constant.RabbitmqVersion)
	if err != nil {
		mlog.Fatal("docker run rabbitmq err:", err)
		return
	}

	mlog.Print("The background account password is admin/123456，Please keep it properly")
	mlog.Print("done!")
}
