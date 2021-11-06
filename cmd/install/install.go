package install

import (
	"runtime"

	"github.com/jettjia/go-micro-frame-cli/util"
)

func Install() {
	goroot := runtime.GOROOT()

	src := ""
	dst := ""

	switch runtime.GOOS {
	case "windows":
		src = "./go-micro-frame-cli.exe"
		dst = `C:\Program Files` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + "/bin" + src
		}

	default:
		src = "./go-micro-frame-cli"
		dst = `/usr/local/bin` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + src
		}
	}

	_, _ = util.Copy(src, dst)

	return
}
