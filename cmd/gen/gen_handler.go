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
		"category":                  GetJsonTagFromCase(req.TableName, "CamelLower"),
		"goods-srv":                 req.SrvName,
		"Category":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"goodsProto":                req.ProtoName + "Proto",
		"mall.com/mall-proto/goods": "mall.com/mall-proto/" + req.ProtoName,
		"{{create}}":                doGenHandlerCreate(req),
		"{{update}}":                doGenHandlerUpdate(req),
		"{{find}}":                  doGenHandlerFindById(req),
		"{{pageOne}}":               doGenHandlerPageOne(req),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// doGenHandlerCreate 创建
func doGenHandlerCreate(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "id" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}
		str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= req." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
	}

	return
}

// doGenHandlerUpdate 修改
func doGenHandlerUpdate(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}
		str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= req." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
	}

	return
}

// doGenHandlerFindById 单条
func doGenHandlerFindById(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}
		str += "reply." + GetJsonTagFromCase(v.Field, "Camel") + "=" + modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
	}

	return
}

// doGenHandlerPageOne page
func doGenHandlerPageOne(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}
		str += GetJsonTagFromCase(v.Field, "Camel") + ":" + "v." + GetJsonTagFromCase(v.Field, "Camel") + ",\n"
	}

	return
}
