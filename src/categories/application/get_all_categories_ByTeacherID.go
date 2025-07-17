package categoryApplication

import (
	categoryDomain "evaluaciones/src/categories/domain"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type GetAllCategoriesByTeacherID struct {
	repo categoryDomain.CategoryInterface
}

func NewGetAllCategoriesByTeacherID(repo categoryDomain.CategoryInterface) *GetAllCategoriesByTeacherID {
	return &GetAllCategoriesByTeacherID{repo: repo}
}

func (gac *GetAllCategoriesByTeacherID) Run(teacherID int32) ([]*categoryEntity.CategoryEntity, error) {
	return gac.repo.GetAllCategoriesByTeacherID(teacherID)
}
