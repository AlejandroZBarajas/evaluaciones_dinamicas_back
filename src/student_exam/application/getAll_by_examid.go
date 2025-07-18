package studentExamApplication

import (
	studentExamDomain "evaluaciones/src/student_exam/domain"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type GetAllByExamID struct {
	repo studentExamDomain.StudentExamInterface
}

func NewGetAllByExamID(repo studentExamDomain.StudentExamInterface) *GetAllByExamID {
	return &GetAllByExamID{repo: repo}
}

func (g *GetAllByExamID) Run(examID int32) ([]studentExamEntity.StudentExamEntity, error) {
	return g.repo.GetAllByExamID(examID)
}
