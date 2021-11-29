package redis

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

func RunRedis() {
	mlog.Print("init redis:" + constant.RedisVersion + " start...")
	// docker pull image
	has, _ := gproc.ShellExec("docker images -q redis:"+constant.RedisVersion)
	if has == "" {
		_, err := gproc.ShellExec("sudo docker pull redis:" + constant.RedisVersion)
		if err != nil {
			mlog.Fatal("pull redis image err", err)
			return
		}
	}

	_, _ = gproc.ShellExec("sudo mkdir -p /mydata/redis/conf")
	_, _ = gproc.ShellExec("sudo touch /mydata/redis/conf/redis.conf")
	_, _ = gproc.ShellExec("sudo touch /mydata/redis/conf/redis.conf")
	_, _ = gproc.ShellExec(`echo -e "appendonly yes" >> /mydata/redis/conf/redis.conf`)
	_, err := gproc.ShellExec(`docker run -p 6379:6379 --name ` + constant.RedisName + ` \
-v /mydata/redis/data:/data \
-v /mydata/redis/conf/redis.conf:/etc/redis/redis.conf \
-d redis redis-server /etc/redis/redis.conf`)
	if err != nil {
		mlog.Fatal("docker run redis err", err)
		return
	}

	mlog.Print("done!")
}
