package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)

var (
	serviceStr = `package service

import (
	"context"

	"goods_srv/domain/model"
	"goods_srv/domain/repository"
	goods_srv "goods_srv/proto/goods_srv"
)

type ICategoryService interface {
	AddCategory(context.Context, *model.Category) (ID uint64, err error)
	DeleteCategory(ctx context.Context, ID uint64) error
	UpdateCategory(context.Context, *model.Category) error
	FindCategoryByID(ctx context.Context, ID uint64) (*model.Category, error)
	FindPage(context.Context, []*goods_srv.Query, *goods_srv.PageData) ([]model.Category, *goods_srv.PageData, error)
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryService(CategoryRepository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{CategoryRepository}
}

// 插入
func (u *CategoryService) AddCategory(ctx context.Context, Category *model.Category) (uint64, error) {
	return u.CategoryRepository.CreateCategory(Category)
}

// 删除
func (u *CategoryService) DeleteCategory(ctx context.Context, ID uint64) error {
	return u.CategoryRepository.DeleteCategoryByID(ID)
}

// 修改
func (u *CategoryService) UpdateCategory(ctx context.Context, Category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(Category)
}

// 通过ID查找
func (u *CategoryService) FindCategoryByID(ctx context.Context, ID uint64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(ID)
}

// 所有
func (u *CategoryService) FindAll(ctx context.Context, query []*goods_srv.Query) ([]model.Category, error) {
	return u.CategoryRepository.FindAll(query)
}

// 分页
func (u *CategoryService) FindPage(ctx context.Context, query []*goods_srv.Query, reqPage *goods_srv.PageData) ([]model.Category, *goods_srv.PageData, error) {
	return u.CategoryRepository.FindPage(query, reqPage)
}`
)
func GenService(req GenReq) {
	newService := strings.Replace(serviceStr, "Category", GetJsonTagFromCase(req.TableName, "Camel"), -1)
	newService = strings.Replace(newService, "goods_srv", req.SrvName, -1)

	util.WriteStringToFileMethod(req.ServiceDir+"/"+req.TableName+"_service.go", newService)
}