package harbor

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/jettjia/go-micro-frame-cli/constant"
	"github.com/jettjia/go-micro-frame-cli/util"
)

// RunHarbor 安装 harbor
func RunHarbor() {
	// down
	mlog.Print("harbor init...")
	if !util.IsExists("/tmp/harbor-offline-installer-v"+constant.HarborVersion+".tgz") {
		_, err := gproc.ShellExec("sudo curl -L https://github.com/goharbor/harbor/releases/download/v"+constant.HarborVersion+"/harbor-offline-installer-v"+constant.HarborVersion+".tgz -o /tmp")
		if err != nil {
			mlog.Fatal("down harbor : ", err)
			return
		}
	}

	_, err := gproc.ShellExec("sudo mkdir /opt/harbor")
	if err != nil {
		mlog.Fatal("mkdir /opt/harbor err:", err)
		return
	}

	res, err := gproc.ShellExec("tar -zxf /tmp/harbor-offline-installer-v"+constant.HarborVersion+".tgz -C /opt/" )
	if err != nil {
		mlog.Fatal("tar -zxf /tmp/harbor-offline-installer err:", err)
		mlog.Fatal("tar -zxf /tmp/harbor-offline-installer res:", res)
		return
	}

	gproc.ShellExec("cp /opt/harbor/harbor.yml.tmpl /opt/harbor/harbor.yml")

	// 修改ip和端口
	gproc.ShellExec(`sed -i "s/hostname: reg.mydomain.com/hostname: ` + util.GetOutboundIP() + `/g"  /opt/harbor/harbor.yml`)
	gproc.ShellExec(`sed -i "s/port: 80/port: 85/g"  /opt/harbor/harbor.yml`)

	gproc.ShellExec(`sed -i "s/port: 443/#port: 443/g"  /opt/harbor/harbor.yml`)
	gproc.ShellExec(`sed -i "s/certificate:/#certificate:/g"  /opt/harbor/harbor.yml`)
	gproc.ShellExec(`sed -i "s/private_key:/#private_key:/g"  /opt/harbor/harbor.yml`)

	// 安装
	res, err = gproc.ShellExec("sh /opt/harbor/prepare")
	if err != nil {
		mlog.Fatal("sh /opt/harbor/prepare err:", err)
		mlog.Fatal("sh /opt/harbor/prepare res:", res)
		return
	}

	res, err = gproc.ShellExec("sh /opt/harbor/install.sh  --with-chartmuseum")
	if err != nil {
		mlog.Fatal("sh /opt/harbor/install.sh  --with-chartmuseum err:", err)
		mlog.Fatal("sh /opt/harbor/install.sh  --with-chartmuseum res:", res)
		return
	}

	res, err = gproc.ShellExec("cd /opt/harbor && docker-compose up -d")
	if err != nil {
		mlog.Fatal("docker-compose up -d err:", err)
		mlog.Fatal("docker-compose up -d res:", res)
		return
	}

	mlog.Print("http://ip:85")
	mlog.Print("The harbor account password is admin/Harbor12345，Please keep it properly")
	mlog.Print("done!")
}
