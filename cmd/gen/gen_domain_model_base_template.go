package gen

const baseTemplateContext = `
package model

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
	Id        uint64`         +"`" +`gorm:"column:created_id;primarykey;type:int;comment:主键" json:"id"` +"`\n"+
	`CreatedAt time.Time`     +"`"+ `gorm:"column:created_at;comment:创建时间" json:"-"` +"`\n" +
	`UpdatedAt time.Time`      +"`"+ `gorm:"column:updated_at;type:timestamp not null;default:current_timestamp;comment:修改时间" json:"-"` +"`\n" +
	`DeletedAt gorm.DeletedAt` +"`"+ `gorm:"column:deleted_at;comment:删除时间" json:"-"` +"`\n" +
	`}
`
