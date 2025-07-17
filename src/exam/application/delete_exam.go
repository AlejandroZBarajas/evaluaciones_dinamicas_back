package examApplication

import examDomain "evaluaciones/src/exam/domain"

type DeleteExam struct {
	repo examDomain.ExamInterface
}

func NewDeleteExam(repo examDomain.ExamInterface) *DeleteExam {
	return &DeleteExam{repo: repo}
}

func (uc *DeleteExam) Run(examID int32) error {
	return uc.repo.DeleteExam(examID)
}
