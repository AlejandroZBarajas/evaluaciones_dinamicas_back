package questionApplication

import questionDomain "evaluaciones/src/question/domain"

type DeleteQuestion struct {
	repo questionDomain.QuestionInterface
}

func NewDeleteQuestion(repo questionDomain.QuestionInterface) *DeleteQuestion {
	return &DeleteQuestion{repo: repo}
}

func (dq *DeleteQuestion) Run(id int32) error {
	return dq.repo.DeleteQuestion(id)
}
