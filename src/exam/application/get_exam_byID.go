package examApplication

import (
	examDomain "evaluaciones/src/exam/domain"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type GetExamByID struct {
	repo examDomain.ExamInterface
}

func NewGetExamByID(repo examDomain.ExamInterface) *GetExamByID {
	return &GetExamByID{repo: repo}
}

func (uc *GetExamByID) Run(examID int32) (examEntity.ExamEntity, error) {
	return uc.repo.GetExamByID(examID)
}
