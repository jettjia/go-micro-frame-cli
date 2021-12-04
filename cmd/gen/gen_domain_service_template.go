package gen

const serviceTemplateContext = `
package service

import (
	"context"

	"goods-srv/domain/model"
	"goods-srv/domain/repository"
	goodsProto "mall.com/mall-common/proto/goods"
)

type ICategoryService interface {
	AddCategory(context.Context, *model.Category) (ID uint64, err error)
	DeleteCategory(ctx context.Context, ID uint64) error
	UpdateCategory(context.Context, *model.Category) error
	FindCategoryByID(ctx context.Context, ID uint64) (*model.Category, error)
	FindPage(context.Context, []*goodsProto.Query, *goodsProto.PageData) ([]model.Category, *goodsProto.PageData, error)
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

// 分页
func (u *CategoryService) FindPage(ctx context.Context, query []*goodsProto.Query, reqPage *goodsProto.PageData) ([]model.Category, *goodsProto.PageData, error) {
	return u.CategoryRepository.FindPage(query, reqPage)
}
`
