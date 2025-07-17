package questionApplication

import (
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type UpdateQuestion struct {
	repo questionDomain.QuestionInterface
}

func NewUpdateQuestion(repo questionDomain.QuestionInterface) *UpdateQuestion {
	return &UpdateQuestion{repo: repo}
}

func (uq *UpdateQuestion) Run(id int32, data map[string]interface{}, categoryID int32, examID int32) error {
	question := questionEntity.CreateQuestion(data, categoryID, examID)
	return uq.repo.UpdateQuestion(id, question)
}
