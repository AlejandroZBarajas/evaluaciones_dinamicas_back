package questionDomain

import questionEntity "evaluaciones/src/question/domain/entity"

type QuestionInterface interface {
	CreateQuestion(question *questionEntity.QuestionEntity) error

	GetQuestionByID(questionID int32) (*questionEntity.QuestionEntity, error)

	UpdateQuestion(questionID int32, question *questionEntity.QuestionEntity) error

	DeleteQuestion(questionID int32) error

	GetAllQuestionsByExam(examid int32) ([]*questionEntity.QuestionEntity, error)

	GetAllQuestionsByCategory(categoryID int32) ([]*questionEntity.QuestionEntity, error)
}
