package drone

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/os/gproc"

	"github.com/jettjia/go-micro-frame-cli/util"
)

func RunDrone() {
	// install drone
	_, err := gproc.ShellExec(`docker run --name=drone \
--volume=/data/drone:/data \
--env=DRONE_AGENTS_ENABLED=true \
--env=DRONE_GOGS_SERVER=http://10.4.7.100:3000 \
--env=DRONE_RPC_SECRET=123456key \
--env=DRONE_SERVER_HOST=10.4.7.100:8080 \
--env=DRONE_SERVER_PROTO=http \
--env=TZ=PRC \
--env=DRONE_DATABASE_DRIVER=mysql \
--env=DRONE_DATABASE_DATASOURCE="root:123456@tcp(` + util.GetOutboundIP() + `:3306)/drone?parseTime=true" \
--publish=8080:80 \
--publish=8443:443 \
--detach=true \
--restart=always \
--restart always \
-d drone/drone
`)
	if err != nil {
		mlog.Fatal("docker run drone:", err)
		return
	}

	// install drone-runner
	_, err = gproc.ShellExec(`docker run -d --name drone-runner \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /root/.ssh:/root/.ssh \
  -e DRONE_RPC_PROTO=http \
  -e DRONE_RPC_HOST=10.4.7.100:8080 \
  -e DRONE_RPC_SECRET=123456key \
  -e DRONE_RUNNER_CAPACITY=2 \
  -e DRONE_RUNNER_NAME=drone-runner \
  -e DRONE_DATABASE_DRIVER=mysql \
  -e DRONE_DATABASE_DATASOURCE="root:123456@tcp(` + util.GetOutboundIP() + `:3306)/drone?parseTime=true" \
  -p 13000:3000 \
  --restart always \
  drone/drone-runner-docker`)
	if err != nil {
		mlog.Fatal("docker run drone-runner:", err)
		return
	}

	// install ssh-runner
	_, err = gproc.ShellExec(`docker run -d \
  -e DRONE_RPC_PROTO=http \
  -e DRONE_RPC_HOST=` + util.GetOutboundIP() + `:8080 \
  -e DRONE_RPC_SECRET=123456key \
  -e DRONE_DEBUG=true \
  -p 10081:3000 \
  --restart always \
  --name ssh-runner \
  drone/drone-runner-ssh
`)
	if err != nil {
		mlog.Fatal("docker run ssh-runner:", err)
		return
	}


	mlog.Print("http://" + util.GetOutboundIP() + "8080/")
	mlog.Print("The Drone account password is your gogs...，Please keep it properly")
	mlog.Print("done!")
}
