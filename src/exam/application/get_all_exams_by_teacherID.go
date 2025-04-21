package examApplication

import (
	examDomain "evaluaciones/src/exam/domain"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type GetAllExamsByTeacherID struct {
	repo examDomain.ExamInterface
}

func NewGetAllExamsByTeacherID(repo examDomain.ExamInterface) *GetAllExamsByTeacherID {
	return &GetAllExamsByTeacherID{repo: repo}
}

func (uc *GetAllExamsByTeacherID) Run(teacherID int32) (*[]examEntity.ExamEntity, error) {
	return uc.repo.GetAllByTeacherID(teacherID)
}
