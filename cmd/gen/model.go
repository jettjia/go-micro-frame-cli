package gen

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

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
	ProtoName     string
	BaseDir       string
	DomainDir     string
	ModelDir      string
	RepositoryDir string
	ServiceDir    string
	HandlerDir    string
	InitializeDir string
	ProtoDir      string
	TestDir       string
	WebBaseDir    string

	TableColumns []TableColumn
}

func GenInit(srvName, tableName, protoName string) GenReq {

	baseDir := "auto-code/" + srvName
	domainDir := baseDir + "/domain"
	modelDir := domainDir + "/model"
	repositoryDir := domainDir + "/repository"
	serviceDir := domainDir + "/service"
	handlerDir := baseDir + "/handler"
	initializeDir := baseDir + "/initialize"
	protoDir := baseDir + "/proto" + "/" + protoName
	testDir := baseDir + "/test"
	webBaseDir := "auto-code/web"

	return GenReq{
		TableName:     tableName,
		SrvName:       srvName,
		ProtoName:     protoName,
		BaseDir:       baseDir,
		DomainDir:     domainDir,
		ModelDir:      modelDir,
		RepositoryDir: repositoryDir,
		ServiceDir:    serviceDir,
		HandlerDir:    handlerDir,
		InitializeDir: initializeDir,
		ProtoDir:      protoDir,
		TestDir:       testDir,
		WebBaseDir:    webBaseDir,
		TableColumns:  GetTableCol(tableName),
	}
}
