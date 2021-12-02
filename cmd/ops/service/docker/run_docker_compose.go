package docker

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/util"
)

func RunDockerCompose() {
	if !util.IsExists("/usr/local/bin/docker-compose") {
		_, err := gproc.ShellExec("sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose")
		if err != nil {
			mlog.Fatal("docker-compose download err:", err)
			return
		}
	}

	gproc.ShellExec("sudo chmod +x /usr/local/bin/docker-compose")

}
