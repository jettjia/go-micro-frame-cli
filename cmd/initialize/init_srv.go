package initialize

import (
	"context"
	"log"
	"strings"

	"github.com/gogf/gf-cli/v2/library/allyes"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	emptySrvProjectName = "https://github.com/jettjia/go-micro-frame-demos"
)

// InitSrv init srv project
func InitSrv(args []string) {
	projectName := args[0]

	dirPath := projectName
	if !gfile.IsEmpty(dirPath) && !allyes.Check() {
		s := gcmd.Scanf(`the folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, projectName)
		if strings.EqualFold(s, "n") {
			return
		}
	}

	log.Println("initializing...")
	//从github 拉取项目
	respMd5, err := g.Client().Get(context.TODO(), emptySrvProjectName+"/archive/refs/heads/master.zip")
	if err != nil {
		log.Printf("get the project zip md5 failed: %s\n", err.Error())
		return
	}
	if respMd5 == nil {
		log.Println("got the project zip md5 failed")
		return
	}

	defer respMd5.Close()
	md5DataStr := respMd5.ReadAllString()

	if md5DataStr == "" {
		log.Println("get the project zip md5 failed: empty md5 value. maybe network issue, try again?")
		return
	}

	// Unzip the zip data.
	if err = gcompress.UnZipContent([]byte(md5DataStr), projectName); err != nil {
		log.Println("unzip project data failed,", err.Error())
		return
	}

	log.Println("initialization done! ")
}
