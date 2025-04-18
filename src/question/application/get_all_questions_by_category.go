package questionApplication

import (
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type GetAllQuestionsByCategory struct {
	repo questionDomain.QuestionInterface
}

func NewGetAllQuestionsByCategory(repo questionDomain.QuestionInterface) *GetAllQuestionsByCategory {
	return &GetAllQuestionsByCategory{repo: repo}
}

func (ga *GetAllQuestionsByCategory) Run(categoryID int32) ([]*questionEntity.QuestionEntity, error) {
	return ga.repo.GetAllQuestionsByCategory(categoryID)
}
