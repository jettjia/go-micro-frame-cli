package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/util/gconv"

	mySql "github.com/jettjia/go-micro-frame/service/gmysql"
)

func GetTableCol(table string) []TableColumn {
	var tc []TableColumn
	DB.Raw("SHOW FULL COLUMNS FROM " + table).Scan(&tc)

	return tc
}

func InitDB(host, port, user, password, db string) {
	defer func() {
		if err := recover(); err != nil {
			mlog.Fatal("db connect err:%v", err)
		}
	}()
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
		mlog.Fatal("db connect err:%v", err)
	}
}
