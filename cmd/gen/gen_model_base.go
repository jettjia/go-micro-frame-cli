package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)


var (
	baseModel = `package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        uint64`         +"`" +`gorm:"column:created_id;primarykey;type:int;comment:主键" json:"id"` +"`\n"+
	`CreatedAt time.Time`     +"`"+ `gorm:"column:created_at;comment:创建时间" json:"-"` +"`\n" +
	`UpdatedAt time.Time`      +"`"+ `gorm:"column:updated_at;type:timestamp not null;default:current_timestamp;comment:修改时间" json:"-"` +"`\n" +
	`DeletedAt gorm.DeletedAt` +"`"+ `gorm:"column:deleted_at;comment:删除时间" json:"-"` +"`\n" +
`}
`

	commonModel = `package model

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"

	goods_proto "goods_proto/proto/goods_proto"
)

var OperatorMap = map[goods_proto.Operator]string{
	goods_proto.Operator_GT:    " > ",
	goods_proto.Operator_EQUAL: " = ",
	goods_proto.Operator_LT:    " < ",
	goods_proto.Operator_NEQ:   " != ",
	goods_proto.Operator_LIKE:  " like ",
	goods_proto.Operator_GTE:   " >= ",
	goods_proto.Operator_LTE:   " <= ",
}

func GenerateQueryCondition(conditions []*goods_proto.Query) string {
	var condition string
	for k, v := range conditions {
		if k > 0 {
			condition += " and "
		}

		if v.Operator == goods_proto.Operator_LIKE {
			condition += fmt.Sprintf("%v%s'%%%v%%'", v.Key, OperatorMap[v.Operator], v.Value)
		} else {
			//bool string int
			_, err := strconv.ParseBool(v.Value)
			if err != nil {
				condition += fmt.Sprintf("%v%s'%v'", v.Key, OperatorMap[v.Operator], v.Value)
			} else {
				condition += fmt.Sprintf("%v%s%v", v.Key, OperatorMap[v.Operator], v.Value)
			}
		}
	}

	return condition
}

// 参考 https://github.com/qicmsg/go_vcard
// 用法参考
/*  where := []interface{}{
	[]interface{}{"id", "in", []int{1, 2}},
	[]interface{}{"username", "=", "chen", "or"},
	}
	db, err = entity.BuildWhere(db, where)
	db.Find(&users)
	// SELECT * FROM users  where (id in ('1','2')) OR (username = 'chen')
*/
func BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	var err error
	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					return nil, errors.New("切片长度不能小于2")
				}
				columnstr := column.(string)
				// 拼接参数形式
				if strings.Index(columnstr, "?") > -1 {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/

					if strings.Index(" in notin ", _opt) > -1 {
						// val 是数组类型
						column = columnstr + " " + opt + " (?)"
					} else if strings.Index(" = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp", _opt) > -1 {
						column = columnstr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			} else if t == reflect.Map /*Map*/ {
				db = db.Where(item)
			} else {
				/*
					// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						db, err = BuildWhere(db, item)
						if err != nil {
							panic(err)
						}
						return db
					})*/

				db, err = BuildWhere(db, item)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		return nil, errors.New("参数有误")
	}
	return db, nil
}`
)

func DoGenBase(req GenReq)  {
	util.WriteStringToFileMethod(req.ModelDir+"/base.go", baseModel)
}


func DoGenCommon(req GenReq) {
	newcommonModel := strings.Replace(commonModel, "goods_proto", req.SrvName, -1)
	util.WriteStringToFileMethod(req.ModelDir+"/common.go", newcommonModel)
}