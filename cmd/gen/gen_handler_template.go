package gen

const handlerBaseTemplateContext = `
package handler

import (
	goodsProto "mall.com/mall-common/proto/goods"
	"mall.com/srv/goods-srv/domain/service"
)

type GoodsServer struct {
	goodsProto.UnimplementedGoodsServer
	CategoryService          service.ICategoryService
}
`

const handlerTemplateContext = `
package handler

import (
	"context"
	"strconv"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	goodsProto "mall.com/mall-common/proto/goods"
	"mall.com/srv/goods-srv/domain/model"
)

// 创建
func (s *GoodsServer) CreateCategory(ctx context.Context, req *goodsProto.CategoryInfoRequest) (*goodsProto.CategoryInfoResponse, error) {
	category := &model.Category{}
	{{create}}
	id, err := s.CategoryService.AddCategory(ctx, category)
	if err != nil {
		return nil, status.Error(codes.Aborted, "CreateCategory func error")
	}

	return &goodsProto.CategoryInfoResponse{Id: id}, nil
}

// 删除
func (s *GoodsServer) DeleteCategory(ctx context.Context, req *goodsProto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := s.CategoryService.DeleteCategory(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "DeleteCategory id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	return &emptypb.Empty{}, nil
}

// 修改
func (s *GoodsServer) UpdateCategory(ctx context.Context, req *goodsProto.CategoryInfoRequest) (*emptypb.Empty, error) {
	category := &model.Category{}
	{{update}}
	err := s.CategoryService.UpdateCategory(ctx, category)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Aborted, "UpdateCategory func error")
	}

	return &emptypb.Empty{}, nil
}

// 根据id查找
func (s *GoodsServer) FindCategoryById(ctx context.Context, req *goodsProto.FindCategoryRequest) (*goodsProto.CategoryInfoResponse, error) {
	reply := &goodsProto.CategoryInfoResponse{}
	category, err := s.CategoryService.FindCategoryByID(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, "FindCategoryById id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	{{find}}
	return reply, nil
}

// 分页查找
func (s *GoodsServer) QueryPageCategory(ctx context.Context, req *goodsProto.QueryPageCategoryRequest) (*goodsProto.QueryPageCategoryResponse, error) {
	var res goodsProto.QueryPageCategoryResponse
	categoryList, resPage, err := s.CategoryService.FindPage(ctx, req.Conditions, req.Page)
	if err != nil {
		return nil, err
	}

	res.Page = resPage

	for _, v := range categoryList {
		info := &goodsProto.CategoryInfoResponse{
			{{pageOne}}
		}
		res.PageList = append(res.PageList, info)
	}

	return &res, nil
}
`