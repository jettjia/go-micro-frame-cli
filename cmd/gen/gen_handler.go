package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)

func GenHandler(req GenReq) {
	doBase(req)
	// todo ,还差生成 category.go这种
}

var (
	baseHandler = `
package handler

import (
	"goods_srv/domain/service"
	goods_srv "goods_srv/proto/goods_srv"
)

type GoodsSrv struct {
	goods_srv.UnimplementedGoodsSrv
	CategoryService          service.ICategoryService
}

`
)
func doBase(req GenReq) {
	var newStr string
	newStr = strings.Replace(baseHandler, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newStr = strings.Replace(baseHandler, "goods_srv", req.SrvName, -1)

	util.WriteStringToFileMethod(req.HandlerDir+"/base.go", newStr)
}
