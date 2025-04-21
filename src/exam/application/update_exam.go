package examApplication

import (
	examDomain "evaluaciones/src/exam/domain"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type UpdateExam struct {
	repo examDomain.ExamInterface
}

func NewUpdateExam(repo examDomain.ExamInterface) *UpdateExam {
	return &UpdateExam{repo: repo}
}

func (uc *UpdateExam) Run(examID int32, updated *examEntity.ExamEntity) (*examEntity.ExamEntity, error) {
	return uc.repo.UpdateExam(examID, updated)
}
