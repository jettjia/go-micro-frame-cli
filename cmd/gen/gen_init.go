package gen

import "github.com/jinzhu/gorm"

func GenInit(srvName, tableName string) GenReq {

	baseDir := "auto_code/" + srvName
	domainDir := baseDir + "/domain"
	modelDir := domainDir + "/model"
	repositoryDir := domainDir + "/repository"
	serviceDir := domainDir + "/service"
	handlerDir := baseDir + "/handler"
	initializeDir := baseDir + "/initialize"
	protoDir := baseDir + "/proto" + "/" + srvName

	return GenReq{
		TableName: tableName,
		SrvName:   srvName,

		BaseDir:       baseDir,
		DomainDir:     domainDir,
		ModelDir:      modelDir,
		RepositoryDir: repositoryDir,
		ServiceDir:    serviceDir,
		HandlerDir:    handlerDir,
		InitializeDir: initializeDir,
		ProtoDir:      protoDir,
		TableColumns : GetTableCol(tableName),
	}
}

var (
	DB           *gorm.DB
)
