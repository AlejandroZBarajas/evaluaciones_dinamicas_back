package questionEntity

type QuestionEntity struct {
	ID           int32                  `json:"id"`
	QuestionData map[string]interface{} `json:"question_data"`
	CategoryID   int32                  `json:"category_id"`
	ExamID       int32                  `json:"exam_id"`
}

func CreateQuestion(question map[string]interface{}, categoryId int32, examenID int32) *QuestionEntity {
	return &QuestionEntity{QuestionData: question, CategoryID: categoryId, ExamID: examenID}
}

/*
question := map[string]interface{}{
    "pregunta":  "Seleccione lenguajes de programación",
    "respuestas": []string{"Go", "Java"},
    "malas":      []string{"Photoshop", "CSS"},
}

newQuestion := questionEntity.CreateQuestion(question, 1, 10)
fmt.Printf("%+v\n", newQuestion)
*/
