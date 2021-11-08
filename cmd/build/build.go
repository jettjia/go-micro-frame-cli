package build

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gproc"
)

func Run(args []string, n string) {
	file := args[0]

	fmt.Println("start building...")

	// build windows
	{
		fmt.Println("build for windows amd64")

		buildCommand := fmt.Sprintf("go build %s", file)
		fmt.Println(buildCommand)
		_, err := gproc.ShellExec(buildCommand)
		if err != nil {
			fmt.Println("build for win amd64 err: ", err.Error())
		}
	}

	// build linux
	{
		fmt.Println("build for linux amd64")
		_, _ = gproc.ShellExec("set CGO_ENABLED=0")
		_, _ = gproc.ShellExec("set GOOS=linux")
		_, _ = gproc.ShellExec("set GOARCH=amd64")

		buildCommand := fmt.Sprintf("go build -ldflags \"-w -s\" -o %s %s", n,file)

		fmt.Println(buildCommand)
		_, err := gproc.ShellExec(buildCommand)
		if err != nil {
			fmt.Println("build for linux amd64 err: ", err.Error())
		}
		_, _ = gproc.ShellExec("set GOOS=")
		_, _ = gproc.ShellExec("set GOARCH=")
	}

	fmt.Println("done!")
}
