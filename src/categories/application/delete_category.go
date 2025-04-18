package categoryApplication

import categoryDomain "evaluaciones/src/categories/domain"

type DeleteCategory struct {
	repo categoryDomain.CategoryInterface
}

func NewDeleteCategory(repo categoryDomain.CategoryInterface) *DeleteCategory {
	return &DeleteCategory{repo: repo}
}

func (dc *DeleteCategory) Run(categoryID int32) error {
	return dc.repo.DeleteCategory(categoryID)
}
