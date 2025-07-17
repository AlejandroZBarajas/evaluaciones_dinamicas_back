package studentExamEntity

type StudentExamResult struct {
	CorrectAnswers int32  `json:"correct_answers"`
	TotalQuestions int32  `json:"total_questions"`
	Score          string `json:"score"`
}
