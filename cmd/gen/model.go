package gen

type TableColumn struct {
	Field   string `gorm:"column:Field"`   // 字段名称
	Type    string `gorm:"column:Type"`    // 字段类型
	Null    string `gorm:"column:Null"`    // 是否空
	Key     string `gorm:"column:Key"`     // 索引
	Default string `gorm:"column:Default"` // 默认值
	Extra   string `gorm:"column:Extra"`   // 扩展
	Comment string `gorm:"column:Comment"` // 备注
}

type GenReq struct {
	TableName     string
	SrvName       string

	BaseDir       string
	DomainDir     string
	ModelDir      string
	RepositoryDir string
	ServiceDir    string
	HandlerDir    string
	InitializeDir string
	ProtoDir      string

	TableColumns []TableColumn
}
