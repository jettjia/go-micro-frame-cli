package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)

var (
	regSrv = `
// RegisterSrv 初始化 服务
package initialize

import (
	"google.golang.org/grpc"

	"goods_srv/domain/repository"
	service2 "goods_srv/domain/service"
	"goods_srv/handler"
	goods_srv "goods_srv/proto/goods_srv"
)

func RegisterSrv(server *grpc.Server) {

	categoryService := service2.NewCategoryService(repository.NewCategoryRepository())
	

	goods_srv.RegisterGoodsServer(server, &handler.GoodsServer{
		CategoryService:          categoryService,
	})
}
`
)

func GenInitlialize(req GenReq) {
	var newStr string
	newStr = strings.Replace(regSrv, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newStr = strings.Replace(newStr, "goods_srv", req.SrvName, -1)

	util.WriteStringToFileMethod(req.InitializeDir+"/registerSrv.go", newStr)
}