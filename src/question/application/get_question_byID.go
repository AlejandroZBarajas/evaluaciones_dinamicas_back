package questionApplication

import (
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type GetQuestionByID struct {
	repo questionDomain.QuestionInterface
}

func NewGetQuestionByID(repo questionDomain.QuestionInterface) *GetQuestionByID {
	return &GetQuestionByID{repo: repo}
}

func (gq *GetQuestionByID) Run(id int32) (*questionEntity.QuestionEntity, error) {
	return gq.repo.GetQuestionByID(id)
}
