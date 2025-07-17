package questionApplication

import (
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type CreateQuestion struct {
	repo questionDomain.QuestionInterface
}

func NewCreateQuestion(repo questionDomain.QuestionInterface) *CreateQuestion {
	return &CreateQuestion{repo: repo}
}

func (cq *CreateQuestion) Run(data map[string]interface{}, categoryID int32, examID int32) error {
	question := questionEntity.CreateQuestion(data, categoryID, examID)
	return cq.repo.CreateQuestion(question)
}
