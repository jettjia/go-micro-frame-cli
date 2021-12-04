package gen

const commonTemplateContext = `
package model

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"

	goodsProto "mall.com/mall-common/proto/goods"
)

var OperatorMap = map[goodsProto.Operator]string{
	goodsProto.Operator_GT:    " > ",
	goodsProto.Operator_EQUAL: " = ",
	goodsProto.Operator_LT:    " < ",
	goodsProto.Operator_NEQ:   " != ",
	goodsProto.Operator_LIKE:  " like ",
	goodsProto.Operator_GTE:   " >= ",
	goodsProto.Operator_LTE:   " <= ",
}

func GenerateQueryCondition(conditions []*goodsProto.Query) string {
	var condition string
	for k, v := range conditions {
		if k > 0 {
			condition += " and "
		}

		if v.Operator == goodsProto.Operator_LIKE {
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
}
`
