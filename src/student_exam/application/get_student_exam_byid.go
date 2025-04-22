package studentExamApplication

import (
	studentExamDomain "evaluaciones/src/student_exam/domain"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type GetStudentExamByID struct {
	repo studentExamDomain.StudentExamInterface
}

func NewGetStudentExamByID(repo studentExamDomain.StudentExamInterface) *GetStudentExamByID {
	return &GetStudentExamByID{repo: repo}
}

func (g *GetStudentExamByID) Run(studentExamID int32) (*studentExamEntity.StudentExamEntity, error) {
	return g.repo.GetStudentExamByID(studentExamID)
}
