package categoryDomain

import categoryEntity "evaluaciones/src/categories/domain/entity"

type CategoryInterface interface {
	CreateCategory(category *categoryEntity.CategoryEntity) error
	GetCategoryByID(id int32) (*categoryEntity.CategoryEntity, error)
	GetAllCategoriesByTeacherID(teacherID int32) ([]*categoryEntity.CategoryEntity, error)
	UpdateCategory(categoryID int32, category *categoryEntity.CategoryEntity) error
	DeleteCategory(categoryID int32) error
}
