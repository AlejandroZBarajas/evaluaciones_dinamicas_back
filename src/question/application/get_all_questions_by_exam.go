package questionApplication

import (
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type GetAllQuestionsByExam struct {
	repo questionDomain.QuestionInterface
}

func NewGetAllQuestionsByExam(repo questionDomain.QuestionInterface) *GetAllQuestionsByExam {
	return &GetAllQuestionsByExam{repo: repo}
}

func (ga *GetAllQuestionsByExam) Run(examID int32) ([]*questionEntity.QuestionEntity, error) {
	return ga.repo.GetAllQuestionsByExam(examID)
}
