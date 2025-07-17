package questionEntity

type QuestionEntity struct {
	ID           int32                  `json:"id"`
	QuestionData map[string]interface{} `json:"question_data"`
	CategoryID   int32                  `json:"category_id"`
	ExamID       int32                  `json:"exam_id"`
}

func CreateQuestion(
	question map[string]interface{},
	categoryId int32,
	examenID int32,
) *QuestionEntity {
	return &QuestionEntity{
		QuestionData: question,
		CategoryID:   categoryId,
		ExamID:       examenID,
	}
}
