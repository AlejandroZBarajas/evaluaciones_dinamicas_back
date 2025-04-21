package examApplication

import (
	examDomain "evaluaciones/src/exam/domain"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type CreateExam struct {
	repo examDomain.ExamInterface
}

func NewCreateExam(repo examDomain.ExamInterface) *CreateExam {
	return &CreateExam{repo: repo}
}

func (uc *CreateExam) Run(name string, totalQuestions int32, teacherID int32, categoryID int32) error {
	exam := examEntity.CreateExam(name, totalQuestions, teacherID, categoryID)
	return uc.repo.CreateExam(exam)
}
