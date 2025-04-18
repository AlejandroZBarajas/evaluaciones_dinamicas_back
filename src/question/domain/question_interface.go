package questionDomain

import questionEntity "evaluaciones/src/question/domain/entity"

type QuestionInterface interface {
	CreateQuestion(question *questionEntity.QuestionEntity) error //post

	GetQuestionByID(questionID int32) (*questionEntity.QuestionEntity, error) //get

	UpdateQuestion(questionID int32, question *questionEntity.QuestionEntity) error //PUT

	DeleteQuestion(questionID int32) error //DELETE

	GetAllQuestionsByExam(examid int32) ([]*questionEntity.QuestionEntity, error)

	GetAllQuestionsByCategory(categoryID int32) ([]*questionEntity.QuestionEntity, error)
}
