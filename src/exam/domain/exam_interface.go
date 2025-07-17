package examDomain

import examEntity "evaluaciones/src/exam/domain/entity"

type ExamInterface interface {
	CreateExam(exam *examEntity.ExamEntity) error
	GetExamByID(examID int32) (examEntity.ExamEntity, error)
	UpdateExam(examID int32, exam *examEntity.ExamEntity) (*examEntity.ExamEntity, error)
	DeleteExam(examID int32) error

	GetAllByTeacherID(teacherID int32) ([]examEntity.ExamEntity, error)

	GetbyTeacherAndCategory(teacherID int32, categoryID int32) ([]examEntity.ExamEntity, error)
}
