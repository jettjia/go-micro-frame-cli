package util

import (
	"fmt"
	"github.com/jettjia/go-micro-frame-cli/util"
	"testing"
)

func Test_GetOutboundIP(t *testing.T) {
	str := util.GetOutboundIP()
	fmt.Println("IP地址是：", str)
}
