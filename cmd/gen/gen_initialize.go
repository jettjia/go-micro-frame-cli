package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

var (
	regSrv = `
// RegisterSrv 初始化 服务
package initialize

import (
	"google.golang.org/grpc"

	goodsProto "mall.com/mall-common/proto/goods"
	"mall.com/srv/goods-srv/domain/repository"
	service2 "mall.com/srv/goods-srv/domain/service"
	"mall.com/srv/goods-srv/handler"
)

func RegisterSrv(server *grpc.Server) {

	categoryService := service2.NewCategoryService(repository.NewCategoryRepository())

	goodsProto.RegisterGoodsServer(server, &handler.GoodsServer{
		CategoryService:          categoryService,
	})
}
`
)

func GenInitlialize(req GenReq) {
	path := req.InitializeDir+"/register_srv.go"

	context := gstr.ReplaceByMap(regSrv, g.MapStrStr{
		"goods-srv":                 req.SrvName,
		"Category":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"categoryService":           GetJsonTagFromCase(req.TableName, "CamelLower")+"Service",
		"goodsProto":                req.ProtoName + "Proto",
		"mall.com/mall-common/proto/goods": "mall.com/mall-common/proto/" + req.ProtoName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}