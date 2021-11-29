package gen

import "testing"

func Test_gen(t *testing.T) {
	host := "10.4.7.71"
	user := "root"
	password := "root"
	port := "3307"
	db := "zhe_pms"
	table := "category_attr"
	serverName := "goods-srv"
	protoName := "goods"
	Run(host, user, password, port, db, table, serverName, protoName)
}
