package gen

// api
const webApiTemplateContext = `
package goods

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"

	goodsProto "mall.com/mall-common/proto/goods"
	"mall.com/web/api-backend-general-system-web/global"
	"mall.com/web/api-backend-general-system-web/trans/dto"
	"mall.com/web/api-backend-general-system-web/trans/vo"
	"mall.com/web/common-web/response"
	"mall.com/web/common-web/util"
)

// ApiCreateProductLog 增
func ApiCreateProductLog(ctx *gin.Context) {
	// 解析参数
	dtoData := &dto.CreateProductLogReq{}
	if err := ctx.BindJSON(dtoData); err != nil {
		response.Failed(ctx, "Failed to parse parameters "+err.Error())
		return
	}

	// 参数过滤
	if err := g.Validator().CheckStruct(context.TODO(), dtoData); err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	// grpc远程调用
	rsp, err := global.GoodsSrvClient.CreateProductLog(context.WithValue(context.Background(), "ginContext", ctx), &goodsProto.ProductLogInfoRequest{
		ProductId: dtoData.ProductId,
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx, "goods-srv")
		return
	}

	// 返回
	voData := vo.CreateProductLogRes{}
	voData.Id = rsp.Id

	response.Success(ctx, voData)
}

// ApiDeleteProductLog 删
func ApiDeleteProductLog(ctx *gin.Context) {
	// 解析参数
	dtoData := &dto.DeleteProductLogReq{}
	if err := ctx.BindJSON(dtoData); err != nil {
		response.Failed(ctx, "Failed to parse parameters "+err.Error())
		return
	}

	// 参数过滤
	if err := g.Validator().CheckStruct(context.TODO(), dtoData); err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	// grpc远程调用
	_, err := global.GoodsSrvClient.DeleteProductLog(context.WithValue(context.Background(), "ginContext", ctx), &goodsProto.DeleteProductLogRequest{
		Id: dtoData.Id,
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx, "goods-srv")
		return
	}

	// 返回
	response.Success(ctx, nil)
}

// ApiUpdateProductLog 改
func ApiUpdateProductLog(ctx *gin.Context) {
	// 解析参数
	dtoData := &dto.UpdateProductLogReq{}
	if err := ctx.BindJSON(dtoData); err != nil {
		response.Failed(ctx, "Failed to parse parameters "+err.Error())
		return
	}

	// 参数过滤
	if err := g.Validator().CheckStruct(context.TODO(), dtoData); err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	// grpc远程调用
	_, err := global.GoodsSrvClient.UpdateProductLog(context.WithValue(context.Background(), "ginContext", ctx), &goodsProto.ProductLogInfoRequest{
		Id:        dtoData.Id,
		ProductId: dtoData.Info.ProductId,
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx, "goods-srv")
		return
	}

	// 返回
	response.Success(ctx, nil)
}

// ApiFindProductLog 单条
func ApiFindProductLog(ctx *gin.Context) {
	// 解析参数
	dtoData := &dto.FindProductLogReq{}
	if err := ctx.BindJSON(dtoData); err != nil {
		response.Failed(ctx, "Failed to parse parameters "+err.Error())
		return
	}

	// 参数过滤
	if err := g.Validator().CheckStruct(context.TODO(), dtoData); err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	// grpc远程调用
	rsp, err := global.GoodsSrvClient.FindProductLogById(context.WithValue(context.Background(), "ginContext", ctx), &goodsProto.FindProductLogRequest{
		Id: dtoData.Id,
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx, "goods-srv")
		return
	}

	// 返回
	voData := vo.FindProductLogRes{}
	voData.Id = rsp.Id
	voData.ProductId = rsp.ProductId
	voData.OpTitle = rsp.OpTitle
	voData.OpInfo = rsp.OpInfo
	voData.OpUserId = rsp.OpUserId
	voData.OpUserName = rsp.OpUserName
	voData.Ip = rsp.Ip

	response.Success(ctx, voData)
}

// ApiQueryPageProductLog 分页
func ApiQueryPageProductLog(ctx *gin.Context) {
	// 解析参数
	dtoData := &dto.QueryPageProductLogReq{}
	if err := ctx.BindJSON(dtoData); err != nil {
		response.Failed(ctx, "Failed to parse parameters "+err.Error())
		return
	}

	// grpc远程调用
	rsp, err := global.GoodsSrvClient.QueryPageProductLog(context.WithValue(context.Background(), "ginContext", ctx), &goodsProto.QueryPageProductLogRequest{
		Conditions: dtoData.QueryList,
		Page:       dtoData.Page,
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx, "goods-srv")
		return
	}

	// 返回
	if rsp.Page.TotalNumber < 1 {
		response.Success(ctx, nil)
	}

	voData := vo.QueryPageProductLogRes{}
	voData.Page = rsp.Page

	var list []vo.FindProductLogRes
	for _, v := range rsp.PageList {
		info := vo.FindProductLogRes{
			Id:         v.Id,
			ProductId:  v.ProductId,
			OpTitle:    v.OpTitle,
			OpInfo:     v.OpInfo,
			OpUserId:   v.OpUserId,
			OpUserName: v.OpUserName,
			Ip:         v.Ip,
		}

		list = append(list, info)
	}
	voData.List = list

	response.Success(ctx, voData)
}

`

