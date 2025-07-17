package examApplication

import (
	examDomain "evaluaciones/src/exam/domain"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type GetbyTeacherAndCategory struct {
	repo examDomain.ExamInterface
}

func NewGetbyTeacherAndCategory(repo examDomain.ExamInterface) *GetbyTeacherAndCategory {
	return &GetbyTeacherAndCategory{repo: repo}
}

func (uc *GetbyTeacherAndCategory) Run(teacherID int32, categoryID int32) ([]examEntity.ExamEntity, error) {
	return uc.repo.GetbyTeacherAndCategory(teacherID, categoryID)
}
