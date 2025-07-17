package studentExamApplication

import (
	studentExamDomain "evaluaciones/src/student_exam/domain"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type EvaluateStudentExam struct {
	repo studentExamDomain.StudentExamInterface
}

func NewEvaluateStudentExam(repo studentExamDomain.StudentExamInterface) *EvaluateStudentExam {
	return &EvaluateStudentExam{repo}
}

func (e *EvaluateStudentExam) Run(submission *studentExamEntity.ExamSubmissionInput) (*studentExamEntity.StudentExamResult, error) {
	return e.repo.EvaluateStudentExam(submission)
}
