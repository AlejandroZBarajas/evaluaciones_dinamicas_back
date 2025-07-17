package resulstEntity

import "time"

type ResultsEntity struct {
	ID            int32     `jaon:"id"`
	StundetExamID int32     `json:"student_exam_id"`
	Result        string    `json:"result"`
	CreatedAt     time.Time `json:"created_at"`
}

func CreateResult(id int32, studentExamID int32, result string, cratedat time.Time) *ResultsEntity {
	return &ResultsEntity{ID: id, StundetExamID: studentExamID, Result: result, CreatedAt: cratedat}
}
