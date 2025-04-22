package studentExamApplication

import (
	studentExamDomain "evaluaciones/src/student_exam/domain"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type CreateStudentExam struct {
	repo studentExamDomain.StudentExamInterface
}

func NewCreateStudentExam(repo studentExamDomain.StudentExamInterface) *CreateStudentExam {
	return &CreateStudentExam{repo: repo}
}

func (c *CreateStudentExam) Run(studentExam *studentExamEntity.StudentExamEntity) error {
	return c.repo.CreateStudentExam(studentExam)
}
