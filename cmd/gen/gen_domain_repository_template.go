package gen

const repositoryBaseTemplate= `
package repository

import "github.com/jinzhu/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
`

const repositoryTemplate = `
	package repository

import (
	"errors"
	"math"
	"time"

	"goods-srv/domain/model"
	"goods-srv/global"
	goodsProto "mall.com/mall-common/proto/goods"
)

type ICategoryRepository interface {
	InitTable() error
	CreateCategory(*model.Category) (ID uint64, err error)
	DeleteCategoryByID(ID uint64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(ID uint64) (*model.Category, error)
	FindPage([]*goodsProto.Query, *goodsProto.PageData) ([]model.Category, *goodsProto.PageData, error)
}

type CategoryRepository struct {
}

func NewCategoryRepository() ICategoryRepository {
	return &CategoryRepository{}
}

func tableName() string {
	return "table_name"
}

// 初始化表
func (u *CategoryRepository) InitTable() error {
	return global.DB.CreateTable(&model.Category{}).Error
}

// 创建
func (u *CategoryRepository) CreateCategory(category *model.Category) (categoryID uint64, err error) {
	category.CreatedAt = time.Now()
	return category.Id, global.DB.Create(category).Error
}

// 根据ID删除
func (u *CategoryRepository) DeleteCategoryByID(CategoryID uint64) error {
	return global.DB.Where("id = ?", CategoryID).Delete(&model.Category{}).Error
}

// 更新信息
func (u *CategoryRepository) UpdateCategory(Category *model.Category) error {
	return global.DB.Model(Category).Update(&Category).Error
}

// 根据ID查找信息
func (u *CategoryRepository) FindCategoryByID(categoryID uint64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, global.DB.First(category, categoryID).Error
}

// 分页查找
func (u *CategoryRepository) FindPage(conditions []*goodsProto.Query, reqPage *goodsProto.PageData) ([]model.Category, *goodsProto.PageData, error) {
	var condition string
	var total uint32
	var categories []model.Category
	var resPage goodsProto.PageData

	condition = model.GenerateQueryCondition(conditions)

	global.DB.Model(&model.Category{}).Count(&total)

	if total == 0 {
		return nil, nil, errors.New("database data is empty")
	}

	err := global.DB.Table(tableName()).
		Select(tableName()+".*").
		Where(condition).
		Scopes(Paginate(int(reqPage.Page), int(reqPage.PageSize))).
		Find(&categories).Error

	if err != nil  {
		return nil, nil, err
	}

	resPage.Page = reqPage.Page
	resPage.PageSize = reqPage.PageSize
	resPage.TotalNumber = total

	resPage.TotalPage = uint32(int(math.Ceil(float64(total) / float64(reqPage.PageSize))))

	return categories, &resPage, err
}
`
