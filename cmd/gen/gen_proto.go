package gen

import (
	"fmt"
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenProto(req GenReq) {
	doGenCommon(req)    // 生成proto公共部分
	doGenSrv(req)       // 生成proto server部分
	doGenMessage(req)   // 生成表操作的具体逻辑部分
	doGenAutoProto(req) // 生成自动编译proto文件
}

func doGenCommon(req GenReq) {
	path := req.ProtoDir + "/common.proto"

	context := gstr.ReplaceByMap(protoCommonTemplateContext, g.MapStrStr{
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func doGenAutoProto(req GenReq) {
	str := `protoc --go_out=plugins=grpc:./ ./*.proto`
	gfile.PutContents(req.ProtoDir+"/auto.bat", str)
}

func doGenSrv(req GenReq) {
	path := req.ProtoDir + "/" + req.ProtoName + ".proto"

	context := gstr.ReplaceByMap(protoSrvTemplateContext, g.MapStrStr{
		"Category": GetJsonTagFromCase(req.TableName, "Camel"),
		"category": req.TableName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func doGenMessage(req GenReq) {
	str := `syntax = "proto3";
option go_package = ".;proto";
import "common.proto";
import public "google/protobuf/timestamp.proto";
`

	str += "\n"

	// req
	str += `message CategoryInfoRequest {`
	str += "\n"
	for k, v := range req.TableColumns {
		typeName := genFieldForProtoMessage(v)
		str += "    " + fmt.Sprintf("%s %s = %d;//%s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), k+1, v.Comment) + "\n"
	}
	str += "\n"
	str += `}`
	str += "\n"

	// res
	str += `message CategoryInfoResponse {`
	str += "\n"
	for k, v := range req.TableColumns {
		typeName := genFieldForProtoMessage(v)
		str += "    " + fmt.Sprintf("%s %s = %d;//%s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), k+1, v.Comment) + "\n"
	}
	str += "\n"
	str += `}`
	str += "\n"

	str += `message DeleteCategoryRequest {
    uint64 id = 1;
}

message FindCategoryRequest {
    uint64 id = 1;
}

message QueryPageCategoryRequest {
    repeated Query conditions = 1;
    PageData page = 2;
}

message QueryPageCategoryResponse {
    repeated CategoryInfoResponse pageList = 1;
    PageData page = 2;
}`

	newStr := strings.Replace(str, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	gfile.PutContents(req.ProtoDir+"/"+req.TableName+".proto", newStr)
}

func genFieldForProtoMessage(field TableColumn) (colStr string) {
	var typeName string

	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)

	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "bytes"
	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint32"
		} else {
			typeName = "int32"
		}
	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}
	case "real":
		typeName = "float"
	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float"
	case "bool":
		typeName = "bool"
	case "datetime", "timestamp", "date", "time":
		typeName = "google.protobuf.Timestamp"
	case "varchar", "longtext", "text":
		typeName = "string"
	case "json":
		typeName = "string"
	}

	return typeName
}
