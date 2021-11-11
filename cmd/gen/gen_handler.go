package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)

func GenHandler(req GenReq) {
	doBase(req)
	doHandler(req)
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
	newStr := strings.Replace(baseHandler, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newStr = strings.Replace(newStr, "goods_srv", req.SrvName, -1)

	util.WriteStringToFileMethod(req.HandlerDir+"/base.go", newStr)
}

var (
	handlerStr = `
package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"

	"goods-srv/domain/model"
	goods_srv "goods-srv/proto/goods_srv"
)

// 创建
func (s *GoodsServer) CreateCategory(ctx context.Context, req *goods_srv.CategoryInfoRequest) (*goods_srv.CategoryInfoResponse, error) {
	category := &model.Category{}
	// todo
	// category.Pid = req.Pid


	id, err := s.CategoryService.AddCategory(ctx, category)
	if err != nil {
		return nil, status.Error(codes.Aborted, "CreateCategory func error")
	}

	return &goods_srv.CategoryInfoResponse{Id: id}, nil
}

// 删除
func (s *GoodsServer) DeleteCategory(ctx context.Context, req *goods_srv.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := s.CategoryService.DeleteCategory(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "DeleteCategory id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	return &emptypb.Empty{}, nil
}

// 修改
func (s *GoodsServer) UpdateCategory(ctx context.Context, req *goods_srv.CategoryInfoRequest) (*emptypb.Empty, error) {
	category := &model.Category{}
	// todo
	// category.Pid = req.Pid

	err := s.CategoryService.UpdateCategory(ctx, category)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Aborted, "UpdateCategory func error")
	}

	return &emptypb.Empty{}, nil
}

// 根据id查找
func (s *GoodsServer) FindCategoryById(ctx context.Context, req *goods_srv.FindCategoryRequest) (*goods_srv.CategoryInfoResponse, error) {
	reply := &goods_srv.CategoryInfoResponse{}
	category, err := s.CategoryService.FindCategoryByID(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, "FindCategoryById id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	// todo
	// reply.Id = category.ID

	return reply, nil
}

// 分页查找
func (s *GoodsServer) QueryPageCategory(ctx context.Context, req *goods_srv.QueryPageRequest) (*goods_srv.QueryPageResponse, error) {
	var res goods_srv.QueryPageResponse
	categories, resPage, err := s.CategoryService.FindPage(ctx, req.Conditions, req.Page)
	if err != nil {
		return nil, err
	}

	res.Page = resPage

	for _, v := range categories {
		info := &goods_srv.CategoryInfoResponse{
			// todo
			// Id = v.ID
		}
		res.PageList = append(res.PageList, info)
	}

	return &res, nil
}
`
)

func doHandler(req GenReq) {
	newStr := strings.Replace(handlerStr, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newStr = strings.Replace(newStr, "goods_srv", req.SrvName, -1)

	util.WriteStringToFileMethod(req.HandlerDir+"/"+req.TableName+".go", newStr)
}
