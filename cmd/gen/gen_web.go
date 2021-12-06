package gen

import (
	"fmt"
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenWeb(req GenReq) {
	runGenApi(req)
	runGenRouter(req)
	runGenInitSrvConn(req)
	runGenInitRouter(req)
	runGenVo(req)
	runGenDto(req)
}

// runGenApi api逻辑
func runGenApi(req GenReq) {
	path := req.WebBaseDir + "/api/" + req.TableName + ".go"
	context := gstr.ReplaceByMap(webApiTemplateContext, g.MapStrStr{
		"ProductLog":                       GetJsonTagFromCase(req.TableName, "Camel"),
		"GoodsSrvClient":                   GetJsonTagFromCase(req.SrvName, "Camel") + "Client",
		"goodsProto":                       req.ProtoName + "Proto",
		"mall.com/mall-common/proto/goods": "mall.com/mall-common/proto/" + req.ProtoName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// runGenInitSrvConn 链接grpc server
func runGenInitSrvConn(req GenReq) {
	path := req.WebBaseDir + "/initialize/srv_conn.go"
	context := gstr.ReplaceByMap(webInitializeSrvConnContext, g.MapStrStr{
		"GoodsSrvClient": GetJsonTagFromCase(req.SrvName, "Camel") + "Client",
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// runGenInitRouter 路由init
func runGenInitRouter(req GenReq) {
	path := req.WebBaseDir + "/initialize/router.go"
	context := gstr.ReplaceByMap(webInitializeRouterContext, g.MapStrStr{
		"GoodsRouter": GetJsonTagFromCase(req.ProtoName, "Camel") + "Router",
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// runGenRouter 路由
func runGenRouter(req GenReq) {
	path := req.WebBaseDir + "/router/router.go"
	context := gstr.ReplaceByMap(webRouterContext, g.MapStrStr{
		"goods":       req.ProtoName,
		"GoodsRouter": GetJsonTagFromCase(req.ProtoName, "Camel") + "Router",
		"productLog":  GetJsonTagFromCase(req.TableName, "CamelLower"),
		"ProductLog":  GetJsonTagFromCase(req.TableName, "Camel"),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// runGenDto 参数
func runGenDto(req GenReq) {
	columnStr := "\n"
	for _, v := range req.TableColumns {
		if v.Field == "id" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}
		colStr := generateStructFieldForDto(v)
		columnStr += colStr + "\n"
	}

	path := req.WebBaseDir + "/trans/dto/" + req.TableName + ".go"
	context := gstr.ReplaceByMap(webDtoContext, g.MapStrStr{
		"{{dtoCreate}}":                    columnStr,
		"goodsProto":                       req.ProtoName + "Proto",
		"mall.com/mall-common/proto/goods": "mall.com/mall-common/proto/" + req.ProtoName,
		"ProductLog":                       GetJsonTagFromCase(req.TableName, "Camel"),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// runGenVo 返回
func runGenVo(req GenReq) {
	columnStr := "\n"
	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}

		fieldName := GetJsonTagFromCase(v.Field, "Camel")
		typeName := generateStructFieldTypeName(v)
		jsonTag := `json:"` + GetJsonTagFromCase(v.Field, "CamelLower") + `"`

		colStr := fieldName + " " + typeName + " " + "`" + jsonTag + "`"
		columnStr += colStr + "\n"
	}

	path := req.WebBaseDir + "/trans/vo/" + req.TableName + ".go"
	context := gstr.ReplaceByMap(webVoContext, g.MapStrStr{
		"{{voInfo}}":                       columnStr,
		"goodsProto":                       req.ProtoName + "Proto",
		"mall.com/mall-common/proto/goods": "mall.com/mall-common/proto/" + req.ProtoName,
		"ProductLog":                       GetJsonTagFromCase(req.TableName, "Camel"),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func generateStructFieldForDto(field TableColumn) (colStr string) {
	var fieldName, typeName, vTyp, node string

	fieldName = GetJsonTagFromCase(field.Field, "Camel")
	typeName = generateStructFieldTypeName(field)

	if field.Null == "YES" {
		fmt.Println("=====================")
		if fieldName == "uint64" || fieldName == "int64" || fieldName == "uint32" || fieldName == "int32" || fieldName == "int" {
			vTyp = `v:"integer|min:1"`
		} else if fieldName == "string" {
			vTyp = `v:"length:0,?"`
		} else if fieldName == "float32" || fieldName == "float64" {
			vTyp = `v:"float"`
		} else if fieldName == "time.Time" {
			vTyp = `v:"datetime"`
		}
	}

	node = `dc:"` + field.Comment + `"`

	colStr = fieldName + "    " + typeName + " `" + vTyp + node + "`"

	return
}
