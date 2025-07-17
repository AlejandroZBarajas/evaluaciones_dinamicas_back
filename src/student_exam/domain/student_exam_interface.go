package studentExamDomain

import (
	questionEntity "evaluaciones/src/question/domain/entity"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type StudentExamInterface interface {
	CreateStudentExam(studentExam *studentExamEntity.StudentExamEntity) error

	GetAllByExamID(examID int32) ([]*studentExamEntity.StudentExamEntity, error)

	GetStudentExamByID(studentExamID int32) (*studentExamEntity.StudentExamEntity, error)

	DeleteStudentExam(studentExamid int32) error

	EvaluateStudentExam(submission *studentExamEntity.ExamSubmissionInput) (*studentExamEntity.StudentExamResult, error)

	GenerateRandomExam(input *studentExamEntity.RandomExamInput) ([]*questionEntity.QuestionEntity, error)
}
