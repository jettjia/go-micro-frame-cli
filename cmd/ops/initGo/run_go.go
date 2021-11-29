package initGo

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/jettjia/go-micro-frame-cli/constant"
	"github.com/jettjia/go-micro-frame-cli/util"
)

func RunGo() {
	goversion := "go" + constant.GOVERSION
	goZipFile := goversion + ".linux-amd64.tar.gz"

	mlog.Print("init " + goversion + " start...")

	if !util.IsExists(goZipFile) {
		mlog.Printf(goZipFile + "downloading... ")
		_, err := gproc.ShellExec("sudo wget https://studygolang.com/dl/golang/" + goZipFile)
		if err != nil {
			mlog.Fatal("down "+goversion+" err", err)
		}
	}

	mlog.Printf("begin install...")
	_, _ = gproc.ShellExec("rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz")
	_, _ = gproc.ShellExec(`echo -e "\n" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export GO111MODULE=on" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export GOPROXY=https://goproxy.cn" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export GOROOT=/usr/local/go" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export GOPATH=/mnt/hgfs/go_work/wingopath" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export GOBIN=\$GOPATH/bin" >> /etc/profile`)
	_, _ = gproc.ShellExec(`echo -e "export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> /etc/profile`)
	_, _ = gproc.ShellExec(`source /etc/profile`)
	str, _ := gproc.ShellExec(`go version`)
	mlog.Print(str)
	mlog.Print("done!")
}