// init srv_conn
const webInitializeSrvConnContext = `
package initialize

import (
	"github.com/jettjia/go-micro-frame/core/register/nacos"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"

	"mall.com/mall-common/proto/goods"
	"mall.com/web/api-backend-general-system-web/global"
)

func InitSrvConn() {
	c := nacos.NewRegistryClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace, global.NacosConfig.User, global.NacosConfig.Password)
	goodsConn, err := c.Discovery(global.ServerConfig.GoodsSrvInfo.Name, global.ServerConfig.Env)
	if err != nil {
		mylogger.Fatal("Failed to link grpc server")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}
`
// init router
const webInitializeRouterContext = `
package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall.com/web/api-backend-general-system-web/middleware"
	"mall.com/web/api-backend-general-system-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("/bgsw/v1")
	router.InitGoodsRouter(ApiGroup)

	return Router
}
`

// router
const webRouterContext = `
package router

import (
	"github.com/gin-gonic/gin"

	"mall.com/web/api-backend-general-system-web/api/goods"
	"mall.com/web/api-backend-general-system-web/middleware"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods").Use(middleware.Trace())
	{
		GoodsRouter.GET("productLog/create", goods.ApiCreateProductLog)
		GoodsRouter.GET("productLog/delete", goods.ApiDeleteProductLog)
		GoodsRouter.GET("productLog/update", goods.ApiUpdateProductLog)
		GoodsRouter.GET("productLog/info", goods.ApiFindProductLog)
		GoodsRouter.GET("productLog/page", goods.ApiQueryPageProductLog)
	}
}
`

// dto
const webDtoContext = `
package dto

import (
	goodsProto "mall.com/mall-common/proto/goods"
)

// 增
type CreateProductLogReq struct {
	{{dtoCreate}}
}

// 删
type DeleteProductLogReq struct {
	Id   uint64 ` + "`" + `v:"integer|min:1" dc:"id"` + "`" + `
}

// 改
type UpdateProductLogReq struct {
	Id   uint64 ` + "`" + `v:"integer|min:1" dc:"id"` + "`" + `
	Info CreateProductLogReq
}

// 单条
type FindProductLogReq struct {
	Id   uint64 ` + "`" + `v:"integer|min:1" dc:"id"` + "`" + `
}

// 分页
type QueryPageProductLogReq struct {
	QueryList []*goodsProto.Query
	Page *goodsProto.PageData
}
`

// vo
const webVoContext = `
package vo

import (
	goodsProto "mall.com/mall-common/proto/goods"
)

type CreateProductLogRes struct {
	Id uint64 ` + "`" + `json:"id" ` + "`" + `
}

type FindProductLogRes struct {
	{{voInfo}}
}

type QueryPageProductLogRes struct {
	List []FindProductLogRes
	Page *goodsProto.PageData
}

`
