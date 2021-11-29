package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

func GenHandler(req GenReq) {
	doBase(req)
	doHandler(req)
}

func doBase(req GenReq) {
	path := req.HandlerDir + "/base.go"

	context := gstr.ReplaceByMap(handlerBaseTemplateContext, g.MapStrStr{
		"goods-srv":   req.SrvName,
		"Category":    GetJsonTagFromCase(req.TableName, "Camel"),
		"GoodsServer": GetJsonTagFromCase(req.ProtoName, "Camel") + "Server",
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func doHandler(req GenReq) {
	path := req.HandlerDir + "/" + req.TableName + ".go"

	context := gstr.ReplaceByMap(handlerTemplateContext, g.MapStrStr{
		"goods-srv":                 req.SrvName,
		"Category":                  GetJsonTagFromCase(req.TableName, "Camel"),
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
