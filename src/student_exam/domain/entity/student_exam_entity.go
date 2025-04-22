package studentExamEntity

import "time"

type StudentExamEntity struct {
	StudentExamID  int32                  `json:"id"`
	ExamID         int32                  `json:"exam_id"`
	StudentID      int32                  `json:"student_id"`
	TotalQuestions int32                  `json:"total_questions"`
	CorrectAnswers int32                  `json:"correct_answers"`
	Responses      map[string]interface{} `json:"responses"`
	SubmittedAt    time.Time              `json:"submitted_at"`
}

func CreateStudentExam(
	id int32,
	examid int32,
	studentid int32,
	totalquestions int32,
	correctAns int32,
	responses map[string]interface{},
	submittedAt time.Time,
) *StudentExamEntity {
	return &StudentExamEntity{
		StudentExamID:  id,
		ExamID:         examid,
		StudentID:      studentid,
		TotalQuestions: totalquestions,
		CorrectAnswers: correctAns,
		Responses:      responses,
		SubmittedAt:    submittedAt,
	}
}
