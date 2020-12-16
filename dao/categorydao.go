package dao

import (
	"finalgo/model"
	"finalgo/tool"

	"github.com/rs/zerolog/log"
)

// Categorydao about category's database function
type Categorydao struct {
	*tool.Orm
}

// CreateCategory use for creating new category
func (cd *Categorydao) CreateCategory(Categorymodel model.Category) int64{
	count, err := cd.InsertOne(&Categorymodel)
	if err != nil {
		log.Error().Err(err)
	}
	// the return num count is the id of the insert one, if it more than 0, the meaning is insert success
	return count
}

// Selectallcategory use for Selecting all the category of the blog
func (cd *Categorydao) Selectallcategory() []model.Category{
	var categorys = make([]model.Category, 0)
	err := cd.SQL("select * from category").Find(&categorys)
	if err != nil {
		log.Error().Err(err)
	}
	return categorys
}

// Deletecategorybycategory use for deleting category
func (cd *Categorydao) Deletecategorybycategory(mc model.Category) bool {
	// implement dao delete
 	result, err := cd.Delete(&mc)
	if err != nil {
		log.Error().Err(err)
	}
	if result > 0 {
		return true
	}
	return false
}

// Updatecategory use for updating category
func (cd *Categorydao) Updatecategory(mc model.Category) bool {
	result, err := cd.Update(&mc)
	if err != nil {
		log.Error().Err(err)
	}
	if result > 0 {
		return true
	}
	return false
}

