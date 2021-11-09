package gen

import (
	"go.uber.org/zap"

	"github.com/gogf/gf/util/gconv"
	mySql "github.com/jettjia/go-micro-frame/service/gmysql"
)

func GetTableCol(table string) []TableColumn {
	var tc []TableColumn
	DB.Raw("SHOW FULL COLUMNS FROM " + table).Scan(&tc)

	return tc
}

func InitDB(host, port, user, password, db string) {
	m := &mySql.Mysql{
		Host:     host,
		Port:     gconv.Int(port),
		User:     user,
		Password: password,
		Db:       db,
	}

	var err error
	DB, err = m.GetDB()
	if err != nil {
		zap.S().Error("db connect err:", err.Error())
	}
}
