package gen

import (
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// GenModel 生成 model
func GenModel(req GenReq) {
	doGenModelBase(req)
	doGenModelCommon(req)
	doGenModel(req)
}

func doGenModelBase(req GenReq) {
	path := req.ModelDir + "/base.go"
	if err := gfile.PutContents(path, strings.TrimSpace(baseTemplateContext)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func doGenModelCommon(req GenReq) {
	path := req.ModelDir + "/common.go"
	context := gstr.ReplaceByMap(commonTemplateContext, g.MapStrStr{
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

func doGenModel(req GenReq) {
	columnStr := "\n"

	for _, v := range req.TableColumns {
		if v.Field == "id" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" {
			continue
		}

		colStr := generateStructFieldForModel(v)

		columnStr += colStr + "\n"
	}

	str := `package model

type ` + GetJsonTagFromCase(req.TableName, "Camel") + ` struct {
	BaseModel` + columnStr +
		`}`

	path := req.ModelDir + "/" + req.TableName + ".go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// generateStructFieldForModel 获取字段解析后的 type所索要的类型
// 要处理的目标
// CateName       string `gorm:"type:varchar(32); not null; default:0;comment:分类名称;" json:"cate_name" ` // 分类名称
// ormTag, gorm:"type:varchar(32); not null; comment:分类名称;"
// jsonTag, json:"cate_name"`
func generateStructFieldForModel(field TableColumn) (colStr string) {
	var fieldName, typeName, ormTag, jsonTag, node string
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint"
		} else {
			typeName = "int"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}

	case "real":
		typeName = "float32"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float64"

	case "bool":
		typeName = "bool"

	case "datetime", "timestamp", "date", "time":
		typeName = "time.Time"
	case "json":
		typeName = "string"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "int"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "float64"
		case strings.Contains(t, "bool"):
			typeName = "bool"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "[]byte"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			typeName = "time.Time"
		default:
			typeName = "string"
		}
	}

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`

	// note 如 // 分类名称
	node = " //" + field.Comment

	// ormTag 如 gorm:"column:category_name; type:varchar(32); not null; default:0; comment:分类名称;"
	ormTag = `gorm:"column:` + field.Field

	if gstr.ContainsI(field.Key, "pri") {
		ormTag += " ,primary"
	}
	if gstr.ContainsI(field.Key, "uni") {
		ormTag += " ,unique"
	}

	ormTag += "; type:" + field.Type + ";"

	if field.Null == "YES" {
		ormTag += "not null;"
	}

	if field.Default != "" {
		ormTag += "default: " + field.Default + ";"
	}

	if field.Comment != "" {
		ormTag += "comment:" + field.Comment + ";"
	}

	ormTag += `"`

	colStr = fieldName + "    " + typeName + "    " + "`" + ormTag + " " + jsonTag + "`" + node

	return
}

func GetJsonTagFromCase(str, caseStr string) string {
	switch gstr.ToLower(caseStr) {
	case gstr.ToLower("Camel"):
		return gstr.CaseCamel(str)

	case gstr.ToLower("CamelLower"):
		return gstr.CaseCamelLower(str)

	case gstr.ToLower("Kebab"):
		return gstr.CaseKebab(str)

	case gstr.ToLower("KebabScreaming"):
		return gstr.CaseKebabScreaming(str)

	case gstr.ToLower("Snake"):
		return gstr.CaseSnake(str)

	case gstr.ToLower("SnakeFirstUpper"):
		return gstr.CaseSnakeFirstUpper(str)

	case gstr.ToLower("SnakeScreaming"):
		return gstr.CaseSnakeScreaming(str)
	}
	return str
}

////////////////////////////////////////////////////////////////////////////////

func GenRepository(req GenReq) {
	doGenRepositoryBase(req)
	doGenRepository(req)
}

func doGenRepositoryBase(req GenReq) {
	path := req.RepositoryDir + "/base.go"
	if err := gfile.PutContents(path, strings.TrimSpace(repositoryBaseTemplate)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func doGenRepository(req GenReq) {
	path := req.RepositoryDir + "/" + req.TableName + "_repository.go"

	context := gstr.ReplaceByMap(repositoryTemplate, g.MapStrStr{
		"goods-srv":                 req.SrvName,
		"Category":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"category":                  GetJsonTagFromCase(req.TableName, "CamelLower"),
		"table_name":                req.TableName,
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

////////////////////////////////////////////////////////////////////////////////
func GenService(req GenReq) {
	path := req.ServiceDir + "/" + req.TableName + "_service.go"

	context := gstr.ReplaceByMap(serviceTemplateContext, g.MapStrStr{
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
