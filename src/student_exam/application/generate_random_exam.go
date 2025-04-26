package studentExamApplication

import (
	questionEntity "evaluaciones/src/question/domain/entity"
	studentExamDomain "evaluaciones/src/student_exam/domain"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type GenerateRandomExam struct {
	repo studentExamDomain.StudentExamInterface
}

func NewGenerateRandomExam(repo studentExamDomain.StudentExamInterface) *GenerateRandomExam {
	return &GenerateRandomExam{repo}
}

func (g *GenerateRandomExam) Run(random *studentExamEntity.RandomExamInput) ([]*questionEntity.QuestionEntity, error) {
	return g.repo.GenerateRandomExam(random)
}
