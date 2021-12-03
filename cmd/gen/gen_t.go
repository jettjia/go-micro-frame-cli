package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// GenT 生成测试案例
func GenT(req GenReq) {
	runGen(req)
	runGenClient(req)
}

func runGen(req GenReq) {
	path := req.TestDir + "/" + req.TableName + "_test.go"

	context := gstr.ReplaceByMap(tContext, g.MapStrStr{
		"CategoryAttr":              GetJsonTagFromCase(req.TableName, "Camel"),
		"goodsProto":                req.ProtoName + "Proto",
		"mall.com/mall-proto/goods": "mall.com/mall-proto/" + req.ProtoName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func runGenClient(req GenReq) {
	path := req.TestDir + "/client.go"

	context := gstr.ReplaceByMap(tClientContext, g.MapStrStr{
		"192.168.106.1":             util.GetOutboundIP(),
		"goodsProto":                req.ProtoName + "Proto",
		"mall.com/mall-proto/goods": "mall.com/mall-proto/" + req.ProtoName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
