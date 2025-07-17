package studentExamEntity

type ExamSubmissionInput struct {
	StudentExamID int32 `json:"student_exam_id"`
	Responses     []struct {
		QuestionID int32       `json:"question_id"`
		Answer     interface{} `json:"answer"`
	} `json:"responses"`
}
