package gen

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jettjia/go-micro-frame-cli/util"
)

func GenProto(req GenReq) {
	doGenAutoProto(req) // 生成自动编译proto文件
	doGenCommon(req)    // 生成proto公共部分
	doGenSrv(req)       // 生成proto server部分
	doGenMessage(req)   // 生成表操作的具体逻辑部分
}

func doGenCommon(req GenReq) {
	str := `syntax = "proto3";
option go_package = ".;proto";

message Query {
    string key = 1; //表字段名称
    string value = 2; //表字段值
    Operator operator = 3; //判断条件
}

enum Operator {
    GT = 0; //大于
    EQUAL = 1; //等于
    LT = 2; //小于
    NEQ = 3; //不等于
    LIKE = 4; //模糊查询
    GTE = 5; // 大于等于
    LTE = 6; // 小于等于
}

message PageData {
    uint32 pageSize = 1; // 一页多少条数据
    uint32 page = 2; // 第几页数据
    uint32 totalNumber = 3; // 一共多少条数据
    uint32 totalPage = 4; // 一共多少页
}`
	util.WriteStringToFileMethod(req.ProtoDir+"/common.proto", str)
}

func doGenAutoProto(req GenReq) {
	str := `protoc --go_out=plugins=grpc:./ ./*.proto`
	util.WriteStringToFileMethod(req.ProtoDir+"/auto.bat", str)
}

func doGenSrv(req GenReq) {
	str := `syntax = "proto3";
import "google/protobuf/empty.proto";
import public "google/protobuf/timestamp.proto";
import "common.proto";
import "category.proto";

option go_package = ".;proto";

service GoodsSrv {
    // 分类
    rpc CreateCategory (CategoryInfoRequest) returns (CategoryInfoResponse); // 新建
    rpc DeleteCategory (CategoryDeleteRequest) returns (google.protobuf.Empty); // 删
    rpc UpdateCategory (CategoryInfoRequest) returns (google.protobuf.Empty); // 修改
    rpc FindCategoryById (CategoryFindByIdRequest) returns (CategoryInfoResponse); // 根据id查找
    rpc QueryPageCategory (CategoryQueryPageRequest) returns (CategoryQueryPageResponse); // 分页List
}`
	var newStr string
	newStr = strings.Replace(str, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newStr = strings.Replace(newStr, "category", req.TableName, -1)

	util.WriteStringToFileMethod(req.ProtoDir+"/"+req.SrvName+".proto", newStr)
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
		str += "    "+fmt.Sprintf("%s %s = %d;//%s", typeName, v.Field, k+1, v.Comment) + "\n"
	}
	str += "\n"
	str += `}`
	str += "\n"


	// res
	str += `message CategoryInfoResponse {`
	str += "\n"
	for k, v := range req.TableColumns {
		typeName := genFieldForProtoMessage(v)
		str += "    "+fmt.Sprintf("%s %s = %d;//%s", typeName, v.Field, k+1, v.Comment) + "\n"
	}
	str += "\n"
	str += `}`
	str += "\n"

	str += `message CategoryDeleteRequest {
    uint64 id = 1;
}

message CategoryFindByIdRequest {
    uint64 id = 1;
}

message CategoryQueryPageRequest {
    repeated Query conditions = 1;
    PageData page = 2;
}

message CategoryQueryPageResponse {
    repeated CategoryInfoResponse pageList = 1;
    PageData page = 2;
}`

	newStr := strings.Replace(str, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	util.WriteStringToFileMethod(req.ProtoDir+"/"+req.TableName+".proto", newStr)
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
